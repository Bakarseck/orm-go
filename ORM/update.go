package orm

import (
	"database/sql"
	"log"
	"reflect"
)

// Modifier represents a modification operation on a database table. It encapsulates
// information about a specific field to be updated, including the new value, any
// additional parameters, and a reference to the table schema.
type Modifier struct {
	field      string
	value      interface{}
	Parameters map[string]interface{}
	Model      *Table
}

// NewModifier creates a new instance of Modifier. It takes parameters for the modification,
// a reference to the table model, and the name of the field to be modified.
func NewModifier(params map[string]interface{}, m *Table, f string) *Modifier {
	return &Modifier{
		field:      f,
		Parameters: params,
		Model:      m,
	}
}

// Update applies the modification to the database. It constructs an SQL update query
// using the Modifier's details and executes it using the provided database connection.
func (m *Modifier) Update(db *sql.DB) {
	builder := NewSQLBuilder()
	query, parameters := builder.Update(m).Where(m.field, m.Parameters[m.field]).Build()
	_, err := db.Exec(query, parameters...)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateField sets the new value for the field that will be updated in the database.
// This method facilitates method chaining by returning a pointer to the Modifier.
func (m *Modifier) UpdateField(value interface{}) *Modifier {
	m.value = value
	return m
}

// SetModel is a method of the ORM struct that initializes a Modifier for a specific
// table model. It queries the database for the current values of the table row
// identified by 'nameField' and 'data', and then creates a Modifier with this
// current state, ready for updates.
func (o *ORM) SetModel(nameField string, data interface{}, tableName string) *Modifier {
	_table := o.GetTable(tableName)
	__params := make(map[string]interface{})

	builder := NewSQLBuilder()

	query, param := builder.Select("*").From(_table).Where(nameField, data).Build()
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
			__params[_table.AllFields[i].Name] = reflect.ValueOf(value).Elem().Interface()
		}
	}

	modif := NewModifier(__params, _table, nameField)
	return modif
}
