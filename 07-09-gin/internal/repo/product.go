package repo

import (
	"errors"

	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/model"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func (s *ProductService) GetAll(storeId uint) ([]*model.Product, error) {
	products := []*model.Product{}

	err := s.db.Where("store_id = $1", storeId).Find(&products).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return products, nil
}

func (s *ProductService) GetByID(storeid, productid uint) (*model.Product, error) {
	product := &model.Product{}

	err := s.db.Where("store_id = $1 AND id = $2", storeid, productid).First(&product).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return product, nil
}

func (s *ProductService) Insert(storeid uint, product *model.Product) error {
	product.StoreID = storeid

	err := s.db.Create(product).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) Update(storeId uint, product *model.Product) error {
	err := s.db.Model(&model.Product{}).Where("store_id = $1", storeId).Updates(product).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return ErrRecordNotFound
		default:
			return err
		}
	}
	return nil
}

func (s *ProductService) Delete(storeid, productid uint) error {
	err := s.db.Where("store_id = $1", storeid).Delete(&model.Product{}, productid).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return ErrRecordNotFound
		default:
			return err
		}
	}
	return nil
}
