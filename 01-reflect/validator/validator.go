package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Value any
}

func New(value any) *Validator {
	return &Validator{
		Value: value,
	}
}

func (v *Validator) Validate() error {
	t := reflect.TypeOf(v.Value)

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("err: to perform validation, value must be struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		typeofField := field.Type.Kind()

		switch typeofField {
		case reflect.String:
			val := reflect.ValueOf(v.Value).Field(i).String()
			if err := validateString(val, field); err != nil {
				return err
			}

		case reflect.Int:
			val := reflect.ValueOf(v.Value).Field(i).Int()
			if err := validateInt(val, field); err != nil {
				return err
			}

		default:
			return fmt.Errorf("err: unsuported type of field. must be int or string")
		}
	}

	return nil
}

func validateString(val string, field reflect.StructField) error {
	if tagval := field.Tag.Get("required"); tagval != "" {
		switch {
		case tagval == "true":
			if len(val) < 1 {
				return fmt.Errorf("err: %v field on tag required. value must be provided", field.Name)
			}

		case tagval != "false":
			return fmt.Errorf("err: %v field on tag required. tag value must be false or true", field.Name)
		}
	}

	if tagval := field.Tag.Get("min"); tagval != "" {
		validTagval, err := strconv.Atoi(tagval)
		if err != nil {
			return fmt.Errorf("err: %v field on tag min. must be valid integer tag value", field.Name)
		}

		if len(val) < validTagval {
			return fmt.Errorf("err: %v field on tag min. length of value must be at least %d bytes", field.Name, validTagval)
		}
	}

	if tagval := field.Tag.Get("max"); tagval != "" {
		validTagval, err := strconv.Atoi(tagval)
		if err != nil {
			return fmt.Errorf("err: %v field on tag max. must be valid integer tag value", field.Name)
		}

		if len(val) > validTagval {
			return fmt.Errorf("err: %v field on tag max. length of value must not be greater than %d bytes", field.Name, validTagval)
		}
	}

	if tagval := field.Tag.Get("isEmail"); tagval != "" {
		switch {
		case tagval == "true":
			if !EmailRX.MatchString(val) {
				return fmt.Errorf("err: %v field on tag isEmail. value must be valid email form", field.Name)
			}

		case tagval != "false":
			return fmt.Errorf("err: %v field on tag isEmail. tag value must be false or true", field.Name)
		}
	}

	return nil
}

func validateInt(val int64, field reflect.StructField) error {
	if tagval := field.Tag.Get("required"); tagval != "" {
		switch {
		case tagval == "true":
			if val == 0 {
				return fmt.Errorf("err: %v field on tag required. value must be provided", field.Name)
			}

		case tagval != "false":
			return fmt.Errorf("err: %v field on tag required. tag value must be false or true", field.Name)
		}
	}

	if tagval := field.Tag.Get("min"); tagval != "" {
		validTagval, err := strconv.ParseInt(tagval, 10, 64)
		if err != nil {
			return fmt.Errorf("err: %v field on tag min. must be valid integer tag value", field.Name)
		}

		if val < validTagval {
			return fmt.Errorf("err: %v field on tag min. value must be greater than %d", field.Name, validTagval)
		}
	}

	if tagval := field.Tag.Get("max"); tagval != "" {
		validTagval, err := strconv.ParseInt(tagval, 10, 64)
		if err != nil {
			return fmt.Errorf("err: %v field on tag max. must be valid integer tag value", field.Name)
		}

		if val > validTagval {
			return fmt.Errorf("err: %v field on tag max. value must be less than %d", field.Name, validTagval)
		}
	}

	return nil
}
