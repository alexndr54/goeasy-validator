package Rule

import (
	"fmt"
	"github.com/alexndr54/goeasy-validator/Validation/Helper/NameChange"
	"reflect"
	"strings"
)

type RequiredRule struct{}

func (r *RequiredRule) GetName() string                   { return "required" }
func (r *RequiredRule) ParseParams(params []string) error { return nil } // Tidak ada parameter

func (r *RequiredRule) Validate(field string, value interface{}) (bool, string) {
	field = NameChange.FieldNameChange(field)
	if value == nil {
		return false, fmt.Sprintf("%s tidak boleh kosong.", field)
	}
	switch v := value.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return false, fmt.Sprintf("%s tidak boleh kosong.", field)
		}
	case int:
		if v == 0 { // Angka 0 mungkin dianggap kosong tergantung konteks, sesuaikan.
			return false, fmt.Sprintf("%s tidak boleh kosong.", field)
		}
	case float64:
		if v == 0.0 { // Angka 0.0 mungkin dianggap kosong tergantung konteks, sesuaikan.
			return false, fmt.Sprintf("%s tidak boleh kosong.", field)
		}
	case bool:
		// Boolean selalu dianggap ada jika bukan nil
	default:
		// Untuk tipe lain, anggap ada jika bukan nil
		val := reflect.ValueOf(value)
		if !val.IsValid() || (val.Kind() == reflect.Ptr && val.IsNil()) {
			return false, fmt.Sprintf("%s tidak boleh kosong.", field)
		}
	}
	return true, ""
}
