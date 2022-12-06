package repositories

import "github.com/jmoiron/sqlx"

type Repositories struct {
    *List
    *Task
}

func InitRepostitories(db *sqlx.DB) *Repositories  {
    return &Repositories{
        List: NewListRepository(db),
        Task: NewTaskRepository(db),
    }
}