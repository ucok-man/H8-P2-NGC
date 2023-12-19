package entity

import (
	"database/sql"
	"fmt"
)

var (
	ErrorNotFound     = fmt.Errorf("record not found")
	ErrDuplicateEntry = fmt.Errorf("dupilcate entry")
)

type Entity struct {
	User   UserServive
	Recipe RecipeService
}

func New(db *sql.DB) Entity {
	return Entity{
		User:   UserServive{db: db},
		Recipe: RecipeService{db: db},
	}
}
