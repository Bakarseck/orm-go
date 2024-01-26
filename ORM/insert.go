package orm

import (
	"fmt"
	"log"
	"reflect"
)

func (o *ORM) Insert(tables ...interface{}) {
	for _, t := range tables {
		if reflect.TypeOf(t).Kind() == reflect.Struct {
			var values []interface{}
			v := reflect.ValueOf(t)
			nameTable := reflect.TypeOf(t).Name()
			_table := o.GetTable(nameTable)

			for i := 0; i < v.NumField(); i++ {
				switch v.Field(i).Kind() {
				case reflect.Int:
				case reflect.Int64:
					values = append(values, v.Field(i).Int())
				case reflect.String:
					values = append(values, v.Field(i).String())
				case reflect.Float32:
				case reflect.Float64:
					values = append(values, v.Field(i).Float())
				case reflect.Struct:
					if v.Field(i).Type().Name() == "Model" {
						_table.AllFields = _table.AllFields[2:]
					}
				}
			}
			for _, v := range _table.AllFields {
				fmt.Println(v.Name)
			}
			builder := NewSQLBuilder()
			query, parameters := builder.Insert(_table, values).Build()
			_, err := o.Db.Exec(query, parameters...)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}