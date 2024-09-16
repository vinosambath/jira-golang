package services

import (
	"jira/models/task"
	"jira/services/input"
)

type FeatureTaskFactory struct{}

func (f *FeatureTaskFactory) CreateTask(input input.TaskFactoryCreateInput) (task.ITask, error) {
	// validation

	return &task.FeatureTask{
		Task: task.Task{
			Title:    input.Title,
			Assignee: input.Assignee,
			DueDate:  input.DueDate,
			Type:     input.TaskType,
		},
		FeatureSummary: input.FeatureSummary,
		// Impact: input.impac todo
		Status: input.Status,
	}, nil
}

func (f *FeatureTaskFactory) UpdateStatus(task task.ITask, newStatus task.TaskStatus) (task.ITask, error) {
	task.SetStatus(newStatus)
	return task, nil
}
