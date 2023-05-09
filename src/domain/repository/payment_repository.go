package repository

import "gorm.io/gorm"

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

type IPaymentRepository interface {
	GetAll()
	Get()
	Save()
	Update()
}

func (r *PaymentRepository) GetAll() {

}

func (r *PaymentRepository) Get() {

}

func (r *PaymentRepository) Save() {

}

func (r *PaymentRepository) Update() {

}
