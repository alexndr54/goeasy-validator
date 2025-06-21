package Filter

import (
	"fmt"
	"strconv"
)

// minRule implementasi untuk aturan 'min'.
type minRule struct {
	minLength int
}

func (m *minRule) GetName() string { return "min" }
func (m *minRule) ParseParams(params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("aturan min membutuhkan parameter panjang minimum")
	}
	length, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("parameter min harus berupa angka: %v", err)
	}
	m.minLength = length
	return nil
}

func (m *minRule) Validate(field string, value interface{}) (bool, string) {
	strValue, ok := value.(string)
	if !ok {
		return false, fmt.Sprintf("%s harus berupa string untuk validasi panjang.", field)
	}
	if len(strValue) < m.minLength {
		return false, fmt.Sprintf("%s harus memiliki minimal %d karakter.", field, m.minLength)
	}
	return true, ""
}
