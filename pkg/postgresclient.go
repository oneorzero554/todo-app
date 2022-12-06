package dbclients

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    "log"
)

type Config struct {
    Host     string
    Port     string
    Username string
    Password string
    DBName   string
    SSLMode  string
}

func NewPostgresClient(cfg Config) (*sqlx.DB, error) {
    db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

    if err != nil {
        log.Fatalln(err)
    }

    return db, nil
}