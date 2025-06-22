package Rule

import (
	"fmt"
	"github.com/alexndr54/goeasy-validator/Validation/Helper/NameChange"
	"regexp"
)

// EmailRule implementasi untuk aturan 'email'.
type EmailRule struct{}

func (e *EmailRule) GetName() string                   { return "email" }
func (e *EmailRule) ParseParams(params []string) error { return nil } // Tidak ada parameter

func (e *EmailRule) Validate(field string, value interface{}) (bool, string) {
	field = NameChange.FieldNameChange(field)
	
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
