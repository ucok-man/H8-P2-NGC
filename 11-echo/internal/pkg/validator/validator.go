package validator

import (
	"fmt"
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Error error
}

func New() *Validator {
	return &Validator{Error: nil}
}

func AddError(key, message string) error {
	return fmt.Errorf(key, message)
}

func (v *Validator) Valid() bool {
	return v.Error == nil
}

func (v *Validator) AddError(key, message string) {
	v.Error = fmt.Errorf("%s %s", key, message)
}

func (v *Validator) Check(ok bool, key, message string) {
	if v.Error != nil {
		return
	}

	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) Reset() {
	v.Error = nil
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
