package computeengine

import "fmt"

var simpleQueue []Job

// AddJob adds a new AI job to the queue.
func AddJob(job Job) {
	simpleQueue = append(simpleQueue, job)
	fmt.Printf("📝 Job %s added to queue (%s)\n", job.ID, job.Model)
}

// GetNextJob retrieves the next available job.
func GetNextJob() *Job {
	if len(simpleQueue) == 0 {
		return nil
	}

	next := simpleQueue[0]
	simpleQueue = simpleQueue[1:]
	return &next
}
