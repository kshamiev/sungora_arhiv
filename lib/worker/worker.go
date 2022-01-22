package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"sungora/lib/logger"
	"sungora/lib/response"
	"sungora/lib/typ"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc/metadata"
)

type Task interface {
	Name() string                     // информация о задаче
	Action(ctx context.Context) error // выполняемая задача
	WaitFor() time.Duration           // время до следующего запуска
}

type scheduler struct {
	wg       sync.WaitGroup       // контроль задач без расписания
	pullWork map[string]chan bool // включенные задачи по расписанию
	pull     []Task               // пулл всех задач
}

var instance *scheduler
var mu sync.RWMutex

func Init() {
	if instance == nil {
		mu.Lock()
		if instance == nil {
			instance = &scheduler{
				pullWork: make(map[string]chan bool),
			}
		}
		mu.Unlock()
	}
}

// //// control method

// Add добавить в scheduler задачу по расписанию
func Add(w Task) {
	instance.pull = append(instance.pull, w)
}

// Start включить конкретную задачу по расписанию
func Start(name string) {
	if _, ok := instance.pullWork[name]; ok {
		return
	}
	for i := range instance.pull {
		if instance.pull[i].Name() == name {
			instance.pullWork[name] = make(chan bool)
			go runScheduler(instance.pull[i], instance.pullWork[name])
			return
		}
	}
}

func AddStart(w Task) {
	Add(w)
	Start(w.Name())
}

// Run запустить задачу на выполенине
func Run(w Task) {
	instance.wg.Add(1)
	go func() {
		action(w)
		instance.wg.Done()
	}()
}

func Stop(name string) {
	if _, ok := instance.pullWork[name]; !ok {
		return
	}
	instance.pullWork[name] <- true
	<-instance.pullWork[name]
	delete(instance.pullWork, name)
}

func CloseWait() {
	for k := range instance.pullWork {
		instance.pullWork[k] <- true
	}
	for k := range instance.pullWork {
		<-instance.pullWork[k]
		delete(instance.pullWork, k)
	}
	instance.wg.Wait()
}

// //// support method

func GetTasks() map[string]Task {
	res := make(map[string]Task)
	for i := range instance.pull {
		res[instance.pull[i].Name()] = instance.pull[i]
	}
	return res
}

func runScheduler(task Task, ch chan bool) {
	for {
		waitFor := task.WaitFor()
		select {
		case <-time.After(waitFor):
			action(task)
		case <-ch:
			ch <- true
			return
		}
	}
}

// action выполнение задачи
func action(task Task) {
	requestID := typ.UUIDNew().StringShort()
	ctx := context.Background()
	lg := logger.Gist(ctx).WithField(response.LogTraceID, requestID)

	ctx = context.WithValue(ctx, response.CtxTraceID, requestID)
	ctx = logger.WithLogger(ctx, lg)
	ctx = boil.WithDebugWriter(ctx, lg.Writer())

	m := make(map[string]string)
	m[response.LogTraceID] = requestID
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(m))

	defer func() {
		if rvr := recover(); rvr != nil {
			lg.Error(fmt.Sprintf("panic: %+v", rvr))
		}
	}()

	if err := task.Action(ctx); err != nil {
		if e, ok := err.(response.Error); ok {
			lg.Error(e.Error())
			for _, t := range e.Trace() {
				lg.Trace(t)
			}
		} else {
			lg.Error(e)
		}
	}
}
