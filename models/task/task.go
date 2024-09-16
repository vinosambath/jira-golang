package task

import "time"

type TaskType int

const (
	TTFEATURE TaskType = iota
	TTBUG
	TTSTORY
)

type TaskStatus int

const (
	TSOPEN TaskStatus = iota
	TSIN_PROGRESS
	TSTESTING
	TSDEPLOYED
	TSFIXED
	TSCOMPLETED
)

type ITask interface {
	GetTitle() string
	SetTitle(title string)
	GetAssignee() string
	SetAssignee(assignee string)
	GetType() TaskType
	SetType(t TaskType)
	GetDueDate() time.Time
	SetDueDate(dueDate time.Time)
	SetStatus(s TaskStatus) error
	GetStatus() TaskStatus
}

type Task struct {
	Title    string
	Assignee string
	Type     TaskType
	DueDate  time.Time
}

func (t *Task) GetTitle() string {
	return t.Title
}

func (t *Task) SetTitle(title string) {
	t.Title = title
}

func (t *Task) GetAssignee() string {
	return t.Assignee
}

func (t *Task) SetAssignee(assignee string) {
	t.Assignee = assignee
}

func (t *Task) GetType() TaskType {
	return t.Type
}

func (t *Task) SetType(taskType TaskType) {
	t.Type = taskType
}

func (t *Task) GetDueDate() time.Time {
	return t.DueDate
}

func (t *Task) SetDueDate(dueDate time.Time) {
	t.DueDate = dueDate
}
