package orm

import "reflect"

func GetType(fieldType reflect.Type) (sqlType string) {
	switch fieldType.Kind() {
	case reflect.Int64:
		sqlType = "INTEGER"
	case reflect.String:
		sqlType = "TEXT"
	case reflect.Float64:
		sqlType = "REAL"
	case reflect.Struct:
		if reflect.TypeOf(fieldType).Name() == "Time" {
			sqlType = "TIME"
		}
	}

	return sqlType
}

type Time struct {}

