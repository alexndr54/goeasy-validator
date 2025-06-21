// requiredRule implementasi untuk aturan 'required'.
type requiredRule struct{}

func (r *requiredRule) GetName() string                   { return "required" }
func (r *requiredRule) ParseParams(params []string) error { return nil } // Tidak ada parameter

func (r *requiredRule) Validate(field string, value interface{}) (bool, string) {
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



package Filter

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationErrors map[string][]string

// Add menambahkan pesan kesalahan ke field tertentu.
func (ve ValidationErrors) Add(field, message string) {
	ve[field] = append(ve[field], message)
}

// HasErrors memeriksa apakah ada kesalahan validasi.
func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}

// Rule interface mendefinisikan metode yang harus diimplementasikan oleh setiap aturan validasi.
type Rule interface {
	Validate(field string, value interface{}) (bool, string)
	GetName() string
	ParseParams(params []string) error // Untuk aturan yang butuh parameter (e.g., min:10)
}

// --- 2. Implementasi Aturan Validasi Bawaan ---





// --- 3. Registrasi Aturan dan Factory ---

var registeredRules = make(map[string]func() Rule)

func init() {
	// Daftarkan semua aturan yang tersedia di sini
	registeredRules["required"] = func() Rule { return &requiredRule{} }
	registeredRules["email"] = func() Rule { return &emailRule{} }
	registeredRules["min"] = func() Rule { return &minRule{} }
	registeredRules["password"] = func() Rule { return &passwordRule{} }
}

