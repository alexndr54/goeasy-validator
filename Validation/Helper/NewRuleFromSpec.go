package Helper

import (
	"fmt"
	"github.com/alexndr54/goeasy-validator/Validation/Init"
	"github.com/alexndr54/goeasy-validator/Validation/Rule"
	"strings"
)

// NewRuleFromSpec membuat instance Rule berdasarkan spesifikasi string.
func NewRuleFromSpec(spec string) (Rule.RuleContract, error) {
	parts := strings.SplitN(spec, ":", 2)
	ruleName := parts[0]
	var params []string
	if len(parts) > 1 {
		params = strings.Split(parts[1], ",")
	}

	ruleFunc, ok := Init.RegisteredRules[ruleName]
	if !ok {
		return nil, fmt.Errorf("aturan '%s' tidak ditemukan", ruleName)
	}

	rule := ruleFunc()
	if err := rule.ParseParams(params); err != nil {
		return nil, fmt.Errorf("kesalahan parsing parameter untuk aturan '%s': %v", ruleName, err)
	}
	return rule, nil
}
