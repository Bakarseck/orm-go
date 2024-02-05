package orm

import (
	"database/sql"
	"reflect"
	"strings"
	"time"
)

// The ORM type represents an object-relational mapping with a database connection and a collection of
// tables.
// @property Db - The `Db` property is a pointer to an instance of the `sql.DB` struct. This is
// typically used to establish a connection to a database and perform various database operations such
// as querying and executing SQL statements.
// @property {[]*Table} Tables - The `Tables` property is a slice of `Table` structs. Each `Table`
// struct represents a database table and contains information about the table's name, columns, and
// relationships with other tables.
type ORM struct {
	Db     *sql.DB
	Tables []*Table
	Custom *SQLBuilder
}

// The above type represents a model with an auto-incrementing primary key and a default created at
// timestamp.
// @property {int64} Id - The `Id` property is an integer field that serves as the primary key for the
// model. It is annotated with `orm-go:"PRIMARY KEY AUTOINCREMENT"`, which indicates that it should be
// automatically incremented when new records are inserted into the database.
// @property CreatedAt - CreatedAt is a property of type time.Time that represents the timestamp of
// when the model instance was created. It is annotated with `orm-go:"DEFAULT CURRENT_TIMESTAMP"`,
// which means that the database will automatically set the value of CreatedAt to the current timestamp
// when a new record is inserted.
type Model struct {
	Id        int64     `orm-go:"PRIMARY KEY AUTOINCREMENT"`
	CreatedAt time.Time `orm-go:"DEFAULT CURRENT_TIMESTAMP"`
}

// The NewORM function returns a new instance of the ORM struct.
func NewORM() *ORM {
	return &ORM{
		Custom: NewSQLBuilder(),
	}
}

// The `AddTable` method of the `ORM` struct is used to add a new table to the ORM instance. It takes a
// pointer to a `Table` struct as a parameter and appends it to the `Tables` slice of the ORM instance.
// This allows the ORM to keep track of all the tables it is working with.
func (o *ORM) AddTable(t *Table) {
	o.Tables = append(o.Tables, t)
}

// The `GetTable` method of the `ORM` struct is used to retrieve a specific table from the ORM instance
// based on its name. It takes a string parameter `table` which represents the name of the table to be
// retrieved.
func (o *ORM) GetTable(table string) *Table {
	for _, t := range o.Tables {
		if t.Name == table {
			return t
		}
	}
	return nil
}

// The type "Field" represents a field in a struct, with properties for name, type, and tag.
// @property {string} Name - The Name property is a string that represents the name of the field.
// @property Type - The "Type" property in the "Field" struct is of type "reflect.Type". This means
// that it stores the type information of a field in a struct.
// @property {string} Tag - The "Tag" property in the "Field" struct is a string that represents any
// additional metadata or information associated with the field. It can be used to store and retrieve
// additional information about the field, such as validation rules, formatting instructions, or any
// other custom data that may be relevant to the field
type Field struct {
	Name string
	Type reflect.Type
	Tag  string
}

// The NewField function creates and returns a new Field object with the specified name, type, and tag.
func NewField(name string, tp reflect.Type, tag string) *Field {
	return &Field{
		Name: name,
		Type: tp,
		Tag:  tag,
	}
}

// The function `TableField` returns a string representation of a field in a table, including its name,
// type, and tag.
func TableField(f *Field) (fd string) {
	fd = strings.Join([]string{f.Name, GetType(f.Type), f.Tag}, " ")
	return
}

// The "Table" type represents a database table with a name, a list of fields, and a list of foreign
// keys.
// @property {string} Name - The Name property of the Table struct is a string that represents the name
// of the table.
// @property {[]*Field} AllFields - AllFields is a slice of pointers to Field objects. It represents
// all the fields/columns in the table.
// @property {[]string} ForeignKey - The `ForeignKey` property is a slice of strings that represents
// the foreign key constraints of the table. Each string in the slice represents a foreign key
// constraint, typically in the format of `foreign_table_name(foreign_column_name)`.
type Table struct {
	Name       string
	AllFields  []*Field
	ForeignKey []string
}

// The NewTable function creates a new instance of the Table struct with the given name.
func NewTable(name string) *Table {
	return &Table{
		Name: name,
	}
}

// The `AddField` method of the `Table` struct is used to add a new field to the table. It takes a
// pointer to a `Field` struct as a parameter and appends it to the `AllFields` slice of the table.
// This allows the table to keep track of all the fields/columns it has.
func (t *Table) AddField(f *Field) {
	t.AllFields = append(t.AllFields, f)
}

// The `GetFieldName` method of the `Table` struct is used to retrieve the names of all the fields in
// the table. It iterates over the `AllFields` slice of the table and appends the name of each field to
// a new slice called `names`. Finally, it returns the `names` slice, which contains all the field
// names.
func (t *Table) GetFieldName() []string {
	var names []string
	for _, field := range t.AllFields {
		names = append(names, field.Name)
	}
	return names
}

func (t *Table) GetField(name string) *Field {
	for _, v := range t.AllFields {
		if v.Name == name {
			return v
		}
	}
	return nil
}
