package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type TimeRequired int

func (tr TimeRequired) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", tr)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (tr *TimeRequired) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	part := strings.Split(unquotedJSONValue, " ")

	if len(part) != 2 || part[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.Atoi(part[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*tr = TimeRequired(i)

	return nil
}
