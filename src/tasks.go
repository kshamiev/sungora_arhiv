package src

import (
	"sungora/lib/worker"
	"sungora/src/task"
)

// initWorkers инициализация фоновых задач
func initWorkers() {
	worker.AddStart(task.NewUserOnlineOff())
}
