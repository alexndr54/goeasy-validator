package Rule

import (
	"fmt"
	"strconv"
)

// MinRule implementasi untuk aturan 'min'.
type MaxRule struct {
	maxLength int
}

func (m *MaxRule) GetName() string { return "min" }
func (m *MaxRule) ParseParams(params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Contoh penggunaan: max:50, aturan min membutuhkan parameter panjang karakter")
	}
	length, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("parameter max harus berupa angka: %v", err)
	}

	m.maxLength = length
	return nil
}

func (m *MaxRule) Validate(field string, value interface{}) (bool, string) {
	strValue, ok := value.(string)
	if !ok {
		return false, fmt.Sprintf("%s harus berupa string untuk validasi panjang.", field)
	}
	if len(strValue) > m.maxLength {
		return false, fmt.Sprintf("Karakter %s terlalu panjang,Maksimal %d.", field, m.maxLength)
	}
	return true, ""
}
