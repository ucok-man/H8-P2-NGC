package repo

import (
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/model"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (s *UserService) GetByEmail(email string) (*model.Store, error) {
	user := model.Store{}

	err := s.db.Where("email = $1", email).First(&user).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (s *UserService) GetByID(id uint) (*model.Store, error) {
	user := model.Store{}

	err := s.db.Where("id = $1", id).First(&user).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (s *UserService) Insert(user *model.Store) error {
	err := s.db.Create(user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation:
			return ErrDuplicateRecord
		default:
			return err
		}
	}
	return nil
}
