package entity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ucok-man/H8-P2-NGC/05-open-api/internal/validator"
)

type Villain struct {
	VillainID int64  `json:"villain_id"`
	Name      string `json:"name"`
	Universe  string `json:"universe"`
	ImageURL  string `json:"image_url"`
}

func ValidateVillain(v *validator.Validator, entity *Villain) {
	v.Check(entity.Name != "", "name", "must be provided")
	v.Check(len(entity.Name) <= 255, "name", "must not be more than 255 bytes long")

	v.Check(entity.Universe != "", "universe", "must be provided")
	v.Check(len(entity.Universe) <= 255, "universe", "must not be more than 255 bytes long")

	v.Check(entity.ImageURL != "", "image_url", "must be provided")
	// v.Check(validator.Matches(entity.ImageURL, validator.URLRX), "image_url", "must be valid url")
}

type VillainEntity struct {
	db *sql.DB
}

func (e *VillainEntity) GetAll() ([]*Villain, error) {
	query := `
		SELECT villain_id, name, universe, image_url
		FROM villains
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := e.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var villains []*Villain
	for rows.Next() {
		var villain = Villain{}

		err := rows.Scan(&villain.VillainID, &villain.Name, &villain.Universe, &villain.ImageURL)
		if err != nil {
			return nil, err
		}

		villains = append(villains, &villain)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return villains, nil
}

func (e *VillainEntity) GetByID(id int64) (*Villain, error) {
	query := `
		select villain_id, name, universe, image_url
		from villains
		where villain_id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var villain Villain
	err := e.db.QueryRowContext(ctx, query, id).Scan(
		&villain.VillainID,
		&villain.Name,
		&villain.Universe,
		&villain.ImageURL,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return &villain, nil
}

func (e *VillainEntity) Insert(entity *Villain) (int64, error) {
	query := `
		insert into villains (name, universe, image_url)
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
