package orm

import "strings"

var (
	// The `Order` variable is a map that maps integers to strings. It is used to represent the order of
	// sorting in a query. The keys of the map are integers, where 0 represents ascending order ("ASC") and
	// 1 represents descending order ("DESC"). This map can be used to easily specify the sorting order in
	// a query by accessing the corresponding string value based on the desired order.
		Order = map[int]string{
			0: "ASC",
			1: "DESC",
		}
	)

// The SQLBuilder type is used to construct SQL queries with parameters.
// @property {string} query - The `query` property is a string that represents the SQL query being
// built by the SQLBuilder. It will be used to store the SQL query as it is being constructed.
// @property {[]interface{}} parameters - The `parameters` property is a slice of `interface{}` type.
// It is used to store the values that will be used as parameters in the SQL query.
type SQLBuilder struct {
	query      string
	parameters []interface{}
	custom     *SQLBuilder
}

// The NewSQLBuilder function returns a new instance of the SQLBuilder struct.
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

// The `Build()` function is a method of the `SQLBuilder` struct. It returns the constructed SQL query
// and the parameters as a tuple of type `(string, []interface{})`. This allows the caller to retrieve
// the final SQL query string and the parameters that were added to the query.
func (b *SQLBuilder) Build() (string, []interface{}) {
	if b.custom != nil {
		b.query += b.custom.query
		b.parameters = append(b.parameters, b.custom.parameters...)
	}
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
	setClauses = append(setClauses, updates.field2+" = ?")
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

// The `OrderBy` method is a function of the `SQLBuilder` struct. It is used to construct an SQL `ORDER
// BY` clause in a query.
func (b *SQLBuilder) OrderBy(column string, order int) *SQLBuilder {
	b.query += " ORDER BY " + column + " " + Order[order]
	return b
}

// The `Limit` method is a function of the `SQLBuilder` struct. It is used to construct an SQL `LIMIT`
// clause in a query.
func (b *SQLBuilder) Limit(limit int) *SQLBuilder {
	b.query += " LIMIT ?"
	b.parameters = append(b.parameters, limit)
	return b
}

// The `Join` method is a function of the `SQLBuilder` struct. It is used to construct an SQL `JOIN`
// clause in a query.
func (b *SQLBuilder) Join(table string, condition string) *SQLBuilder {
	b.query += " JOIN " + table + " ON " + condition
	return b
}

// The `GroupBy` method is a function of the `SQLBuilder` struct. It is used to construct an SQL `GROUP
// BY` clause in a query.
func (b *SQLBuilder) GroupBy(column string) *SQLBuilder {
	b.query += " GROUP BY " + column
	return b
}

// The `Having` method is a function of the `SQLBuilder` struct. It is used to construct an SQL
// `HAVING` clause in a query. The `HAVING` clause is used to filter the results of a query based on a
// condition that applies to the aggregated values.
func (b *SQLBuilder) Having(condition string) *SQLBuilder {
	b.query += " HAVING " + condition
	return b
}