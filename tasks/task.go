package tasks

import (
	"errors"
	"sync"

	"github.com/veeainc/utils.go/logging"
)

var _taskLog = logging.GetNamedLogger()

// A channel used to receive a stop signal
type StopChannel = chan struct{}
type taskFn = func(stop StopChannel, args ...interface{})

// An utility to create stoppable goroutine task easily.
type Task struct {
	name string         // for logging
	fn   taskFn         // Func to execute
	stop StopChannel    // close this channel to stop
	wg   sync.WaitGroup // wait goroutine finished
}

func NewTask(name string, task taskFn) *Task {
	return &Task{
		name: name,
		fn:   task,
		stop: nil,
	}
}

func (t *Task) Run(args ...interface{}) error {
	if t.stop != nil {
		return errors.New("task " + t.name + " is already running")
	}

	t.stop = make(StopChannel)
	t.wg = sync.WaitGroup{} // reset
	t.wg.Add(1)

	_taskLog.Tracef("starting task %s", t.name)
	go func() {
		t.fn(t.stop, args...) // run in goroutine
		_taskLog.Tracef("task %s is done", t.name)
		t.wg.Done() // after the execution, set as done
	}()

	return nil
}

func (t *Task) IsDone() bool {
	return t.stop == nil
}

func (t *Task) WaitStop() error {
	if t.stop == nil {
		return errors.New("not started")
	}

	_taskLog.Tracef("request to stop task %s, waiting...", t.name)
	close(t.stop) // close channel, the fn must stop when possible
	t.stop = nil
	t.wg.Wait()
	return nil
}

func (t *Task) WaitDone() error {
	if t.stop == nil {
		return errors.New("not started")
	}

	t.wg.Wait()
	close(t.stop) // close channel, the fn must stop when possible
	t.stop = nil
	return nil
}
