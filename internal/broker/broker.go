package broker

import (
	"github.com/kvn-alcantara/tiny-dispatcher/internal/job"
)

// Broker is a simple job queue
type Broker struct {
	queue chan job.Job
}

// New creates a new Broker with the provided job queue
func New(queue chan job.Job) *Broker {
	return &Broker{
		queue: queue,
	}
}

// Enqueue adds a job to the queue
func (b *Broker) Enqueue(j job.Job) {
	b.queue <- j
}

// Next returns the next job from the queue.
func (b *Broker) Next() (job.Job) {
	select {
	case j := <-b.queue:
		return j
	default:
		return job.Job{}
	}
}
