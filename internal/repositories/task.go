package repositories

import (
    "github.com/jmoiron/sqlx"
    entity "todo-app/internal/entities"
)

type Task struct {
    db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *Task  {
    return &Task{db: db}
}

func (t *Task) Create(task entity.Task)  {

}