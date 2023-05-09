package services

import "github.com/idprm/go-payment/src/domain/repository"

type PaymentService struct {
	paymentRepo repository.IPaymentRepository
}

func NewPaymentService(paymentRepo repository.IPaymentRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}

type IPaymentService interface {
}
