package Validation

import (
	"fmt"
	"github.com/alexndr54/goeasy-validator/Validation/Helper"
	"github.com/alexndr54/goeasy-validator/Validation/Model"
	"strings"
)

// NewValidator membuat instance Validator baru.
func NewValidator(data map[string]interface{}, rules map[string]string) *Validator {
	return &Validator{
		Data:   data,
		Rules:  rules,
		Errors: make(ValidationErrors),
	}
}

type ValidationErrors map[string][]string

func (ve ValidationErrors) Add(field, message string) {
	ve[field] = append(ve[field], message)
}

// HasErrors memeriksa apakah ada kesalahan validasi.
func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}

// Validate menjalankan proses validasi.
