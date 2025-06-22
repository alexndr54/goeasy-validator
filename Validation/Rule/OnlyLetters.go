package Rule

import (
	"fmt"
	"github.com/alexndr54/goeasy-validator/Validation/Helper/NameChange"
	"regexp"
)

// MinRule implementasi untuk aturan 'min'.
type OnlyLetters struct {
	maxLength int
}

func (m *OnlyLetters) GetName() string { return "min" }
func (m *OnlyLetters) ParseParams(params []string) error {
	return nil
}

func (m *OnlyLetters) Validate(field string, value interface{}) (bool, string) {
	field = NameChange.FieldNameChange(field)
	strValue, ok := value.(string)

	if !ok {
		return false, fmt.Sprintf("%s hanya boleh berupa string untuk validasi huruf.", field)
	}

	if match, _ := regexp.MatchString("^[a-zA-Z]*$", strValue); !match && strValue != "" {
		return false, fmt.Sprintf("%s hanya boleh huruf", field)
	}

	return true, ""
}
