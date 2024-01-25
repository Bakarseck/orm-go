package orm

import "strings"

type SQLBuilder struct {
	query      string
	parameters []interface{}
}

func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

func (builder *SQLBuilder) Insert(table *Table, values []interface{}) *SQLBuilder{

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

func (b *SQLBuilder) Build() (string, []interface{}) {
	return b.query, b.parameters
}