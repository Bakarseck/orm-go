package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func (o *ORM) InitDB(name string) {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	o.Db, err = sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTable(name string, fields ...*Field) string {
	sqlTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", name)
	var all []string
	for _, field := range fields {
		all = append(all, TableField(field))
	}
	sqlTable += strings.Join(all, ",\n") + "\n)"
	return sqlTable
}

func (o *ORM) AutoMigrate(tables ...interface{}) {
	for _, table := range tables {
		v := reflect.TypeOf(table)
		_table := NewTable(v.Name())

		for i := 0; i < v.NumField(); i++ {

			field := v.Field(i)
			fieldType := v.Field(i).Type
			if fieldType.Kind() == reflect.Struct {

				for i := 0; i < fieldType.NumField(); i++ {
					struct_field := fieldType.Field(i)
					ormgoTag := struct_field.Tag.Get(("orm-go"))
					_table.AddField(NewField(struct_field.Name, struct_field.Type, ormgoTag))
				}

			} else {
				ormgoTag := field.Tag.Get("orm-go")

				_table.AddField(NewField(field.Name, fieldType, ormgoTag))
			}
		}

		o.AddTable(_table)
		fmt.Println(CreateTable(v.Name(), _table.AllFields...))
		_, err := o.Db.Exec(CreateTable(v.Name(), _table.AllFields...))
		if err != nil {
			panic(err)
		}
		
	}
}
