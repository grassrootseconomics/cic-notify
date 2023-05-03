package tasker

import "encoding/json"

type (
	QueueName string
	TaskName  string
)

type Task struct {
	Id      string          `json:"id"`
	Payload json.RawMessage `json:"payload"`
}

const (
	// Push action tasks.
	AtPushTask TaskName = "push:at"
	TgPushTask TaskName = "push:tg"
	// Process/Prepare message payload tasks.
	PrepareMessage                TaskName = "msg:prepare"
	ProcessFailedMsgTask          TaskName = "msg:failed"
	ProcessSuccessReceivedMsgTask TaskName = "msg:success_received"
	ProcessSuccessSentMsgTask     TaskName = "msg:success_sent"
)

const (
	HighPriority    QueueName = "high_priority"
	DefaultPriority QueueName = "default_priority"
)
