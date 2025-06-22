package Validation

// NewValidator membuat instance Validator baru.
func NewValidator(data map[string]interface{}, rules map[string]string) *Validator {
	return &Validator{
		Data:   data,
		Rules:  rules,
		Errors: make(ValidationErrors),
	}
}

// NewSimpleValidator membuat instance Validator sederhana yang mengembalikan kesalahan sebagai map,Example: map[Username: Wajib di isi,Email: Format email]
func NewSimpleValidator(data map[string]interface{}, rules map[string]string) map[string]string {
	validator4 := NewValidator(data, rules)
	errors4 := validator4.Validate()

	if errors4.HasErrors() {
		dataError := make(map[string]string)
		for field, msgs := range errors4 {
			dataError[field] = msgs[0]
		}

		return dataError
	} else {
		return nil
	}
}

// NewSingleRuleValidator memvalidasi semua aturan tetapi hanya mengembalikan 1 kesalahan (type string/field,type string/mesasge) (nama_lengkap: require,Wajib di isi)
func NewSingleRuleValidator(data map[string]interface{}, rules map[string]string) (string, string) {
	validator4 := NewValidator(data, rules)
	errors4 := validator4.Validate()

	if errors4.HasErrors() {
		for field, msgs := range errors4 {
			return field, msgs[0]
		}
	}

	return "", "" // Jika tidak ada kesalahan, kembalikan string kosong
}
