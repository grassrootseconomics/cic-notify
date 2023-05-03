package task

import (
	"context"
	"encoding/json"

	"github.com/grassrootseconomics/cic-notify/internal/notify"
	"github.com/hibiken/asynq"
	"github.com/kamikazechaser/africastalking"
)

type atPayload struct {
	RecepientPhone string
	Message        string
}

func AtPushProcessor(n *notify.Notify) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		var (
			payload atPayload
		)

		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			return err
		}

		msg := africastalking.BulkSMSInput{
			From:    n.AtShortCode,
			Message: payload.Message,
			To:      []string{payload.RecepientPhone},
		}

		atResponse, err := n.AtClient.SendBulkSMS(ctx, msg)
		if err != nil {
			return err
		}

		if err := n.Store.CreateAtReceipt(
			ctx,
			atResponse.SMSMessageData.Recipients[0].StatusCode,
			atResponse.SMSMessageData.Recipients[0].MessageID,
		); err != nil {
			return err
		}

		return nil
	}
}
