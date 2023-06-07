package task

import (
	"context"
	"encoding/json"

	"github.com/grassrootseconomics/cic-notify/internal/graphql"
	"github.com/grassrootseconomics/cic-notify/internal/locale"
	"github.com/grassrootseconomics/cic-notify/internal/notify"
	"github.com/hibiken/asynq"
)

type failedMsg struct {
	FailReason string
	// These are passed to the channel provider e.g. AfricasTalking, Telegram, e.t.c.
	ChannelType       graphql.Interface_type_enum
	ChannelIdentifier string
	Language          string
}

func FailedMsgProcessor(n *notify.Notify) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		var (
			payload failedMsg
		)

		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			return err
		}

		templatePayload := struct {
			FailReason string
		}{
			FailReason: payload.FailReason,
		}

		msgPayload := n.Templates.PrepareLocale(
			locale.FailedTemeplate,
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
