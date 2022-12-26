package repositories

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    entity "todo-app/internal/entities"
)

const ListTable = "list"

type List struct {
    db *sqlx.DB
}

func NewListRepository(db *sqlx.DB) *List  {
    return &List{db: db}
}

func (l *List) GetAll() ([]entity.List, error){
    var lists = make([]entity.List, 0)

    query := fmt.Sprintf("SELECT id, name, description FROM %s", ListTable)
    err := l.db.Select(&lists, query)

    return lists, err
}

func (l *List) GetById(id int) (entity.List, error){
    var list entity.List

    query := fmt.Sprintf("SELECT id, name, description FROM %s WHERE id = $1", ListTable)
    err := l.db.Get(&list, query, id)

    return list, err
}

func (l *List) Update(list *entity.List) (*entity.List, error) {
    query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2 WHERE id = $3 RETURNING *", ListTable)
    err := l.db.Get(list, query, list.Name, list.Description, list.ID)

    return list, err
}

func (l *List) Create(list *entity.List) (*entity.List, error) {
    query := fmt.Sprintf("INSERT INTO %s (name, description) VALUES ($1, $2) RETURNING *", ListTable)
    err := l.db.Get(list, query, list.Name, list.Description)

    return list, err
}

func (l *List) Delete(id int) error {
    query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", ListTable)

    _, err := l.db.Exec(query, id)

    return err
}