// Filename: internal/validator/validator.go

package validator

import (
	"regexp"
)

var (
	PositiveNumberRX = regexp.MustCompile(`^[1-9]+[0-9]*$`)
	EmailRX          = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	PhoneRX          = regexp.MustCompile(`^\+?\(?[0-9]{3}\)?\s?-\s?[0-9]{3}\s?-\s?[0-9]{4}$`)
)

// create a type validator errors map
type Validator struct {
	Errors map[string]string
}

// A new instance
func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// Valid checks the Errors map
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// In() checks if elements exist in the list of elements
func In(elements string, list ...string) bool {
	for i := range elements {

		if elements == list[i] {
			return true
		}

	}
	return false
}

// Match() returns true if match regex pattern
func Match(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// AddError apends an error entry to the Error map
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check() preforms the validation checks and calls the AddError method in turn if there is an error
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// Unique() checks for unique in the slice
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(uniqueValues) == len(values)
}
