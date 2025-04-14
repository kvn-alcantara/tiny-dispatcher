package job

import "github.com/google/uuid"

// Status represents the status of a job
type Status string

// Represents the possible statuses of a job
const (
	Pending  Status = "PENDING"
	Running  Status = "RUNNING"
	Success  Status = "SUCCESS"
	Failed   Status = "FAILED"
)

// Job represents a unit of work to be executed
type Job struct {
	ID      string
	Command string
	Args    []string
	Status  Status
	Retries int
}

// Executor is an interface that defines the method to execute a job
type Executor interface {
	Run(job Job) (output string, err error)
}

// New creates a new job with the given command and arguments
func New(command string, args []string) *Job {
	return &Job{
		ID:      uuid.NewString(),
		Command: command,
		Args:    args,
		Status:  Pending,
		Retries: 0,
	}
}
