package services

import (
	"jira/models/task"
	"jira/services/input"
)

type StoryTaskFactory struct{}

func (s *StoryTaskFactory) CreateTask(input input.TaskFactoryCreateInput) (task.ITask, error) {
	// validation

	return &task.StoryTask{
		Task: task.Task{
			Title:    input.Title,
			Assignee: input.Assignee,
			DueDate:  input.DueDate,
			Type:     input.TaskType,
		},
		StorySummary: input.StorySummary,
		Status:       input.Status,
	}, nil
}

func (s *StoryTaskFactory) UpdateStatus(task task.ITask, newStatus task.TaskStatus) (task.ITask, error) {
	task.SetStatus(newStatus)
	return task, nil
}
