package models

import (
	"jira/models/task"
	"time"
)

type SprintStatus int

const (
	SSOPEN SprintStatus = iota
	SSINPROGRESS
	SSCOMPLETED
)

type Sprint struct {
	Tasks     []*task.ITask
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Status    SprintStatus
}
