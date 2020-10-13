package counter

import "sync"

// ThreadCounter is a...
type ThreadCounter struct {
	threads int
	count   int
	mux     *sync.Mutex
}

// New creates a new instance of a ThreadCounter
func New(threads int) ThreadCounter {
	return ThreadCounter{count: 0, threads: threads, mux: &sync.Mutex{}}
}

// Completed marks one thread as completed
func (mc *ThreadCounter) Completed() {
	mc.mux.Lock()
	mc.count++
	mc.mux.Unlock()
}

// AreCompleted returns `true` if all the threads are marked as completed
func (mc *ThreadCounter) AreCompleted() bool {
	mc.mux.Lock()
	// Lock so only one goroutine at a time can access the map counter
	defer mc.mux.Unlock()
	return mc.count >= mc.threads
}
