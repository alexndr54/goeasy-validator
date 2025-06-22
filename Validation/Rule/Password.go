package Rule

import (
	"fmt"
	"github.com/alexndr54/goeasy-validator/Validation/Helper/NameChange"
	"regexp"
)

// PasswordRule implementasi untuk aturan 'password' (contoh kompleks).
type PasswordRule struct{}

func (p *PasswordRule) GetName() string                   { return "password" }
func (p *PasswordRule) ParseParams(params []string) error { return nil }

func (p *PasswordRule) Validate(field string, value interface{}) (bool, string) {
	field = NameChange.FieldNameChange(field)
	strValue, ok := value.(string)
	if !ok {
		return false, fmt.Sprintf("%s harus berupa string untuk validasi kata sandi.", field)
	}

	if len(strValue) < 8 {
		return false, fmt.Sprintf("Kata sandi harus memiliki minimal 8 karakter.")
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(strValue)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(strValue)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(strValue)

	if !hasUpper || !hasLower || !hasDigit {
		return false, fmt.Sprintf("Kata sandi harus mengandung setidaknya satu huruf besar, satu huruf kecil, dan satu angka.")
	}
	return true, ""
}
