package orm

import (
	"log"
	"reflect"

	"github.com/Bakarseck/orm/utils"
)

// The `Scan` function is a method of the `ORM` struct. It takes in a `table` interface and a variadic
// parameter `columns` of type string. It returns a map with string keys and slice of interface{}
// values.
func (o *ORM) Scan(table interface{}, columns ...string) interface{} {
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

	structSlice := utils.MapToStructs(__results, reflect.TypeOf(table))

	return structSlice
}
