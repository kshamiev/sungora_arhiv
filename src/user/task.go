package user

import (
	"context"
	"time"

	"sungora/lib/logger"
	"sungora/lib/storage/pgsql"
)

const TaskOnlineOffName = "TaskUserOnlineOff"

// TaskOnlineOff Обновление онлайн статуса ушедших пользователей
type TaskOnlineOff struct {
	st *pgsql.Storage
}

func NewTaskOnlineOff(st *pgsql.Storage) *TaskOnlineOff {
	return &TaskOnlineOff{st: st}
}

func (task *TaskOnlineOff) Action(ctx context.Context) error {
	lg := logger.GetLogger(ctx)
	lg.Info("Its's Work")
	return nil
}

func (task *TaskOnlineOff) WaitFor() time.Duration {
	return time.Minute
}

func (task *TaskOnlineOff) Name() string {
	return TaskOnlineOffName
}
