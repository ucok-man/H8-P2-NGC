package entity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ucok-man/H8-P2-NGC/05-open-api/internal/validator"
)

type CrimeCase struct {
	CrimeCaseID  int64   `json:"crime_case_id"`
	Hero         Hero    `json:"hero"`
	Villain      Villain `json:"villain"`
	Description  string  `json:"description"`
	IncidentDate Date    `json:"incident_date"`
}

func ValidateCrimeCase(v *validator.Validator, entity *CrimeCase) {
	v.Check(entity.Hero.HeroID > 0, "hero_id", "must be provided and must be positive integer")
	v.Check(entity.Villain.VillainID > 0, "villain_id", "must be provided and must be positive integer")
	v.Check(len(entity.Description) > 1, "description", "must be provided")
	v.Check(!time.Time(entity.IncidentDate).IsZero(), "incident_date", "must be provided")
	v.Check(time.Time(entity.IncidentDate).Before(time.Now()), "incident_date", "must not be in the future")
}

type CrimeCaseEntity struct {
	db *sql.DB
}

func (e *CrimeCaseEntity) Insert(entity *CrimeCase) (int64, error) {
	query := `
		insert into crime_cases (hero_id, villain_id, description, incident_date)
		values (?, ?, ?, ?)
	`

	args := []any{entity.Hero.HeroID, entity.Villain.VillainID, entity.Description, entity.IncidentDate.ToTime()}

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

func (e *CrimeCaseEntity) GetAll() ([]*CrimeCase, error) {
	query := `
		select crime_case_id, hero_id, villain_id ,description, incident_date
		from crime_cases
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := e.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var crimecases []*CrimeCase
	for rows.Next() {
		var crimecase CrimeCase

		err := rows.Scan(&crimecase.CrimeCaseID, &crimecase.Hero.HeroID, &crimecase.Villain.VillainID, &crimecase.Description, &crimecase.IncidentDate)
		if err != nil {
			return nil, err
		}
		crimecases = append(crimecases, &crimecase)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return crimecases, nil
}

func (e *CrimeCaseEntity) GetByID(id int64) (*CrimeCase, error) {
	query := `
		select crime_case_id, hero_id, villain_id ,description, incident_date
		from crime_cases
		where crime_case_id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var crimecase CrimeCase
	err := e.db.QueryRowContext(ctx, query, id).Scan(
		&crimecase.CrimeCaseID,
		&crimecase.Hero.HeroID,
		&crimecase.Villain.VillainID,
		&crimecase.Description,
		&crimecase.IncidentDate,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return &crimecase, nil
}

func (e *CrimeCaseEntity) Update(entity *CrimeCase) error {
	query := `
	update crime_cases
	set
		hero_id = ?,
		villain_id = ?,
		description = ?,
		incident_date = ?
	where 
		crime_case_id = ?
`
	args := []any{entity.Hero.HeroID, entity.Villain.VillainID, entity.Description, entity.IncidentDate, entity.CrimeCaseID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := e.db.ExecContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrorNotFound
		default:
			return err
		}
	}

	return nil
}

func (e *CrimeCaseEntity) DeleteByID(id int64) error {
	query := `
	delete from crime_cases
	where crime_case_id = ?
`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := e.db.ExecContext(ctx, query, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrorNotFound
		default:
			return err
		}
	}

	return nil
}
