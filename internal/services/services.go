package services

import "todo-app/internal/repositories"

type Services struct {
    *List
    *Task
}

func InitServices(repos *repositories.Repositories) *Services  {
    return &Services{
        List: NewListService(repos.List),
        Task: NewTaskService(repos.Task),
    }
}