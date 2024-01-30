package orm

import (
	"log"
	"reflect"
)

var (
	__BUILDER__ = NewSQLBuilder()
	__QUERY__   string
	__PARAMS__  []interface{}
)

// The `Insert` function is a method of the `ORM` struct. It takes in one or more tables as arguments,
// which are of type `interface{}`.
func (o *ORM) Insert(tables ...interface{}) {
	for _, t := range tables {

		_, __TABLE__ := InitTable(t)

		tType := reflect.TypeOf(t)
		tValue := reflect.ValueOf(t)

		if tType.Kind() == reflect.Struct {
			var values []interface{}
			
			for i := 0; i < tValue.NumField(); i++ {
				field := tValue.Field(i)
				fieldType := field.Type()

				switch fieldType.Kind() {
				case reflect.Int, reflect.Int64:
					values = append(values, field.Int())
				case reflect.String:
					values = append(values, field.String())
				case reflect.Float32, reflect.Float64:
					values = append(values, field.Float())
				}

				if fieldType.Kind() == reflect.Struct && fieldType.Name() == "Model" {
					__TABLE__.AllFields = __TABLE__.AllFields[2:]
				}
			}

			if len(values) > 0 {
				__QUERY__, __PARAMS__ = __BUILDER__.Insert(__TABLE__, values).Build()
				_, err := o.Db.Exec(__QUERY__, __PARAMS__...)
				if err != nil {
					log.Fatal(err)
				}
				__BUILDER__.Clear()
			}
		}
	}
}
