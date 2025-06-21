package Rule

type RuleContract interface {
	Validate(field string, value interface{}) (bool, string)
	GetName() string
	ParseParams(params []string) error // Untuk aturan yang butuh parameter (e.g., min:10)
}
