package Validation

type ValidationErrors map[string][]string

func (ve ValidationErrors) Add(field, message string) {
	ve[field] = append(ve[field], message)
}

// HasErrors memeriksa apakah ada kesalahan validasi.
func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}
