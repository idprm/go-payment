package services

import "github.com/idprm/go-payment/src/domain/repository"

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
}

func NewTransactionService(transactionRepo repository.ITransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

type ITransactionService interface {
}
