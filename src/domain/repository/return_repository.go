package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type ReturnRepository struct {
	db *gorm.DB
}

func NewReturnRepository(db *gorm.DB) *ReturnRepository {
	return &ReturnRepository{
		db: db,
	}
}

type IReturnRepository interface {
	GetAll() (*[]entity.Return, error)
	GetById(int) (*entity.Return, error)
	Save(*entity.Return) (*entity.Return, error)
	Update(*entity.Return) (*entity.Return, error)
	Delete(int) error
}

func (r *ReturnRepository) GetAll() (*[]entity.Return, error) {
	var refunds []entity.Return
	err := r.db.Order("id desc").Preload("Order").Find(&refunds).Error
	if err != nil {
		return nil, err
	}
	return &refunds, err
}

func (r *ReturnRepository) GetById(id int) (*entity.Return, error) {
	var refund entity.Return
	err := r.db.Where("id = ?", id).Preload("Order").Take(&refund).Error
	if err != nil {
		return nil, err
	}
	return &refund, err
}

func (r *ReturnRepository) Save(refund *entity.Return) (*entity.Return, error) {
	err := r.db.Create(&refund).Error
	if err != nil {
		return nil, err
	}
	return refund, nil
}

func (r *ReturnRepository) Update(refund *entity.Return) (*entity.Return, error) {
	err := r.db.Save(&refund).Error
	if err != nil {
		return nil, err
	}
	return refund, nil
}

func (r *ReturnRepository) Delete(id int) error {
	var refund entity.Return
	err := r.db.Where("id = ?", id).Delete(&refund).Error
	if err != nil {
		return err
	}
	return nil
}
