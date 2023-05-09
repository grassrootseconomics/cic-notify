package task

import (
	"context"
	"encoding/json"
	"fmt"

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

		passThru, err := n.RedisClient.Get(ctx, passThruKey).Bool()
		if err != nil {
			return err
		}

		if passThru {
			return nil
		}

		msg := africastalking.BulkSMSInput{
			From:    n.AtShortCode,
			Message: payload.Message,
			To:      []string{payload.RecepientPhone},
		}

		atResponse, err := n.AtClient.SendBulkSMS(ctx, msg)
		if err != nil {
			return fmt.Errorf("AT push failed: %v: %w", err, asynq.SkipRetry)
		}
		n.Logg.Info("at_push_processor: AT push successful", "payload", payload.Message, "sent_to", payload.RecepientPhone, "message_id", atResponse.SMSMessageData.Recipients[0].MessageID)

		if err := n.Store.CreateAtReceipt(
			ctx,
			atResponse.SMSMessageData.Recipients[0].StatusCode,
			atResponse.SMSMessageData.Recipients[0].MessageID,
		); err != nil {
			return fmt.Errorf("AT receipt save failed: %v: %w", err, asynq.SkipRetry)
		}

		return nil
	}
}
