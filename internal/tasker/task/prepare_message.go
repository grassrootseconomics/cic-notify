package task

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang-module/carbon/v2"
	"github.com/grassrootseconomics/cic-notify/internal/graphql"
	"github.com/grassrootseconomics/cic-notify/internal/notify"
	"github.com/grassrootseconomics/cic-notify/internal/tasker"
	"github.com/hibiken/asynq"
)

type chainEvent struct {
	Block           uint64 `json:"block"`
	From            string `json:"from"`
	To              string `json:"to"`
	ContractAddress string `json:"contractAddress"`
	Success         bool   `json:"success"`
	Timestamp       uint64 `json:"timestamp"`
	TxHash          string `json:"transactionHash"`
	TxIndex         uint   `json:"transactionIndex"`
	Value           uint64 `json:"value"`
}

// PrepareMsgProcessor will attempp to determine which kind of message should be prepared for sending.
// It attempts to also support cross custodial non-custodial transfers.
// It handles approx 4 scenarios:
// 1. Failed transfer from a Custodial user
// 2. Custodial -> Custodial transfer
// 3. Non-custodial -> Custodial transfer
// 4. Custodial -> Non-custodial transfer
func PrepareMsgProcessor(n *notify.Notify) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		var (
			payload chainEvent
		)

		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			return err
		}

		resp, err := graphql.PrepareMessagePayload(
			ctx,
			n.GraphQLClient,
			payload.From,
			payload.To,
			payload.ContractAddress,
		)
		if err != nil {
			return err
		}

		// Check if any of the parties are registered in the GE graph realm.
		// If not, we cannot proceed with any notification action.
		// We additionally also check if the voucher is certified within the GE graph realm.
		if len(resp.Accounts) < 1 || len(resp.Vouchers) < 1 {
			return nil
		}

		// If the transfer failed on-chain, we only notify the transfer initiator.
		if !payload.Success {
			if len(resp.Accounts) > 1 || resp.Accounts[0].Blockchain_address == payload.From {
				failedMsgJobPayload, err := json.Marshal(failedMsg{
					FailReason:        "Tx failed on chain",
					ChannelType:       resp.Accounts[0].User.Interface_type,
					ChannelIdentifier: resp.Accounts[0].User.Interface_identifier,
				})
				if err != nil {
					return nil
				}

				_, err = n.TaskerClient.CreateTask(
					ctx,
					tasker.ProcessFailedMsgTask,
					tasker.DefaultPriority,
					&tasker.Task{
						Payload: failedMsgJobPayload,
					},
				)
				if err != nil {
					return err
				}
			}

			return nil
		}

		// If both parties are registered in the GE graph realm.
		// The first party is always the sender, while the second one is the receiver.
		if len(resp.Accounts) > 1 {
			successSentMsgData := successSentMsg{
				ShortHash:     formatShortHash(payload.TxHash),
				TransferValue: payload.Value / 1000000,
				VoucherSymbol: resp.Vouchers[0].Symbol,
				SentTo: formatIdentifier(
					resp.Accounts[1].User.Personal_information.Given_names,
					resp.Accounts[1].User.Personal_information.Family_name,
					resp.Accounts[1].User.Interface_identifier,
					payload.To,
				),
				DateString:        formatDate(payload.Timestamp, n.Timezone),
				ChannelType:       resp.Accounts[0].User.Interface_type,
				ChannelIdentifier: resp.Accounts[0].User.Interface_identifier,
			}

			successSentJobPayload, err := json.Marshal(successSentMsgData)
			if err != nil {
				return nil
			}

			_, err = n.TaskerClient.CreateTask(
				ctx,
				tasker.ProcessSuccessSentMsgTask,
				tasker.DefaultPriority,
				&tasker.Task{
					Payload: successSentJobPayload,
				},
			)
			if err != nil {
				return err
			}

			successReceivedMsgdata := successReceivedMsg{
				ShortHash:     formatShortHash(payload.TxHash),
				TransferValue: payload.Value / 1000000,
				VoucherSymbol: resp.Vouchers[0].Symbol,
				ReceivedFrom: formatIdentifier(
					resp.Accounts[0].User.Personal_information.Given_names,
					resp.Accounts[0].User.Personal_information.Family_name,
					resp.Accounts[0].User.Interface_identifier,
					payload.From,
				),
				DateString:        formatDate(payload.Timestamp, n.Timezone),
				ChannelType:       resp.Accounts[1].User.Interface_type,
				ChannelIdentifier: resp.Accounts[1].User.Interface_identifier,
			}

			successReceivedJobPayload, err := json.Marshal(successReceivedMsgdata)
			if err != nil {
				return nil
			}

			_, err = n.TaskerClient.CreateTask(
				ctx,
				tasker.ProcessSuccessReceivedMsgTask,
				tasker.DefaultPriority,
				&tasker.Task{
					Payload: successReceivedJobPayload,
				},
			)
			if err != nil {
				return err
			}
		} else {
			// We can only send a message to one party.
			// We need to determine what kind of message we need to send.
			if resp.Accounts[0].Blockchain_address == payload.To {
				// A custodial user has received funds from a user outside the custodial system.
				successReceivedMsgdata := successReceivedMsg{
					ShortHash:     formatShortHash(payload.TxHash),
					TransferValue: payload.Value / 1000000,
					VoucherSymbol: resp.Vouchers[0].Symbol,
					ReceivedFrom: formatIdentifier(
						"",
						"",
						"",
						payload.From,
					),
					DateString:        formatDate(payload.Timestamp, n.Timezone),
					ChannelType:       resp.Accounts[1].User.Interface_type,
					ChannelIdentifier: resp.Accounts[1].User.Interface_identifier,
				}

				successReceivedJobPayload, err := json.Marshal(successReceivedMsgdata)
				if err != nil {
					return nil
				}

				_, err = n.TaskerClient.CreateTask(
					ctx,
					tasker.ProcessSuccessReceivedMsgTask,
					tasker.DefaultPriority,
					&tasker.Task{
						Payload: successReceivedJobPayload,
					},
				)
				if err != nil {
					return err
				}
			}

			if resp.Accounts[0].Blockchain_address == payload.From {
				// A custodial user had sent funds to a user outside the custodial system.
				successSentMsgData := successSentMsg{
					ShortHash:     formatShortHash(payload.TxHash),
					TransferValue: payload.Value / 1000000,
					VoucherSymbol: resp.Vouchers[0].Symbol,
					SentTo: formatIdentifier(
						"",
						"",
						"",
						payload.To,
					),
					DateString:        formatDate(payload.Timestamp, n.Timezone),
					ChannelType:       resp.Accounts[0].User.Interface_type,
					ChannelIdentifier: resp.Accounts[0].User.Interface_identifier,
				}

				successSentJobPayload, err := json.Marshal(successSentMsgData)
				if err != nil {
					return nil
				}

				_, err = n.TaskerClient.CreateTask(
					ctx,
					tasker.ProcessSuccessSentMsgTask,
					tasker.DefaultPriority,
					&tasker.Task{
						Payload: successSentJobPayload,
					},
				)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}
}

// formatDate takes a unix timestamp and timezone and returns a formatted string
// 1649735755981 + Europe/Moscow
// 2021-01-27 13:14:15
func formatDate(unixTimestamp uint64, timeZone string) string {
	return carbon.CreateFromTimestamp(int64(unixTimestamp)).SetTimezone(timeZone).ToDateTimeString()
}

// formatIdentifier takes the first name and last name and returns a formatted name string.
func formatIdentifier(firstName string, lastName string, identifier string, blockchainAddress string) string {
	if firstName != "" || lastName != "" {
		return strings.ToUpper(strings.TrimSpace(fmt.Sprintf("%s %s %s", firstName, lastName, identifier)))
	} else {
		return blockchainAddress
	}
}

// formatShortHash takes a full txHash and returns the last 8 chars of the Ethereum transaction hash (hex).
// 0x1562767d2a01098da599cdea23ff798838a530a17e6072838c425d48.
// Should return 7837424A.
func formatShortHash(txHash string) string {
	return strings.ToUpper(txHash[58:])
}
