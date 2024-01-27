package orm

import (
	"reflect"
)

// The function GetType takes a reflect.Type as input and returns a string representing the
// corresponding SQL data type.
func GetType(fieldType reflect.Type) (sqlType string) {
	switch fieldType.Kind() {
	case reflect.Int:	
	case reflect.Int64:
		sqlType = "INTEGER"
	case reflect.String:
		sqlType = "TEXT"
	case reflect.Float64:
		sqlType = "REAL"
	case reflect.Struct:

		if fieldType.Name() == "Time" {
			sqlType = "DATETIME"
		}
	}
	return sqlType
}
