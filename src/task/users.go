package task

import (
	"context"
	"time"

	"sungora/lib/logger"
	"sungora/lib/storage"
	"sungora/lib/storage/stpg"
)

const UserOnlineOffName = "UserOnlineOff"

// Обновление онлайн статуса ушедших пользователей
type UserOnlineOff struct {
	storage.Face
}

func NewUserOnlineOff() *UserOnlineOff { return &UserOnlineOff{&stpg.Storage{}} }

func (task *UserOnlineOff) Action(ctx context.Context) error {
	lg := logger.GetLogger(ctx)
	lg.Info("Its's Work")
	return nil
}

func (task *UserOnlineOff) WaitFor() time.Duration {
	return time.Minute
}

func (task *UserOnlineOff) Name() string {
	return UserOnlineOffName
}
