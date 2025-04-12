package broker

import (
	"github.com/kvn-alcantara/tiny-dispatcher/internal/job"
)

// Broker is a simple job queue
type Broker struct {
	queue chan job.Job
}

// Enqueue adds a job to the queue
func (b *Broker) Enqueue(job job.Job) {
	b.queue <- job
}

// Next returns the next job from the queue
func (b *Broker) Next() job.Job {
	return <-b.queue
}
