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
			fmt.Println("Query: ", query)
			fmt.Println("Parameters: ", parameters)
			_, err := o.Db.Exec(query, parameters...)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// func (o *ORM) Insert(tables ...interface{}) {

// 	for _, t := range tables {

// 		if reflect.TypeOf(t).Kind() == reflect.Struct {
// 			v := reflect.ValueOf(t)
// 			nameTable := reflect.TypeOf(t).Name()
// 			deb := 0
// 			query := fmt.Sprintf("INSERT INTO %s (", nameTable)

// 			v1 := reflect.TypeOf(t)
// 			for i := 0; i < v1.NumField(); i++ {
// 				if v1.Field(i).Name != "Model" {
// 					query = fmt.Sprintf("%s %s,", query, v1.Field(i).Name)
// 				}

// 			}
// 			query = fmt.Sprintf("%s) VALUES (", query[:len(query)-1])

// 			for i := 0; i < v.NumField(); i++ {
// 				if v.Field(i).Type().Name() == "Model" {
// 					deb = 1
// 				}
// 				switch v.Field(i).Kind() {
// 				case reflect.Int:
// 				case reflect.Int64:
// 					if i == deb {
// 						query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
// 					} else {
// 						query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
// 					}
// 				case reflect.String:
// 					if i == deb {
// 						query = fmt.Sprintf("%s \"%s\"", query, v.Field(i).String())
// 					} else {
// 						query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
// 					}
// 				case reflect.Float32:
// 				case reflect.Float64:
// 					if i == deb {
// 						query = fmt.Sprintf("%s %f", query, v.Field(i).Float())
// 					} else {
// 						query = fmt.Sprintf("%s, %f", query, v.Field(i).Float())
// 					}
// 				}
// 			}
// 			query = fmt.Sprintf("%s)", query)
// 			//fmt.Println(query)
// 			_, err := o.db.Exec(query)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 	}
// }
