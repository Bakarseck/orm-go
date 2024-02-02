package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// The `InitDB` function is responsible for initializing the database connection and creating the
// necessary files and directories for database migration.
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

// The CreateTable function creates a SQL table with the given name and fields.
func CreateTable(name string, fields ...*Field) string {
	sqlTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", name)
	var all []string
	for _, field := range fields {
		all = append(all, "\t"+TableField(field))
	}
	sqlTable += strings.Join(all, ",\n") + "\n)"
	return sqlTable
}

// The `AutoMigrate` function is responsible for automatically creating database tables based on the
// provided struct definitions. It takes in a variadic parameter `tables` which represents the struct
// definitions of the tables to be created.
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

		upFileName := fmt.Sprintf("create-table-%s.up.sql", v.Name())
		downFileName := fmt.Sprintf("drop-table-%s.down.sql", v.Name())

		if _, err := os.Stat(upFileName); os.IsNotExist(err) {
			file, err := os.Create(upFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			_, err = file.WriteString(createTableSQL)
			if err != nil {
				log.Fatal(err)
			}

			downFile, err := os.Create(downFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer downFile.Close()

			dropTableSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s;", v.Name())
			_, err = downFile.WriteString(dropTableSQL)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// The function `InitTable` initializes a table by extracting field information from a given struct
// type and creating corresponding fields in the table.
func InitTable(table interface{}) (reflect.Type, *Table) {
	v := reflect.TypeOf(table)
	_table := NewTable(v.Name())
	var foreignKeys []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := field.Type

		if fieldType.Kind() == reflect.Struct {
			// Gérer les sous-structs récursivement.
			for j := 0; j < fieldType.NumField(); j++ {
				structField := fieldType.Field(j)
				ormgoTag, fk := GetTags(structField)
				foreignKeys = append(foreignKeys, fk...)
				_table.AddField(NewField(structField.Name, structField.Type, ormgoTag))
			}
		} else {
			// Gérer les champs normaux.
			ormgoTag, fk := GetTags(field)
			foreignKeys = append(foreignKeys, fk...)
			_table.AddField(NewField(field.Name, fieldType, ormgoTag))
		}
		_table.ForeignKey = foreignKeys
	}

	return v, _table
}
