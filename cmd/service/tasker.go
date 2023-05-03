package main

import (
	"context"
	"time"

	"github.com/grassrootseconomics/cic-custodial/pkg/redis"
	"github.com/grassrootseconomics/cic-notify/internal/notify"
	"github.com/grassrootseconomics/cic-notify/internal/tasker"
	"github.com/grassrootseconomics/cic-notify/internal/tasker/task"
	"github.com/hibiken/asynq"
)

const (
	fixedRetryCount  = 25
	fixedRetryPeriod = time.Second * 1
)

// Load tasker handlers, injecting any necessary handler dependencies from the system container.
func initTasker(notifyContainer *notify.Notify, redisPool *redis.RedisPool) *tasker.TaskerServer {
	taskerServerOpts := tasker.TaskerServerOpts{
		Concurrency:  ko.MustInt("asynq.worker_count"),
		Logg:         lo,
		LogLevel:     asynq.InfoLevel,
		RedisPool:    redisPool,
		RetryHandler: retryHandler,
	}

	taskerServer := tasker.NewTaskerServer(taskerServerOpts)

	taskerServer.RegisterMiddlewareStack([]asynq.MiddlewareFunc{
		observibilityMiddleware(),
	})

	taskerServer.RegisterHandlers(tasker.AtPushTask, task.AtPushProcessor(notifyContainer))
	taskerServer.RegisterHandlers(tasker.TgPushTask, task.TgPushProcessor(notifyContainer))
	taskerServer.RegisterHandlers(tasker.PrepareMessage, task.PrepareMsgProcessor(notifyContainer))
	taskerServer.RegisterHandlers(tasker.ProcessFailedMsgTask, task.FailedMsgProcessor(notifyContainer))
	taskerServer.RegisterHandlers(tasker.ProcessSuccessReceivedMsgTask, task.SuccessReceivedMsgProcessor(notifyContainer))
	taskerServer.RegisterHandlers(tasker.ProcessSuccessSentMsgTask, task.SuccessSentMsgProcessor(notifyContainer))

	return taskerServer
}

func retryHandler(count int, err error, task *asynq.Task) time.Duration {
	if count < fixedRetryCount {
		return fixedRetryPeriod
	} else {
		return asynq.DefaultRetryDelayFunc(count, err, task)
	}
}

func observibilityMiddleware() asynq.MiddlewareFunc {
	return func(handler asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			taskId, _ := asynq.GetTaskID(ctx)

			err := handler.ProcessTask(ctx, task)
			if err != nil {
				lo.Error("tasker: handler error", "err", err, "task_type", task.Type(), "task_id", taskId)
			} else if asynq.IsPanicError(err) {
				lo.Error("tasker: handler panic", "err", err, "task_type", task.Type(), "task_id", taskId)
			} else {
				lo.Info("tasker: process task", "task_type", task.Type(), "task_id", taskId)
			}

			return err
		})
	}
}
