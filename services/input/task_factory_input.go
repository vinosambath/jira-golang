package input

import (
	"jira/models/task"
	"time"
)

type TaskFactoryCreateInput struct {
	Title    string
	Assignee string
	TaskType task.TaskType
	DueDate  time.Time
	Status   task.TaskStatus

	// optional
	StorySummary   string
	FeatureSummary string
	BugSeverity    task.BugTaskSeverity
}

func NewTaskFactoryCreateInput(title string, assignee string, taskType task.TaskType, dueDate time.Time) *TaskFactoryCreateInput {
	return &TaskFactoryCreateInput{
		Title:    title,
		Assignee: assignee,
		TaskType: taskType,
		DueDate:  dueDate,
	}
}

func (i *TaskFactoryCreateInput) SetStorySummary(storySummary string) {
	i.StorySummary = storySummary
}

func (i *TaskFactoryCreateInput) SetFeatureSummary(featureSummary string) {
	i.FeatureSummary = featureSummary
}

func (i *TaskFactoryCreateInput) SetBugSeverity(bugSeverity task.BugTaskSeverity) {
	i.BugSeverity = bugSeverity
}
