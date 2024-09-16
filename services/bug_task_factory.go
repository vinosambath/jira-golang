package services

import (
	"jira/models/task"
	"jira/services/input"
)

type BugTaskFactory struct{}

func (b *BugTaskFactory) CreateTask(input input.TaskFactoryCreateInput) (task.ITask, error) {
	// validate input

	return &task.BugTask{
		Task: task.Task{
			Title:    input.Title,
			Assignee: input.Assignee,
			DueDate:  input.DueDate,
			Type:     input.TaskType,
		},
		Severity: input.BugSeverity,
		Status:   input.Status,
	}, nil
}

func (b *BugTaskFactory) UpdateStatus(task task.ITask, newStatus task.TaskStatus) (task.ITask, error) {
	// validate
	task.SetStatus(newStatus)
	return task, nil
}
