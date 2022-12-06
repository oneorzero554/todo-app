package repositories

import (
    "github.com/jmoiron/sqlx"
    entity "todo-app/internal/entities"
)

type List struct {
    db *sqlx.DB
}

func NewListRepository(db *sqlx.DB) *List  {
    return &List{db: db}
}

func (l *List) Create(list entity.List)  {
    
}