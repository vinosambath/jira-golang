package main

import (
	"fmt"
	"jira/handler"
	"jira/models/task"
	"jira/services/input"
	"time"
)

/*
---------
Models:
	User -> id, name
	Task -> Title, Assignee, Type, Due Date
		Type of Tasks ->
			Feature - Feature summary, impact, status
			Bug - Severity, status
			Story - Summary, sub track, allowed status
				Title, status, parent task
	Sprint -> Collection of tasks, name, start date, end date, status


struct Task {
	title
	assignee
	type
	due date
}

Feature {
	Title,
	Feature summary
}

----------
Service:
	Task
		Create, change status, assignee, get tasks by user and task type
	Sub track
		add subtrack to existing story, change status
	Sprint
		Create, Remove task, Add task, Display sprint snapshot, start or end a sprint

---------
*/

func main() {
	taskHandler := handler.GettaskHandler()
	_, _ = taskHandler.CreateTask(input.TaskFactoryCreateInput{
		Title:          "task1",
		Assignee:       "vinsam",
		TaskType:       task.TTFEATURE,
		DueDate:        time.Now(),
		Status:         task.TSOPEN,
		FeatureSummary: "sometext",
	})

	_, _ = taskHandler.CreateTask(input.TaskFactoryCreateInput{
		Title:    "task2",
		Assignee: "vinsam",
		TaskType: task.TTBUG,
		DueDate:  time.Now(),
		Status:   task.TSOPEN,
	})

	_, _ = taskHandler.CreateTask(input.TaskFactoryCreateInput{
		Title:          "task3",
		Assignee:       "anotheruser",
		TaskType:       task.TTBUG,
		DueDate:        time.Now(),
		Status:         task.TSOPEN,
		FeatureSummary: "sometext",
	})

	tasks, _ := taskHandler.TaskAssignedToUserByTaskType("vinsam", task.TTFEATURE)
	fmt.Println(tasks)
}
