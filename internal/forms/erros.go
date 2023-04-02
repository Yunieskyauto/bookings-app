package forms



type errors map[string][]string

// Adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(fiels string) string {
	es := e[fiels]

	if len(es) == 0 {
		return ""
	}

	return es[0]
}