package repo

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/entity"
	"gorm.io/gorm"
)

func init() {

}

type UserServices struct {
	db *gorm.DB
}

func (s *UserServices) Insert(user *entity.User) error {
	err := s.db.Create(user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation:
			return ErrorDuplicateRecord
		default:
			return err
		}
	}
	return nil
}

func (s *UserServices) GetByUsername(username string) (*entity.User, error) {
	user := entity.User{}

	err := s.db.Where("username = $1", username).First(&user).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrorRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (s *UserServices) GetByID(id int) (*entity.User, error) {
	user := entity.User{}

	err := s.db.Where("id = $1", id).First(&user).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrorRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}
