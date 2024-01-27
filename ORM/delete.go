package orm

func (o *ORM) Delete(table interface{}, column string, value interface{}) {
	_, __table := InitTable(table)
	__BUILDER__ := NewSQLBuilder()
	query, param := __BUILDER__.Delete().From(__table).Where(column, value).Build()
	_, err := o.Db.Exec(query, param...)
	if err != nil {
		panic(err)
	}
}