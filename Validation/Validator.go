package Validation

import (
	"fmt"
	"github.com/alexndr54/goeasy-validator/Validation/Helper"
	"strings"
)

type Validator struct {
	Data   map[string]interface{}
	Rules  map[string]string // FieldName -> "rule1|rule2:param1,param2"
	Errors ValidationErrors
}

func (v *Validator) Validate() ValidationErrors {
	for field, ruleString := range v.Rules {
		// Dapatkan nilai dari data
		value, exists := v.Data[field]

		// Pisahkan multiple rules per field
		individualRules := strings.Split(ruleString, "|")

		for _, ruleSpec := range individualRules {
			rule, err := Helper.NewRuleFromSpec(ruleSpec)
			if err != nil {
				// Handle error parsing rule, mungkin tambahkan ke errors atau log
				v.Errors.Add(field, fmt.Sprintf("Kesalahan internal: %v", err))
				continue
			}

			// Khusus untuk 'required', kita perlu tahu apakah fieldnya ada atau tidak
			// Untuk rule lainnya, jika field tidak ada dan bukan 'required', kita lewati
			if !exists && rule.GetName() != "required" {
				continue // Lewati validasi jika field tidak ada dan bukan required
			}

			isValid, message := rule.Validate(field, value)
			if !isValid {
				v.Errors.Add(field, message)
				// Jika satu aturan gagal, kita bisa memilih untuk menghentikan validasi untuk field ini
				// Atau melanjutkan ke aturan berikutnya. Laravel defaultnya berhenti untuk field tersebut.
				break // Berhenti validasi untuk field ini jika ada satu aturan yang gagal
			}
		}
	}
	return v.Errors
}
