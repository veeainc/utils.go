# sync

```go
import "github.com/veeainc/utils.go/sync"
```

A semaphore implementation using Go a go buffered channel.

It allow to restrict a resource access to N user at the same time when using Goroutines.

It works a bit like context managers in Python3.

## Example

```go
inFlightMsg = sync.NewSemaphore(4)  // max 4 messages at the same time

// somewhere else with goroutines

inFlightMsg.With(func() (any, error) {
	// you can enter here only 4 times at the same time
    // other call to .With will be put on hold
}

```