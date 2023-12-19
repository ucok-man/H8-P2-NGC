package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidTimeRequired = errors.New("invalid time required format (n mins)")

type TimeRequired int

func (tr TimeRequired) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", tr)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (tr *TimeRequired) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidTimeRequired
	}

	part := strings.Split(unquotedJSONValue, " ")

	if len(part) != 2 || part[1] != "mins" {
		return ErrInvalidTimeRequired
	}

	i, err := strconv.Atoi(part[0])
	if err != nil {
		return ErrInvalidTimeRequired
	}

	*tr = TimeRequired(i)

	return nil
}
