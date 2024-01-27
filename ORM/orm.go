package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
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

	if _, err := os.Stat("migrates"); os.IsNotExist(err) {
		err := os.Mkdir("migrates", 0755)
		if err != nil {
			log.Fatal(err)
		}
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
		all = append(all, "\t" + TableField(field))
	}
	sqlTable += strings.Join(all, ",\n") + "\n)"
	return sqlTable
}

func (o *ORM) AutoMigrate(tables ...interface{}) {
	for _, table := range tables {
		v, _table := InitTable(table)

		createTableSQL := CreateTable(v.Name(), _table.AllFields...)
		if len(_table.ForeignKey) > 0 {
			createTableSQL = strings.TrimSuffix(createTableSQL, "\n)")
			createTableSQL += ",\n" + "\t" + strings.Join(_table.ForeignKey, ",\n") + "\n)"
		}

		o.AddTable(_table)
		_, err := o.Db.Exec(createTableSQL)
		if err != nil {
			panic(err)
		}

		currentTime := time.Now()
		fileName := fmt.Sprintf("migrates/%s-create-table-%s.sql", currentTime.Format("2006-01-02-15-04-05"), v.Name())
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

func InitTable(table interface{}) (reflect.Type, *Table) {
	v := reflect.TypeOf(table)
	_table := NewTable(v.Name())
	var foreignKeys []string

	for i := 0; i < v.NumField(); i++ {

		field := v.Field(i)
		fieldType := v.Field(i).Type
		if fieldType.Kind() == reflect.Struct {

			for i := 0; i < fieldType.NumField(); i++ {
				struct_field := fieldType.Field(i)
				ormgoTag := struct_field.Tag.Get(("orm-go"))

				if strings.HasPrefix(ormgoTag, "FOREIGN_KEY") {
					foreignKeyDetails := strings.Split(ormgoTag, ":")
					if len(foreignKeyDetails) == 3 {
						foreignKeys = append(foreignKeys, fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s (%s)", struct_field.Name, foreignKeyDetails[1], foreignKeyDetails[2]))
					}
					ormgoTag = strings.TrimSpace(ormgoTag[:strings.Index(ormgoTag, "FOREIGN_KEY")])
				}
				_table.AddField(NewField(struct_field.Name, struct_field.Type, ormgoTag))
				_table.ForeignKey = foreignKeys
			}

		} else {
			ormgoTag := field.Tag.Get("orm-go")

			if strings.HasPrefix(ormgoTag, "FOREIGN_KEY") {
				foreignKeyDetails := strings.Split(ormgoTag, ":")
				if len(foreignKeyDetails) == 3 {
					foreignKeys = append(foreignKeys, fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s (%s)", field.Name, foreignKeyDetails[1], foreignKeyDetails[2]))
				}
				ormgoTag = strings.TrimSpace(ormgoTag[:strings.Index(ormgoTag, "FOREIGN_KEY")])
			}

			_table.AddField(NewField(field.Name, fieldType, ormgoTag))
			_table.ForeignKey = foreignKeys
		}
	}
	return v, _table
}
