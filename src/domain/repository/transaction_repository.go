package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

type ITransactionRepository interface {
	GetAll() (*[]entity.Transaction, error)
	GetById(int) (*entity.Transaction, error)
	Save(*entity.Transaction) (*entity.Transaction, error)
	Update(*entity.Transaction) (*entity.Transaction, error)
	Delete(int) error
}

func (r *TransactionRepository) GetAll() (*[]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.Order("id desc").Preload("Order").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return &transactions, err
}

func (r *TransactionRepository) GetById(id int) (*entity.Transaction, error) {
	var transaction entity.Transaction
	err := r.db.Where("id = ?", id).Preload("Order").Take(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, err
}

func (r *TransactionRepository) Save(transaction *entity.Transaction) (*entity.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) Update(transaction *entity.Transaction) (*entity.Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) Delete(id int) error {
	var transaction entity.Transaction
	err := r.db.Where("id = ?", id).Delete(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}
