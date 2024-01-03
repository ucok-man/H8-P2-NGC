package api

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/contract"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/pkg/validator"
)

type AppValidator struct {
	validator *validator.Validator
}

func NewAppValidator() *AppValidator {
	return &AppValidator{
		validator: validator.New(),
	}
}

func (v *AppValidator) Validate(i interface{}) error {
	if !structs.IsStruct(i) {
		panic("[api.Validate] ERROR : input must be struct type")
	}

	if v.validator.Error != nil {
		v.validator.Reset()
	}

	switch e := i.(type) {
	case *contract.ReqLogin:
		return e.Validate(v.validator)
	case *contract.ReqRegister:
		return e.Validate(v.validator)
	default:
		panic(fmt.Sprintf("[api.Validate] ERROR : type is not registered %v", structs.Name(e)))
	}
}
