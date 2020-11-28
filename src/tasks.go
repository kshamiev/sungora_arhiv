package src

import (
	"sungora/lib/worker"
	"sungora/src/task"
)

// инициализация фоновых задач
func initWorkers() {
	worker.AddStart(task.NewUserOnlineOff())
}
