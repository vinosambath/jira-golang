package services

import (
	"jira/models/task"
	"jira/services/input"
)

type TaskFactory interface {
	CreateTask(input input.TaskFactoryCreateInput) (task.ITask, error)
	UpdateStatus(task task.ITask, newStatus task.TaskStatus) (task.ITask, error)
	//UpdateAssignee(task task.ITask, assignee string) error
	//TaskAssignedToUserByTaskType(user user.User, taskType task.TaskType) ([]*task.ITask, error)
}
