package services

import (
    "todo-app/internal/repositories"
)

type List struct {
    repo *repositories.List
}

func NewListService(repo *repositories.List) *List  {
    return &List{repo: repo}
}