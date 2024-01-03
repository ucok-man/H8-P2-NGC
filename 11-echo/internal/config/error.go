package config

import "strings"

type ErrorMissingValue struct {
	Listenv []string
	Errors  []error
}

func (e ErrorMissingValue) Error() string {
	buff := strings.Builder{}
	for _, err := range e.Errors {
		buff.WriteString(err.Error())
	}
	return buff.String()
}
