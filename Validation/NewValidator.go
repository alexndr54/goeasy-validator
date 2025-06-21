package Validation

// NewValidator membuat instance Validator baru.
func NewValidator(data map[string]interface{}, rules map[string]string) *Validator {
	return &Validator{
		Data:   data,
		Rules:  rules,
		Errors: make(ValidationErrors),
	}
}

func NewSimpleValidator(data map[string]interface{}, rules map[string]string) map[string][]string {
	validator4 := NewValidator(data, rules)
	errors4 := validator4.Validate()

	if errors4.HasErrors() {
		dataError := make(map[string][]string)
		for field, msgs := range errors4 {
			for _, msgg := range msgs {
				dataError[field] = append(dataError[field], msgg)
			}
		}

		return dataError
	} else {
		return nil
	}
}

