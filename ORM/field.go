package orm

import (
	"reflect"
	"strings"
)

func NewField(name string, tp reflect.Type, tag string) *Field {
	return &Field{
		Name: name,
		Type: tp,
		Tag: tag,
	}
}

func TableField(f *Field) (fd string) {
	fd = strings.Join([]string{f.Name, GetType(f.Type), f.Tag}, " ")
	return
}