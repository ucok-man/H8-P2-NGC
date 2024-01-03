package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID        int     `json:"user_id"`
	Username      string  `json:"username"`
	Password      string  `json:"-"`
	DepositAmount float64 `json:"deposit_amount"`
}

func (u *User) SetPassword(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) MatchesPassword(plaintextPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plaintextPassword))
	if err != nil {
		return err
	}

	return nil
}
