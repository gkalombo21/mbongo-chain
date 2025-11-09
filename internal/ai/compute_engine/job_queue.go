package computeengine

import (
	"errors"
	"sync"
)

// JobQueue defines the behaviour required to buffer incoming jobs before execution.
type JobQueue interface {
	Enqueue(Job) error
	Dequeue() (Job, error)
	Len() int
}

// memoryQueue implements JobQueue using a mutex-protected slice.
type memoryQueue struct {
	mu   sync.Mutex
	jobs []Job
	cap  int
}

// NewMemoryQueue constructs a new in-memory queue with the optional capacity limit.
func NewMemoryQueue(capacity int) JobQueue {
	return &memoryQueue{jobs: make([]Job, 0), cap: capacity}
}

// Enqueue adds a job to the queue respecting capacity limits when configured.
func (q *memoryQueue) Enqueue(job Job) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.cap > 0 && len(q.jobs) >= q.cap {
		return errors.New("job queue is at capacity")
	}

	q.jobs = append(q.jobs, job)
	return nil
}

// Dequeue removes the next job in FIFO order.
func (q *memoryQueue) Dequeue() (Job, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.jobs) == 0 {
		return Job{}, errors.New("job queue is empty")
	}

	job := q.jobs[0]
	q.jobs = q.jobs[1:]
	return job, nil
}

// Len returns the current number of queued jobs.
func (q *memoryQueue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.jobs)
}
