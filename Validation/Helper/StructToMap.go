package Helper

import (
	"reflect"
	"strings"
)

// StructToMap --- Fungsi Helper untuk Mengambil Data dari Struct (Opsional) ---
// Ini untuk mensimulasikan $request->all() jika inputnya adalah struct Go.
func StructToMap(obj interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		panic("Input harus berupa struct atau pointer ke struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i).Interface()
		data[strings.ToLower(field.Name)] = fieldValue // Mengubah nama field ke lowercase
	}
	return data
}
