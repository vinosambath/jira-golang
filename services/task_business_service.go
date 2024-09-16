package services

import (
	"fmt"
	"jira/models/task"
	"jira/services/input"
	"sync"
)

type TaskBusinessService interface {
	CreateTask(input input.TaskFactoryCreateInput) (task.ITask, error)
	UpdateStatus(task task.ITask, newStatus task.TaskStatus) (task.ITask, error)
	TaskAssignedToUserByTaskType(user string, taskType task.TaskType) ([]*task.ITask, error)
}

type taskBusinessService struct {
	tasksByUserByTaskType map[string]map[task.TaskType][]*task.ITask
	FeatureTaskFactory
	BugTaskFactory
	StoryTaskFactory
}

var TaskBusinessServiceInstance TaskBusinessService
var TaskBusinessServiceInstanceSingleton = sync.Once{}

func NewTaskBusinessService() TaskBusinessService {
	TaskBusinessServiceInstanceSingleton.Do(func() {
		TaskBusinessServiceInstance = &taskBusinessService{
			tasksByUserByTaskType: make(map[string]map[task.TaskType][]*task.ITask),
			FeatureTaskFactory:    FeatureTaskFactory{},
			BugTaskFactory:        BugTaskFactory{},
			StoryTaskFactory:      StoryTaskFactory{},
		}
	})

	return TaskBusinessServiceInstance
}

func (s *taskBusinessService) CreateTask(input input.TaskFactoryCreateInput) (task.ITask, error) {
	factory := s.GetTaskFactoryFromTaskType(input.TaskType)
	createdTask, err := factory.CreateTask(input)

	if err != nil {
		return nil, err
	}

	if taskByUser, ok := s.tasksByUserByTaskType[input.Assignee]; ok {
		if taskByUserByTaskType, ok := taskByUser[input.TaskType]; ok {
			s.tasksByUserByTaskType[input.Assignee][input.TaskType] = append(taskByUserByTaskType, &createdTask)
		} else {
			s.tasksByUserByTaskType[input.Assignee][input.TaskType] = []*task.ITask{&createdTask}
		}
	} else {
		s.tasksByUserByTaskType[input.Assignee] = make(map[task.TaskType][]*task.ITask)
		s.tasksByUserByTaskType[input.Assignee][input.TaskType] = []*task.ITask{&createdTask}
	}

	fmt.Println(s.tasksByUserByTaskType)
	return createdTask, nil
}

func (s *taskBusinessService) UpdateStatus(task task.ITask, newStatus task.TaskStatus) (task.ITask, error) {
	factory := s.GetTaskFactoryFromTaskType(task.GetType())

	updatedTask, err := factory.UpdateStatus(task, newStatus)
	return updatedTask, err
}

func (s *taskBusinessService) TaskAssignedToUserByTaskType(user string, taskType task.TaskType) ([]*task.ITask, error) {
	tasks := s.tasksByUserByTaskType[user][taskType]
	return tasks, nil
}

func (s *taskBusinessService) GetTaskFactoryFromTaskType(taskType task.TaskType) TaskFactory {
	switch taskType {
	case task.TTFEATURE:
		return &s.FeatureTaskFactory
	case task.TTSTORY:
		return &s.StoryTaskFactory
	case task.TTBUG:
		return &s.BugTaskFactory
	}
	return nil
}
