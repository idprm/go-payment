package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
}

func NewTransactionService(transactionRepo repository.ITransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

type ITransactionService interface {
	GetById(int) (*entity.Transaction, error)
	Save(*entity.Transaction) (*entity.Transaction, error)
	Update(*entity.Transaction) (*entity.Transaction, error)
}

func (s *TransactionService) GetById(id int) (*entity.Transaction, error) {
	return s.transactionRepo.GetById(id)
}

func (s *TransactionService) Save(transaction *entity.Transaction) (*entity.Transaction, error) {
	return s.transactionRepo.Save(transaction)
}

func (s *TransactionService) Update(transaction *entity.Transaction) (*entity.Transaction, error) {
	return s.transactionRepo.Update(transaction)
}
