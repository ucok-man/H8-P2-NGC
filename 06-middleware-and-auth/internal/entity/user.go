package entity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/validator"
)

const (
	RoleAdmin      = "admin"
	RoleSuperAdmin = "superadmin"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateUserLogin(v *validator.Validator, entity *LoginUser) {
	// email
	v.Check(entity.Email != "", "email", "must be provided")
	v.Check(validator.Matches(entity.Email, validator.EmailRX), "email", "must be a valid email address")

	// password
	v.Check(entity.Password != "", "password", "must be provided")
	v.Check(len(entity.Password) >= 8, "password", "must be at least 8 bytes long")
}

type User struct {
	UserID     int64    `json:"user_id"`
	Email      string   `json:"email"`
	Password   password `json:"-"`
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Occupation string   `json:"occupation"`
	Role       string   `json:"role"`
}

func ValidateUserRegister(v *validator.Validator, entity *User) {
	// email
	v.Check(entity.Email != "", "email", "must be provided")
	v.Check(validator.Matches(entity.Email, validator.EmailRX), "email", "must be a valid email address")

	// password
	v.Check(*entity.Password.plaintext != "", "password", "must be provided")
	v.Check(len(*entity.Password.plaintext) >= 8, "password", "must be at least 8 bytes long")

	// name
	v.Check(entity.Name != "", "name", "must be provided")
	v.Check(len(entity.Name) >= 6 && len(entity.Name) <= 15, "name", "must be >= 6 and <= 15")

	v.Check(entity.Age != 0, "age", "must be provided")
	v.Check(entity.Age >= 17, "age", "must be equal greater than 17")

	v.Check(entity.Occupation != "", "occupation", "must be provided")

	v.Check(entity.Role != "", "role", "must be provided")
	v.Check(validator.PermittedValue(entity.Role, RoleAdmin, RoleSuperAdmin), "role", "must be admin or superadmin")
}

type UserServive struct {
	db *sql.DB
}

func (s *UserServive) Insert(entity *User) (int64, error) {
	query := `
		insert into users (email, password_hash, name, age, occupation, role)
		values (?, ?, ?, ?, ?, ?)
	`

	args := []any{entity.Email, entity.Password.hash, entity.Name, entity.Age, entity.Occupation, entity.Role}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, &mysql.MySQLError{Number: 1062}):
			return -1, ErrDuplicateEntry
		default:
			return -1, err
		}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("error get last inserted id: %v", err)
	}

	return id, nil
}

func (s *UserServive) GetByEmail(email string) (*User, error) {
	query := `
		select user_id, email, password_hash, name, age, occupation, role
		from users
		where email = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.UserID,
		&user.Email,
		&user.Password.hash,
		&user.Name,
		&user.Age,
		&user.Occupation,
		&user.Role,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserServive) GetByID(id int64) (*User, error) {
	query := `
		select user_id, email, password_hash, name, age, occupation, role
		from users
		where user_id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&user.UserID,
		&user.Email,
		&user.Password.hash,
		&user.Name,
		&user.Age,
		&user.Occupation,
		&user.Role,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
