package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form it creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}


// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}


// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field,"This field cannot be balnck")
		}
	}
}


// Has check if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		return false
	}

	return true
}


// Valid returns true if there are not errors, otherwise return true
func (f *Form) Valid() bool {
  return len(f.Errors) == 0
}


// MinLenght checks for string min lenght
func (f *Form) MinLength(field string, lenght int, r *http.Request) bool {
  x := r.Form.Get(field)
  if len(x) < lenght {
	f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long",lenght))
	return false
  }
  return true
}