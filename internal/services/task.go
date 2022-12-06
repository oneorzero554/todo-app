package services

import (
    "todo-app/internal/repositories"
)

type Task struct {
    repo *repositories.Task
}

func NewTaskService(repo *repositories.Task) *Task  {
    return &Task{repo: repo}
}
