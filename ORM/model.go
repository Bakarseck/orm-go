package orm

import (
	"database/sql"
	"reflect"
	"time"
)

type ORM struct {
	db *sql.DB
}

type Model struct {
	ID        int     `orm-go:"PRIMARY KEY AUTOINCREMENT"`
	CreatedAt time.Time `orm-go:"DEFAULT CURRENT_TIMESTAMP"`
}

type Field struct {
	Name string
	Type reflect.Type
	Tag  string
}
