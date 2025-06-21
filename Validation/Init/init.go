package Init

import "github.com/alexndr54/goeasy-validator/Validation/Rule"

var RegisteredRules = make(map[string]func() Rule.RuleContract)

func init() {
	// Daftarkan semua aturan yang tersedia di sini
	RegisteredRules["required"] = func() Rule.RuleContract { return &Rule.RequiredRule{} }
	RegisteredRules["email"] = func() Rule.RuleContract { return &Rule.EmailRule{} }
	RegisteredRules["min"] = func() Rule.RuleContract { return &Rule.MinRule{} }
	RegisteredRules["max"] = func() Rule.RuleContract { return &Rule.MaxRule{} }
	RegisteredRules["password"] = func() Rule.RuleContract { return &Rule.PasswordRule{} }
	RegisteredRules["only_letters"] = func() Rule.RuleContract { return &Rule.OnlyLetters{} }
}
