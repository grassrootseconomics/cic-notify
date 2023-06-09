package task

import (
	"context"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/grassrootseconomics/cic-notify/internal/notify"
	"github.com/hibiken/asynq"
)

type tgPayload struct {
	ChatId  int64
	Message string
}

func TgPushProcessor(n *notify.Notify) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		var (
			payload tgPayload
		)

		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			return err
		}

		if n.DisablePush {
			n.Logg.Info("tg_push_processor: skipping push", "payload", payload.Message, "sent_to", payload.ChatId)
			return nil
		}

		msg := tgbotapi.NewMessage(payload.ChatId, payload.Message)

		tgResponse, err := n.TgClient.Send(msg)
		if err != nil {
			return err
		}
		n.Logg.Info("tg_push_processor: TG push successful", "payload", payload.Message, "sent_to", payload.ChatId, "message_id", tgResponse.MessageID)

		if err := n.Store.CreateTgReceipt(
			ctx,
			tgResponse.MessageID,
		); err != nil {
			return err
		}

		return nil
	}
}
