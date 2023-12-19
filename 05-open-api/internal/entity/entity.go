package entity

import (
	"database/sql"
	"fmt"
)

var (
	ErrorNotFound = fmt.Errorf("record not found")
)

type Entity struct {
	Hero      HeroEntity
	Villain   VillainEntity
	CrimeCase CrimeCaseEntity
}

func New(db *sql.DB) *Entity {
	return &Entity{
		Hero:      HeroEntity{db: db},
		Villain:   VillainEntity{db: db},
		CrimeCase: CrimeCaseEntity{db: db},
	}
}
