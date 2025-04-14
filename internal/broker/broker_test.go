package broker_test

import (
	"testing"

	"github.com/kvn-alcantara/tiny-dispatcher/internal/broker"
	"github.com/kvn-alcantara/tiny-dispatcher/internal/job"
)

func TestEnqueue(t *testing.T) {
	queue := make(chan job.Job, 1)
	b := broker.New(queue)
	j := job.New("test-job", nil)

	b.Enqueue(*j)

	select {
	case queuedJob := <-queue:
		if queuedJob.ID != j.ID {
			t.Errorf("Expected %v, got %v", j, queuedJob)
		}
	default:
		t.Error("Expected job to be in the queue, but it was not")
	}
}

func TestNext(t *testing.T) {
	queue := make(chan job.Job, 1)
	b := broker.New(queue)
	j := job.New("test-job", nil)

	b.Enqueue(*j)

	nextJob := b.Next()

	if nextJob.ID != j.ID {
		t.Errorf("Expected %v, got %v", j, nextJob)
	}
}

func TestNextEmptyQueue(t *testing.T) {
	queue := make(chan job.Job, 1)
	b := broker.New(queue)

	nextJob := b.Next()

	if nextJob.ID != "" {
		t.Errorf("Expected nil, got %v", nextJob)
	}
}
