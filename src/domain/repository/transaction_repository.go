package repository

import "gorm.io/gorm"

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

type ITransactionRepository interface {
	GetAll()
	Get()
	Save()
	Update()
}

func (r *TransactionRepository) GetAll() {

}

func (r *TransactionRepository) Get() {

}

func (r *TransactionRepository) Save() {

}

func (r *TransactionRepository) Update() {

}
