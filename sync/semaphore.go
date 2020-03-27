package sync

// A semaphore implementation using a buffered channel
type Semaphore struct {
	sem chan int // a buffered channel
}

func NewSemaphore(count int) *Semaphore {
	return &Semaphore{
		sem: make(chan int, count),
	}
}

func (s *Semaphore) Acquire() {
	s.sem <- 1
}

func (s *Semaphore) Release() {
	<- s.sem
}

// Exec the handler inside the semaphore
func (s *Semaphore) With(handler func() (interface{}, error)) (interface{}, error) {
	s.Acquire()
	defer s.Release()
	return handler()
}
