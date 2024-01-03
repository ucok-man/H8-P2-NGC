package model

import "golang.org/x/crypto/bcrypt"

type password struct {
	Plaintext *string `gorm:"-:all"`
	Hash      []byte  `gorm:"type:bytes;not null"`
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.Plaintext = &plaintextPassword
	p.Hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) error {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintextPassword))
	if err != nil {
		return err
	}

	return nil
}
