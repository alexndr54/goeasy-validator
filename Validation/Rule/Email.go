package Filter

import (
	"fmt"
	"regexp"
)

// emailRule implementasi untuk aturan 'email'.
type emailRule struct{}

func (e *emailRule) GetName() string                   { return "email" }
func (e *emailRule) ParseParams(params []string) error { return nil } // Tidak ada parameter

func (e *emailRule) Validate(field string, value interface{}) (bool, string) {
	strValue, ok := value.(string)
	if !ok {
		return false, fmt.Sprintf("%s harus berupa string untuk validasi email.", field)
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(strValue) {
		return false, fmt.Sprintf("Format %s tidak valid.", field)
	}
	return true, ""
}
