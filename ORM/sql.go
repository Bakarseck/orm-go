package orm

import "strings"

// The SQLBuilder type is used to construct SQL queries with parameters.
// @property {string} query - The `query` property is a string that represents the SQL query being
// built by the SQLBuilder. It will be used to store the SQL query as it is being constructed.
// @property {[]interface{}} parameters - The `parameters` property is a slice of `interface{}` type.
// It is used to store the values that will be used as parameters in the SQL query.
type SQLBuilder struct {
	query      string
	parameters []interface{}
}

// The NewSQLBuilder function returns a new instance of the SQLBuilder struct.
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

// The `Build()` function is a method of the `SQLBuilder` struct. It returns the constructed SQL query
// and the parameters as a tuple of type `(string, []interface{})`. This allows the caller to retrieve
// the final SQL query string and the parameters that were added to the query.
func (b *SQLBuilder) Build() (string, []interface{}) {
	return b.query, b.parameters
}

// The `Clear()` function is a method of the `SQLBuilder` struct. It is used to reset the `query` and
// `parameters` properties of the `SQLBuilder` instance to their initial values.
func (b *SQLBuilder) Clear() {
	b.parameters = nil
	b.query = ""
}

// The `Insert` method is a function of the `SQLBuilder` struct. It is used to construct an SQL
// `INSERT` statement.
func (builder *SQLBuilder) Insert(table *Table, values []interface{}) *SQLBuilder {

	builder.query += "INSERT INTO " + table.Name + " (" + strings.Join(table.GetFieldName(), ", ") + ")" + " VALUES ("
	for i := 0; i < len(table.AllFields); i++ {

		if i > 0 {
			builder.query += ", "
		}
		builder.query += "?"
		builder.parameters = append(builder.parameters, values[i])
	}
	builder.query += ")"
	return builder
}

// The `Update` method is a function of the `SQLBuilder` struct. It is used to construct an SQL
// `UPDATE` statement.
func (b *SQLBuilder) Update(updates *Modifier) *SQLBuilder {
	b.query += "UPDATE " + updates.Model.Name + " SET "
	var setClauses []string
	setClauses = append(setClauses, updates.field+" = ?")
	b.parameters = append(b.parameters, updates.value)
	b.query += strings.Join(setClauses, ", ")
	return b
}

// The `Delete()` method is a function of the `SQLBuilder` struct. It is used to construct an SQL
// `DELETE` statement.
func (b *SQLBuilder) Delete() *SQLBuilder {
	b.query += "DELETE "
	return b
}

// The `Select` method is a function of the `SQLBuilder` struct. It is used to construct an SQL
// `SELECT` statement.
func (b *SQLBuilder) Select(columns ...string) *SQLBuilder {
	b.query += "SELECT " + strings.Join(columns, ", ")
	return b
}


// The `From` method is a function of the `SQLBuilder` struct. It is used to construct an SQL `FROM`
// clause in a query.
func (b *SQLBuilder) From(table *Table) *SQLBuilder {
	b.query += " FROM " + table.Name
	return b
}

// The `Where` method is a function of the `SQLBuilder` struct. It is used to construct an SQL `WHERE`
// condition in a query.
func (b *SQLBuilder) Where(column string, value interface{}) *SQLBuilder {
	b.query += " WHERE " + column + " = ?"
	b.parameters = append(b.parameters, value)
	return b
}

// The `And` method is a function of the `SQLBuilder` struct. It is used to construct an SQL `AND`
// condition in a query.
func (b *SQLBuilder) And(column string, value interface{}) *SQLBuilder {
	b.query += " AND " + column + " = ?"
	b.parameters = append(b.parameters, value)
	return b
}

// The `Or` method is a function of the `SQLBuilder` struct. It is used to construct an SQL `OR`
// condition in a query.
func (b *SQLBuilder) Or(column string, value interface{}) *SQLBuilder {
	b.query += " OR " + column + " = ?"
	b.parameters = append(b.parameters, value)
	return b
}


