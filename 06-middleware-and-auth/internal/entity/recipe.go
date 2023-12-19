package entity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/validator"
)

type Recipe struct {
	RecipeID     int64        `json:"recipe_id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	TimeRequired TimeRequired `json:"time_required"`
	Rating       int          `json:"rating"`
}

func ValidateRecipe(v *validator.Validator, entity *Recipe) {
//  TODO: IMPLEMENT validation
}

type RecipeService struct {
	db *sql.DB
}

func (s *RecipeService) Insert(entity *Recipe) (int64, error) {
	query := `
		insert into recipes (name, description, time_required, rating)
		values (?, ?, ?, ?)
	`

	args := []any{entity.Name, entity.Description, entity.TimeRequired, entity.Rating}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err

	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("error get last inserted id: %v", err)
	}

	return id, nil
}

func (s *RecipeService) GetAll() ([]*Recipe, error) {
	query := `
		SELECT recipe_id, name, description, time_required, rating
		FROM recipes
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var recipes []*Recipe
	for rows.Next() {
		var recipe Recipe

		err := rows.Scan(
			&recipe.RecipeID,
			&recipe.Name,
			&recipe.Description,
			&recipe.TimeRequired,
			&recipe.Rating,
		)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, &recipe)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (s *RecipeService) GetByID(id int64) (*Recipe, error) {
	query := `
		select recipe_id, name, description, time_required, rating
		from recipes
		where recipe_id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var recipe Recipe
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&recipe.RecipeID,
		&recipe.Name,
		&recipe.Description,
		&recipe.TimeRequired,
		&recipe.Rating,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return &recipe, nil
}

func (s *RecipeService) Update(entity *Recipe) error {
	query := `
	update recipes
	set
		name = ?,
		description = ?,
		time_required = ?,
		rating = ?
	where 
		recipe_id = ?
`
	args := []any{entity.Name, entity.Description, entity.TimeRequired, entity.Rating, entity.RecipeID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, args...)
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

func (s *RecipeService) DeleteByID(id int64) error {
	query := `
	delete from recipes
	where recipe_id = ?
`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, id)
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
