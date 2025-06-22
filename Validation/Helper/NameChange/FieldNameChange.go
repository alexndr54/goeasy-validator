package NameChange

import "strings"

func FieldNameChange(field string) string {
	field = strings.ReplaceAll(field, "_", " ")
	if len(field) == 0 {
		return field
	}
	return strings.ToUpper(field[:1]) + field[1:]
}
