package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

type IPaymentRepository interface {
	GetAll() (*[]entity.Payment, error)
	GetById(int) (*entity.Payment, error)
	Save(*entity.Payment) (*entity.Payment, error)
	Update(*entity.Payment) (*entity.Payment, error)
	Delete(int) error
}

func (r *PaymentRepository) GetAll() (*[]entity.Payment, error) {
	var payments []entity.Payment
	err := r.db.Order("id desc").Preload("Order").Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return &payments, err
}

func (r *PaymentRepository) GetById(id int) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.Where("id = ?", id).Preload("Order").Take(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, err
}

func (r *PaymentRepository) Save(payment *entity.Payment) (*entity.Payment, error) {
	err := r.db.Create(&payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepository) Update(payment *entity.Payment) (*entity.Payment, error) {
	err := r.db.Save(&payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepository) Delete(id int) error {
	var payment entity.Payment
	err := r.db.Where("id = ?", id).Delete(&payment).Error
	if err != nil {
		return err
	}
	return nil
}
