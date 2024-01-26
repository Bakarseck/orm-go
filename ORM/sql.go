package orm

import "strings"

type SQLBuilder struct {
	query      string
	parameters []interface{}
}

func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

func (b *SQLBuilder) Build() (string, []interface{}) {
	return b.query, b.parameters
}

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

func (b *SQLBuilder) Update(updates *Modifier) *SQLBuilder {
	b.query += "UPDATE " + updates.Model.Name + " SET "
	var setClauses []string
	setClauses = append(setClauses, updates.field+" = ?")
	b.parameters = append(b.parameters, updates.value)
	b.query += strings.Join(setClauses, ", ")
	return b
}

func (b *SQLBuilder) From(table *Table) *SQLBuilder {
	b.query += " FROM " + table.Name
	return b
}

func (b *SQLBuilder) Where(column string, value interface{}) *SQLBuilder {
	b.query += " WHERE " + column + " = ?"
	b.parameters = append(b.parameters, value)
	return b
}

func (b *SQLBuilder) And(column string, value interface{}) *SQLBuilder {
	b.query += " AND " + column + " = ?"
	b.parameters = append(b.parameters, value)
	return b
}

func (b *SQLBuilder) Or(column string, value interface{}) *SQLBuilder {
	b.query += " OR " + column + " = ?"
	b.parameters = append(b.parameters, value)
	return b
}

func (b *SQLBuilder) Select() *SQLBuilder {
	b.query += "SELECT *"
	return b
}
