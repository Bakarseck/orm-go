package utils

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
)

func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

// The function `MapToStructs` takes a map of string slices and a target type, and returns a slice of
// structs with the values from the map assigned to the corresponding fields.
func MapToStructs(result map[string][]interface{}, targetType reflect.Type) interface{} {

	sliceLen := GetSliceLength(result)
	slice := reflect.MakeSlice(reflect.SliceOf(targetType), sliceLen, sliceLen)

	for i := 0; i < sliceLen; i++ {
		elem := slice.Index(i).Addr().Elem()
		for k, v := range result {
			field := elem.FieldByName(k)
			if field.IsValid() && field.CanSet() && i < len(v) {
				field.Set(reflect.ValueOf(v[i]))
			}
		}
	}

	return slice.Interface()
}

// The function "GetSliceLength" returns the length of the first slice found in the given map.
func GetSliceLength(result map[string][]interface{}) int {
	for _, v := range result {
		return len(v)
	}
	return 0
}
