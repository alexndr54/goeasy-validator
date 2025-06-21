package Rule

import (
	"fmt"
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
	strValue, ok := value.(string)

	if !ok {
		return false, fmt.Sprintf("%s hanya boleh berupa string untuk validasi huruf.", field)
	}

	if match, _ := regexp.MatchString("^[a-zA-Z]*$", strValue); !match && strValue != "" {
		return false, fmt.Sprintf("%s hanya boleh huruf", field)
	}

	return true, ""
}
