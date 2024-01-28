package orm

import (
	"fmt"
	"reflect"
	"strings"
)

// The GetTags function parses the "orm-go" tag of a struct field and returns the SQL attributes and
// foreign keys specified in the tag.
func GetTags(structField reflect.StructField) (string, []string) {
	ormgoTag := structField.Tag.Get("orm-go")
	if ormgoTag == "" {
		return "", nil
	}

	attributes := strings.Split(ormgoTag, " ")

	var sqlAttributes []string
	var foreignKeys []string

	for _, attr := range attributes {
		if strings.HasPrefix(attr, "FOREIGN_KEY") {
			foreignKeyDetails := strings.Split(attr, ":")
			if len(foreignKeyDetails) == 3 {
				foreignKey := fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s (%s)", structField.Name, foreignKeyDetails[1], foreignKeyDetails[2])
				foreignKeys = append(foreignKeys, foreignKey)
			}
		} else {
			sqlAttributes = append(sqlAttributes, attr)
		}
	}

	return strings.Join(sqlAttributes, " "), foreignKeys
}
