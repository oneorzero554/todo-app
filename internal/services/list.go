package services

import (
    entity "todo-app/internal/entities"
    "todo-app/internal/repositories"
)

type List struct {
    repo *repositories.List
}

func NewListService(repo *repositories.List) *List  {
    return &List{repo: repo}
}

func (t *Task) Create(list entity.List) (int, error) {
    return t.repo.Create(list)
}