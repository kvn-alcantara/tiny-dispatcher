package job_test

import (
	"testing"

	"github.com/kvn-alcantara/tiny-dispatcher/internal/job"
)

func TestGenerateUniqueId(t *testing.T) {
	j1 := job.New("First", nil)
	j2 := job.New("Second", nil)

	if j1.ID == j2.ID {
		t.Errorf("IDs not unique: %q, %q", j1.ID, j2.ID)
	}
}

func TestInitialStatus(t *testing.T) {
	j := job.New("test", nil)

	if j.Status != job.Pending {
		t.Errorf("Initial job status should be %q, got %q", job.Pending, j.Status)
	}
}
