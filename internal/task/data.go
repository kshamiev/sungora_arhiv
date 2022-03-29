package task

import (
	"context"
	"time"

	"sample/internal/model"
	"sample/lib/storage"
	"sample/lib/storage/stpg"
)

const TaskStorageClearName = "TaskStorageClear"

type TaskStorageClear struct {
	st storage.Face
}

func NewTaskStorageClear() *TaskStorageClear {
	return &TaskStorageClear{st: stpg.Gist()}
}

func (task *TaskStorageClear) Action(ctx context.Context) error {
	stM := model.NewData(task.st, "")
	return stM.RemoveNotConfirm(ctx)
}

func (task *TaskStorageClear) WaitFor() time.Duration {
	return time.Hour * 24
}

func (task *TaskStorageClear) Name() string {
	return TaskStorageClearName
}
