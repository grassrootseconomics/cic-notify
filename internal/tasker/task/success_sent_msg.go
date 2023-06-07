package task

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/grassrootseconomics/celoutils"
	"github.com/grassrootseconomics/cic-notify/internal/graphql"
	"github.com/grassrootseconomics/cic-notify/internal/locale"
	"github.com/grassrootseconomics/cic-notify/internal/notify"
	"github.com/grassrootseconomics/w3-celo-patch/module/eth"
	"github.com/hibiken/asynq"
)

type successSentMsg struct {
	ShortHash     string
	TransferValue string
	VoucherSymbol string
	SentTo        string
	DateString    string
	// These are passed to the channel provider e.g. AfricasTalking, Telegram, e.t.c.
	ChannelType       graphql.Interface_type_enum
	ChannelIdentifier string
	Language          string
	// These are used to update the balance value as it is on chain.
	BlockchainAddress string
	VoucherAddress    string
}

func SuccessSentMsgProcessor(n *notify.Notify) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		var (
			payload successSentMsg
			balance big.Int
		)

		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			return err
		}

		if err := n.CeloProvider.Client.CallCtx(
			ctx,
			eth.CallFunc(
				funcBalanceOf,
				celoutils.HexToAddress(payload.VoucherAddress),
				celoutils.HexToAddress(payload.BlockchainAddress),
			).Returns(&balance),
		); err != nil {
			return err
		}

		templatePayload := struct {
			ShortHash      string
			TransferValue  string
			VoucherSymbol  string
			SentTo         string
			DateString     string
			CurrentBalance string
		}{
			payload.ShortHash,
			payload.TransferValue,
			payload.VoucherSymbol,
			payload.SentTo,
			payload.DateString,
			truncateVoucherValue(balance.Uint64()),
		}

		msgPayload := n.Templates.PrepareLocale(
			locale.SuccessSentTemplate,
			"eng",
			templatePayload,
		)

		if err := routeMessage(
			ctx,
			n.TaskerClient,
			payload.ChannelType,
			msgPayload,
			payload.ChannelIdentifier,
		); err != nil {
			return nil
		}

		return nil
	}
}
