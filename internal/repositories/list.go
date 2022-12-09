package repositories

import (
    "github.com/jmoiron/sqlx"
)

type List struct {
    db *sqlx.DB
}

func NewListRepository(db *sqlx.DB) *List  {
    return &List{db: db}
}