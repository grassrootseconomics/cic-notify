package task

import (
	"context"
	"encoding/json"

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

// PrepareMsgProcessor will determine which kind of message should be prepared for sending.
// It also supports cross custodial - non-custodial transfer notifications.
// It handles 4 scenarios:
// 1. Failed transfer from a Custodial user -> Custodial/Non-custodial user.
// 2. Custodial -> Custodial transfer.
// 3. Non-custodial -> Custodial transfer.
// 4. Custodial -> Non-custodial transfer.
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
		// We only send notifications for certified vouchers.
		if (len(resp.Receiver) < 1 && len(resp.Sender) < 1) || len(resp.Vouchers) < 1 {
			n.Logg.Debug("prepare_msg_processor: dropping notification", "tx_hash", payload.TxHash)
			return nil
		}

		// If the transfer failed on-chain, we only notify the transfer initiator/sender and stop further processing.
		if !payload.Success {
			n.Logg.Warn("prepare_msg_processor: on-chain failed transfer", "tx_hash", payload.TxHash, "payload", t.Payload())
			if len(resp.Sender) > 0 {
				failedMsgJobPayload, err := json.Marshal(failedMsg{
					FailReason:        "Tx failed on chain",
					ChannelType:       resp.Sender[0].User.Interface_type,
					ChannelIdentifier: resp.Sender[0].User.Interface_identifier,
					Language:          resp.Sender[0].User.Personal_information.Language_code,
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

		// If the sender is registered on the graph, we send them a message.
		if len(resp.Sender) > 0 {
			n.Logg.Debug("prepare_msg_processor: preparing msg for sender", "tx_hash", payload.TxHash)
			var (
				sentTo string
			)

			// We support sending messages when vouchers are sent to non-custodial/external users too.
			if len(resp.Receiver) < 1 {
				sentTo = formatIdentifier("", "", "", payload.To)
			} else {
				sentTo = formatIdentifier(
					resp.Receiver[0].User.Personal_information.Given_names,
					resp.Receiver[0].User.Personal_information.Family_name,
					resp.Receiver[0].User.Interface_identifier,
					resp.Receiver[0].Blockchain_address,
				)
			}

			successSentMsgData := successSentMsg{
				ShortHash:         formatShortHash(payload.TxHash),
				TransferValue:     truncateVoucherValue(payload.Value),
				VoucherSymbol:     resp.Vouchers[0].Symbol,
				SentTo:            sentTo,
				DateString:        formatDate(payload.Timestamp, n.Timezone),
				ChannelType:       resp.Sender[0].User.Interface_type,
				ChannelIdentifier: resp.Sender[0].User.Interface_identifier,
				Language:          resp.Sender[0].User.Personal_information.Language_code,
				BlockchainAddress: resp.Sender[0].Blockchain_address,
				VoucherAddress:    payload.ContractAddress,
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

		// // If the receiver is registered on the graph, we send them a message from whom they received the vouchers from.
		if len(resp.Receiver) > 0 {
			n.Logg.Debug("prepare_msg_processor: preparing msg for receiver", "tx_hash", payload.TxHash)
			var (
				receivedFrom string
			)

			// We support sending messages when vouchers are received from non-custodial/external users too.
			if len(resp.Sender) < 1 {
				receivedFrom = formatIdentifier("", "", "", payload.From)
			} else {
				receivedFrom = formatIdentifier(
					resp.Sender[0].User.Personal_information.Given_names,
					resp.Sender[0].User.Personal_information.Family_name,
					resp.Sender[0].User.Interface_identifier,
					resp.Sender[0].Blockchain_address,
				)
			}

			successReceivedMsgdata := successReceivedMsg{
				ShortHash:         formatShortHash(payload.TxHash),
				TransferValue:     truncateVoucherValue(payload.Value),
				VoucherSymbol:     resp.Vouchers[0].Symbol,
				ReceivedFrom:      receivedFrom,
				DateString:        formatDate(payload.Timestamp, n.Timezone),
				ChannelType:       resp.Receiver[0].User.Interface_type,
				ChannelIdentifier: resp.Receiver[0].User.Interface_identifier,
				Language:          resp.Receiver[0].User.Personal_information.Language_code,
				BlockchainAddress: resp.Receiver[0].Blockchain_address,
				VoucherAddress:    payload.ContractAddress,
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
		n.Logg.Info("prepare_msg_processor: processing chain event success", "tx_hash", payload.TxHash, "full_payload", t.Payload())

		return nil
	}
}
