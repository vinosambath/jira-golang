package task

import "fmt"

type BugTaskSeverity int

const (
	BTSP0 BugTaskSeverity = iota
	BTSP1
	BTSP2
)

var allowedStatusChanges = map[TaskStatus][]TaskStatus{
	TSOPEN:        {TSIN_PROGRESS},
	TSIN_PROGRESS: {TSTESTING},
}

type BugTask struct {
	Task
	Severity BugTaskSeverity
	Status   TaskStatus
}

func (b *BugTask) SetStatus(s TaskStatus) error {
	currentStatus := b.GetStatus()

	if allowedStatusChanges[currentStatus] == nil {
		return fmt.Errorf("error status change not allowed")
	}

	b.Status = s
	return nil
}

func (b *BugTask) GetStatus() TaskStatus {
	return b.Status
}
