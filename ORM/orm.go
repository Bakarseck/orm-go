package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
)

func NewORM() *ORM {
	return &ORM{}
}

func (o *ORM) InitDB(name string) {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	o.db, err = sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTable(name string, fields ...*Field) string {
	sqlTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", name)
	for i, field := range fields {
		if i == len(fields)-1 {
			sqlTable += "\t" + TableField(field) + "\n)"
		} else {
			sqlTable += "\t" + TableField(field) + ",\n"
		}

	}
	return sqlTable
}

func (o *ORM) AutoMigrate(tables ...interface{}) {

	for _, table := range tables {
		var AllField []*Field
		v := reflect.TypeOf(table)

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := v.Field(i).Type
			if fieldType.Kind() == reflect.Struct {
				for i := 0; i < fieldType.NumField(); i++ {
					struct_field := fieldType.Field(i)
					ormgoTag := struct_field.Tag.Get(("orm-go"))
					AllField = append(AllField, NewField(struct_field.Name, struct_field.Type, ormgoTag))
				}
			} else {
				ormgoTag := field.Tag.Get("orm-go")
				AllField = append(AllField, NewField(field.Name, fieldType, ormgoTag))
			}
		}

		_, err := o.db.Exec(CreateTable(v.Name(), AllField...))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
