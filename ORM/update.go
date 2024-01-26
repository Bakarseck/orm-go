package orm

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type Modifier struct {
	field      string
	value      interface{}
	Parameters map[string]interface{}
	Model      *Table
}

func NewModifier(params map[string]interface{}, m *Table) *Modifier {
	return &Modifier{
		Parameters: params,
		Model:      m,
	}
}

func (m *Modifier) Update(db *sql.DB) {
	builder := NewSQLBuilder()
	query, parameters := builder.Update(m).Where(m.field, m.Parameters[m.field]).Build()
	_, err := db.Exec(query, parameters...)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Modifier) UpdateField(nameField string, value interface{}) *Modifier {
	m.field = nameField
	m.value = value
	return m
}

func (o *ORM) SetModel(column string, data interface{}, tableName string) *Modifier {
	_table := o.GetTable(tableName)
	__params := make(map[string]interface{})

	builder := NewSQLBuilder()

	query, param := builder.Select().From(_table).Where(column, data).Build()
	result, err := o.Db.Query(query, param...)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next() {
		values := make([]interface{}, 0)
		for _, v := range _table.AllFields {
			values = append(values, reflect.New(v.Type).Interface())
		}
		err := result.Scan(values...)
		if err != nil {
			log.Fatal(err)
		}

		for i, value := range values {
			fmt.Println(reflect.ValueOf(value).Elem().Interface())
			__params[_table.AllFields[i].Name] = reflect.ValueOf(value).Elem().Interface()
		}
	}

	modif := NewModifier(__params, _table)
	fmt.Println(modif)

	return modif
}
