package orm

import (
	"log"
	"reflect"
)

// This function `Scan` in the ORM package is responsible for executing a SQL query, scanning the
// results into struct fields, and returning a slice of the scanned results.
func (o *ORM) Scan(table interface{}, columns ...string) interface{} {
	Type, __table := InitTable(table)
	__BUILDER__ := NewSQLBuilder()
	var query string
	var param []interface{}
	__BUILDER__.custom = o.Custom
	query, param = __BUILDER__.Select(columns...).From(__table).Build()

	rows, err := o.Db.Query(query, param...)
	defer __BUILDER__.Clear()
	if err != nil {
		return nil
	}
	defer rows.Close()

	__results := reflect.MakeSlice(reflect.SliceOf(Type), 0, 0)

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

		newStruct := reflect.New(Type).Elem()

		for i, value := range values {
			val := reflect.ValueOf(value)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}

			newStruct.FieldByName(columns[i]).Set(val)
		}

		__results = reflect.Append(__results, newStruct)
	}

	return __results.Interface()
}
