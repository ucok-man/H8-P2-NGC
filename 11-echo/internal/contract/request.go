package contract

import (
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/entity"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/pkg/validator"
)

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (e *ReqLogin) Validate(v *validator.Validator) error {
	v.Check(e.Username != "", "username", "must be provided")
	v.Check(e.Password != "", "password", "must be provided")
	v.Check(len(e.Password) >= 8, "password", "must be at least 8 bytes long")

	if !v.Valid() {
		return v.Error
	}
	return nil
}

type ReqRegister struct {
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
}

func (e *ReqRegister) Validate(v *validator.Validator) error {
	v.Check(e.Username != "", "username", "must be provided")
	v.Check(e.Password != "", "password", "must be provided")
	v.Check(len(e.Password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(e.DepositAmount >= 0, "deposit_amount", "must be provided and positive integer")

	if !v.Valid() {
		return v.Error
	}
	return nil
}

func (r ReqRegister) ToUser() (*entity.User, error) {
	user := &entity.User{
		Username:      r.Username,
		DepositAmount: r.DepositAmount,
	}
	if err := user.SetPassword(r.Password); err != nil {
		return nil, err
	}
	return user, nil
}
