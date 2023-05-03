package sub

import (
	"context"

	"github.com/grassrootseconomics/cic-notify/internal/tasker"
	"github.com/nats-io/nats.go"
)

func (s *Sub) processEventHandler(ctx context.Context, msg *nats.Msg) error {
	_, err := s.taskerClient.CreateTask(
		ctx,
		tasker.PrepareMessage,
		tasker.DefaultPriority,
		&tasker.Task{
			Payload: msg.Data,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
