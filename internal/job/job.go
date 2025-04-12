package job

// Status represents the status of a job
type Status int

// Represents the possible statuses of a job
const (
	Queued  Status = iota
	Running
	Success
	Failed
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
