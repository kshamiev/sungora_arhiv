package data

import (
	"context"
	"time"

	"sungora/lib/storage"
)

const TaskStorageClearName = "TaskStorageClear"

type TaskStorageClear struct {
	st storage.Face
}

func NewTaskStorageClear(st storage.Face) *TaskStorageClear { return &TaskStorageClear{st: st} }

func (task *TaskStorageClear) Action(ctx context.Context) error {
	stM := NewModel(task.st, "")
	return stM.RemoveNotConfirm(ctx)
}

func (task *TaskStorageClear) WaitFor() time.Duration {
	return time.Hour * 24
}

func (task *TaskStorageClear) Name() string {
	return TaskStorageClearName
}
