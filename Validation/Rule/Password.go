package Filter

import (
	"fmt"
	"regexp"
)

// passwordRule implementasi untuk aturan 'password' (contoh kompleks).
type passwordRule struct{}

func (p *passwordRule) GetName() string                   { return "password" }
func (p *passwordRule) ParseParams(params []string) error { return nil }

func (p *passwordRule) Validate(field string, value interface{}) (bool, string) {
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
