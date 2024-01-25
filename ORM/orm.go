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
		var foreignKeys []string
		v := reflect.TypeOf(table)

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := field.Type
			if fieldType.Kind() == reflect.Struct {
				for j := 0; j < fieldType.NumField(); j++ {
					structField := fieldType.Field(j)
					ormgoTag := structField.Tag.Get("orm-go")

					if strings.HasPrefix(ormgoTag, "FOREIGN_KEY") {
						foreignKeyDetails := strings.Split(ormgoTag, ":")
						if len(foreignKeyDetails) == 2 {
							foreignKeys = append(foreignKeys, fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s", structField.Name, foreignKeyDetails[1]))
						}
						ormgoTag = strings.TrimSpace(ormgoTag[:strings.Index(ormgoTag, "FOREIGN_KEY")])
					}

					AllField = append(AllField, NewField(structField.Name, structField.Type, ormgoTag))
				}
			} else {
				ormgoTag := field.Tag.Get("orm-go")

				if strings.HasPrefix(ormgoTag, "FOREIGN_KEY") {
					foreignKeyDetails := strings.Split(ormgoTag, ":")
					if len(foreignKeyDetails) == 2 {
						foreignKeys = append(foreignKeys, fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s", field.Name, foreignKeyDetails[1]))
					}
					ormgoTag = strings.TrimSpace(ormgoTag[:strings.Index(ormgoTag, "FOREIGN_KEY")])
				}

				AllField = append(AllField, NewField(field.Name, fieldType, ormgoTag))
			}
		}

		createTableSQL := CreateTable(v.Name(), AllField...)
		if len(foreignKeys) > 0 {
			createTableSQL = strings.TrimSuffix(createTableSQL, "\n)")
			createTableSQL += ",\n" + "\t"+ strings.Join(foreignKeys, ",\n") + "\n)"
		}

		_, err := o.db.Exec(createTableSQL)
        if err != nil {
            fmt.Println(err)
            return
        }

		// Créer un fichier SQL pour la table
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
