package transform

import (
	"reflect"
	"strings"
)

func MapToExportedStruct(m map[string]any, s any) {
	// We expect "s" to be a pointer to a struct, hence we want to get the actual struct (Elem())
	v := reflect.ValueOf(s).Elem()

	for key, value := range m {
		// search the struct by field name using the capitalized key of the map
		// we use the capitalized key because the fields of the struct are exported hence capitalized
		field := v.FieldByName(capitalize(key))

		if field.IsValid() && field.CanSet() {
			val := reflect.ValueOf(value)

			// If the value of the key in the map is convertible to the type of the field in the struct
			// then we set the field of the struct to the converted value
			if val.Type().ConvertibleTo(field.Type()) {
				field.Set(val.Convert(field.Type()))
			}
		}
	}
}

func capitalize(word string) string {
	runes := []rune(word)

	return strings.ToUpper(string(runes[0])) + string(runes[1:])
}
