package orm

import (
	"log"
	"reflect"
)

func (o *ORM) Scan(table interface{}, columns ...string) map[string][]interface{} {
	_, __table := InitTable(table)
	__BUILDER__ := NewSQLBuilder()
	query, param := __BUILDER__.Select(columns...).From(__table).Build()
	rows, err := o.Db.Query(query, param...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	__results := make(map[string][]interface{})
	for rows.Next() {
		values := make([]interface{}, 0)
		for _, name := range columns {
			for _, v := range __table.AllFields {
				if name == v.Name {
					values = append(values, reflect.New(v.Type).Interface())
				}
			}
		}
		err := rows.Scan(values...)
		if err != nil {
			log.Fatal(err)
		}

		for i, value := range values {
			__results[columns[i]] = append(__results[columns[i]], reflect.ValueOf(value).Elem().Interface())
		}
	}
	return __results
}
