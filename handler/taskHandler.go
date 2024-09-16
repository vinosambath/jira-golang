package handler

import (
	"jira/models/task"
	"jira/services"
	"jira/services/input"
)

type TaskHandler interface {
	CreateTask(input input.TaskFactoryCreateInput) (task.ITask, error)
	UpdateStatus(task task.ITask, newStatus task.TaskStatus) (task.ITask, error)
	TaskAssignedToUserByTaskType(user string, taskType task.TaskType) ([]*task.ITask, error)
}

type taskHandler struct {
	taskBusinessService services.TaskBusinessService
}

func GettaskHandler() TaskHandler {
	return &taskHandler{
		taskBusinessService: services.NewTaskBusinessService(),
	}
}

func (h *taskHandler) CreateTask(input input.TaskFactoryCreateInput) (task.ITask, error) {
	return h.taskBusinessService.CreateTask(input)
}

func (h *taskHandler) UpdateStatus(task task.ITask, newStatus task.TaskStatus) (task.ITask, error) {
	return h.taskBusinessService.UpdateStatus(task, newStatus)
}

func (h *taskHandler) TaskAssignedToUserByTaskType(user string, taskType task.TaskType) ([]*task.ITask, error) {
	return h.taskBusinessService.TaskAssignedToUserByTaskType(user, taskType)
}
