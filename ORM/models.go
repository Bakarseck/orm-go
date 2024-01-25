package orm

import (
	"database/sql"
	"reflect"
	"strings"
	"time"
)

type ORM struct {
	db     *sql.DB
	Tables []*Table
}

type Model struct {
	Id        int64       `orm-go:"PRIMARY KEY AUTOINCREMENT"`
	CreatedAt time.Time `orm-go:"DEFAULT CURRENT_TIMESTAMP"`
}

func NewORM() *ORM {
	return &ORM{}
}

func (o *ORM) AddTable(t *Table) {
	o.Tables = append(o.Tables, t)
}

func (o *ORM) GetTable(table string) *Table {
	for _, t := range o.Tables {
		if t.Name == table {
			return t
		}
	}
	return nil
}

type Field struct {
	Name string
	Type reflect.Type
	Tag  string
}

func NewField(name string, tp reflect.Type, tag string) *Field {
	return &Field{
		Name: name,
		Type: tp,
		Tag:  tag,
	}
}

func TableField(f *Field) (fd string) {
	fd = strings.Join([]string{f.Name, GetType(f.Type), f.Tag}, " ")
	return
}

type Table struct {
	Name      string
	AllFields []*Field
}

func NewTable(name string) *Table {
	return &Table{
		Name: name,
	}
}

func (t *Table) AddField(f *Field) {
	t.AllFields = append(t.AllFields, f)
}

func (t *Table) GetFieldName() []string {
	var names []string
	for _, field := range t.AllFields {
		names = append(names, field.Name)
	}
	return names
}
