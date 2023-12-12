package entity

import "database/sql"

type Entity struct {
	Hero    HeroEntity
	Villain VillainEntity
}

func New(db *sql.DB) *Entity {
	return &Entity{
		Hero:    HeroEntity{db: db},
		Villain: VillainEntity{db: db},
	}
}
