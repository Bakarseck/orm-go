package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
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

	if _, err := os.Stat("migrates"); os.IsNotExist(err) {
        err := os.Mkdir("migrates", 0755)
        if err != nil {
            log.Fatal(err)
        }
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

		createTableSQL := CreateTable(v.Name(), AllField...)
		_, err = o.db.Exec(createTableSQL)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Créer un fichier SQL pour la table
		currentTime := time.Now()
		fileName := fmt.Sprintf("migrates/%s-create-table-%s.sql", currentTime.Format("15-04-2006-05-s"), v.Name())
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString(createTableSQL)
		if err != nil {
			log.Fatal(err)
		}
	}

}
