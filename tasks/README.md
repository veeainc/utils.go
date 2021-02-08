# tasks

```go
import "github.com/veeainc/utils.go/tasks"
```

The `Task` structure facilitates the management of Goroutine.

## Example

Creating a Task to read on a serial port in background:

```go
// task name is used in the logs
func readingTaskWork(stop utils.StopChannel, args ...interface{}) {
	buffer := make([]byte, uartBufferSize)
	
	for { // infinite loop
		select {
		default:
			if n, err := g.port.Read(buffer); err != nil {
				uartLog.Panic(err)
			} else {
				data := buffer[:n]
				g.dataReceived(data)
			}
		case <-stop: // we enter here after a call to readingTask.WaitStop()
			return
		}
	}
}

readingTask = tasks.NewTask("my-reading-task", readingTaskWork)
readingTask.Run() // you can pass args here to be passed in your handler
readingTask.IsDone() // false
readingTask.WaitDone()
readingTask.IsDone() // true

// readingTask.WaitStop()

```