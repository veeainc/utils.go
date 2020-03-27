package tasks

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTask_WaitStop(t *testing.T) {
	task := NewTask("read", func(stop StopChannel, args ...interface{}) {
		for {
			select {
			default:
				time.Sleep(500 * time.Millisecond)
				fmt.Println("working")
			case <-stop:
				return
			}
		}
	})

	err := task.Run()
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
	err = task.WaitStop()
	assert.NoError(t, err)

	err = task.Run()
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
	err = task.WaitStop()
	assert.NoError(t, err)
}

func TestTask_WaitDone(t *testing.T) {
	task := NewTask("read", func(stop StopChannel, args ...interface{}) {
		count := 2
		for {
			select {
			default:
				if count == 0 {
					return
				}
				time.Sleep(500 * time.Millisecond)
				fmt.Println("working")
				count--
			case <-stop:
				return
			}
		}
	})

	err := task.Run()
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
	err = task.WaitDone()
	assert.NoError(t, err)

	err = task.Run()
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
	err = task.WaitDone()
	assert.NoError(t, err)
}