package entity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ucok-man/H8-P2-NGC/05-open-api/internal/validator"
)

type Hero struct {
	HeroID   int64  `json:"hero_id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	ImageURL string `json:"image_url"`
}

func ValidateHero(v *validator.Validator, entity *Hero) {
	v.Check(entity.Name != "", "name", "must be provided")
	v.Check(len(entity.Name) <= 255, "name", "must not be more than 255 bytes long")

	v.Check(entity.Universe != "", "universe", "must be provided")
	v.Check(len(entity.Universe) <= 255, "universe", "must not be more than 255 bytes long")

	v.Check(entity.ImageURL != "", "image_url", "must be provided")
	// v.Check(validator.Matches(entity.ImageURL, validator.URLRX), "image_url", "must be valid url")
}

type HeroEntity struct {
	db *sql.DB
}

func (e *HeroEntity) GetAll() ([]*Hero, error) {
	query := `
		SELECT hero_id, name, universe, image_url
		FROM heroes
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := e.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var heroes []*Hero
	for rows.Next() {
		var hero = Hero{}

		err := rows.Scan(&hero.HeroID, &hero.Name, &hero.Universe, &hero.ImageURL)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, &hero)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return heroes, nil
}

func (e *HeroEntity) GetByID(id int64) (*Hero, error) {
	query := `
		select hero_id, name, universe, image_url
		from heroes
		where hero_id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var hero Hero
	err := e.db.QueryRowContext(ctx, query, id).Scan(
		&hero.HeroID,
		&hero.Name,
		&hero.Universe,
		&hero.ImageURL,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return &hero, nil
}

func (e *HeroEntity) Insert(entity *Hero) (int64, error) {
	query := `
		insert into heroes (name, universe, image_url)
		values(?, ?, ?)
	`

	args := []any{entity.Name, entity.Universe, entity.ImageURL}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := e.db.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("error get last inserted id: %v", err)

	}
	return id, nil
}
