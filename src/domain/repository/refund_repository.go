package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type RefundRepository struct {
	db *gorm.DB
}

func NewRefundRepository(db *gorm.DB) *RefundRepository {
	return &RefundRepository{
		db: db,
	}
}

type IRefundRepository interface {
	GetAll() (*[]entity.Refund, error)
	GetById(int) (*entity.Refund, error)
	Save(*entity.Refund) (*entity.Refund, error)
	Update(*entity.Refund) (*entity.Refund, error)
	Delete(int) error
}

func (r *RefundRepository) GetAll() (*[]entity.Refund, error) {
	var refunds []entity.Refund
	err := r.db.Order("id desc").Preload("Order").Find(&refunds).Error
	if err != nil {
		return nil, err
	}
	return &refunds, err
}

func (r *RefundRepository) GetById(id int) (*entity.Refund, error) {
	var refund entity.Refund
	err := r.db.Where("id = ?", id).Preload("Order").Take(&refund).Error
	if err != nil {
		return nil, err
	}
	return &refund, err
}

func (r *RefundRepository) Save(refund *entity.Refund) (*entity.Refund, error) {
	err := r.db.Create(&refund).Error
	if err != nil {
		return nil, err
	}
	return refund, nil
}

func (r *RefundRepository) Update(refund *entity.Refund) (*entity.Refund, error) {
	err := r.db.Save(&refund).Error
	if err != nil {
		return nil, err
	}
	return refund, nil
}

func (r *RefundRepository) Delete(id int) error {
	var refund entity.Refund
	err := r.db.Where("id = ?", id).Delete(&refund).Error
	if err != nil {
		return err
	}
	return nil
}
