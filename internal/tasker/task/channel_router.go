package task

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/grassrootseconomics/cic-notify/internal/graphql"
	"github.com/grassrootseconomics/cic-notify/internal/tasker"
)

// routeChannel checks the interface type from the graph and routes the task to the appropriate registered handler.
func routeMessage(
	ctx context.Context,
	taskerClient *tasker.TaskerClient,
	channelType graphql.Interface_type_enum,
	msgPayload string,
	channelIdentifier string,
) error {
	switch channelType {
	case graphql.Interface_type_enumUssd:
		atPayload, err := json.Marshal(atPayload{
			RecepientPhone: channelIdentifier,
			Message:        msgPayload,
		})
		if err != nil {
			return err
		}

		if _, err := taskerClient.CreateTask(
			ctx,
			tasker.AtPushTask,
			tasker.HighPriority,
			&tasker.Task{
				Payload: atPayload,
			},
		); err != nil {
			return err
		}
	case graphql.Interface_type_enumTelegram:
		chatId, err := strconv.ParseInt(channelIdentifier, 10, 64)
		if err != nil {
			return err
		}

		tgPayload, err := json.Marshal(tgPayload{
			ChatId:  chatId,
			Message: msgPayload,
		})
		if err != nil {
			return err
		}

		if _, err := taskerClient.CreateTask(
			ctx,
			tasker.TgPushTask,
			tasker.HighPriority,
			&tasker.Task{
				Payload: tgPayload,
			},
		); err != nil {
			return err
		}
	}

	return nil
}
