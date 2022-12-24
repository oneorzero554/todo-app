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

func (l *List) GetAll() ([]entity.List, error) {
    return l.repo.GetAll()
}

func (l *List) GetById(id int) (entity.List, error) {
    return l.repo.GetById(id)
}

func (l *List) Delete(id int) error {
    return l.repo.Delete(id)
}

func (l *List) Create(list *entity.List) (*entity.List, error) {
    return l.repo.Create(list)
}

func (l *List) Update(list *entity.List) (*entity.List, error) {
    return l.repo.Update(list)
}