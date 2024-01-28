package orm

import (
	"database/sql"
	"log"
	"reflect"
)

// The `Modifier` type represents a modification to a table field with associated parameters and a
// reference to the table model.
// @property {string} field - The "field" property is a string that represents the name of the field
// being modified.
// @property value - The `value` property is a variable that can hold any type of value. It can be used
// to store the value of a field in a data model or any other value that needs to be stored in the
// `Modifier` struct.
// @property Parameters - Parameters is a map that stores additional parameters for the Modifier. The
// keys in the map are strings, and the values can be of any type. These parameters can be used to
// provide additional information or configuration for the Modifier.
// @property Model - The `Model` property is a pointer to a `Table` struct.
type Modifier struct {
	field      string
	value      interface{}
	Parameters map[string]interface{}
	Model      *Table
}

// The function NewModifier creates a new Modifier object with the given parameters.
func NewModifier(params map[string]interface{}, m *Table, f string) *Modifier {
	return &Modifier{
		field:      f,
		Parameters: params,
		Model:      m,
	}
}

// The `Update` method is used to update a record in the database based on the parameters provided in
// the `Modifier` struct.
func (m *Modifier) Update(db *sql.DB) {
	builder := NewSQLBuilder()
	query, parameters := builder.Update(m).Where(m.field, m.Parameters[m.field]).Build()
	_, err := db.Exec(query, parameters...)
	if err != nil {
		log.Fatal(err)
	}
}

// The `UpdateField` method is used to set the value of a field that will be updated in the database.
// It takes a value as a parameter and assigns it to the `value` field of the `Modifier` struct. It
// then returns a pointer to the `Modifier` struct, allowing for method chaining.
func (m *Modifier) UpdateField(value interface{}) *Modifier {
	m.value = value
	return m
}

// The `SetModel` function is a method of the `ORM` struct. It is used to set the model for a specific
// table in the database.
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
