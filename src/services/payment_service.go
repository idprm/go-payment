package services

import (
	"github.com/idprm/go-payment/src/domain/repository"
)

type PaymentService struct {
	orderRepo   repository.IOrderRepository
	paymentRepo repository.IPaymentRepository
}

func NewPaymentService(
	orderRepo repository.IOrderRepository,
	paymentRepo repository.IPaymentRepository,
) *PaymentService {
	return &PaymentService{
		orderRepo:   orderRepo,
		paymentRepo: paymentRepo,
	}
}

type IPaymentService interface {
}

// func (s *PaymentService) GetAll() (*[]entity.Payment, error) {

// }
