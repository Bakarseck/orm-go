package orm

// The `Delete` function is a method of the `ORM` struct. It takes three parameters: `table`, `column`,
// and `value`.
func (o *ORM) Delete(table interface{}, column string, value interface{}) {
	_, __table := InitTable(table)
	__BUILDER__ := NewSQLBuilder()
	query, param := __BUILDER__.Delete().From(__table).Where(column, value).Build()
	_, err := o.Db.Exec(query, param...)
	if err != nil {
		panic(err)
	}
}