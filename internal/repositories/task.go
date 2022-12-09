package repositories

import (
    "github.com/jmoiron/sqlx"
)

type Task struct {
    db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *Task  {
    return &Task{db: db}
}