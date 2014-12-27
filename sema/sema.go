package sema

type Semaphore struct {
	ch chan struct{}
}

// Create a new Semaphore, setting max concurrent jobs to count
func New(count int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, count)}
}

// Create a new Semaphore, setting max concurrent jobs to 1
func Mutex() *Semaphore { return New(1) }

// Aquire a job, blocks until one slot is free
func (s *Semaphore) Acquire() *Semaphore { s.ch <- struct{}{}; return s }

// Release one job
func (s *Semaphore) Release() { <-s.ch }

// Try to lock, return true if locked
func (s *Semaphore) Try() bool {
	select {
	case s.ch <- struct{}{}:
		return true
	default:
	}
	return false
}
