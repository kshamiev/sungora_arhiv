package task

import (
	"context"
	"time"

	"sample/lib/logger"
	"sample/lib/storage"
	"sample/lib/storage/stpg"
)

const TaskOnlineOffName = "TaskUserOnlineOff"

// TaskOnlineOff Обновление онлайн статуса ушедших пользователей
type TaskOnlineOff struct {
	st storage.Face
}

func NewTaskOnlineOff() *TaskOnlineOff {
	return &TaskOnlineOff{st: stpg.Gist()}
}

func (task *TaskOnlineOff) Action(ctx context.Context) error {
	lg := logger.Get(ctx)
	lg.Info("Its's Work")
	return nil
}

func (task *TaskOnlineOff) WaitFor() time.Duration {
	return time.Minute
}

func (task *TaskOnlineOff) Name() string {
	return TaskOnlineOffName
}
