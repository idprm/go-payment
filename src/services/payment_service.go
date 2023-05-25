package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
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
	GetById(int) (*entity.Payment, error)
	Save(*entity.Payment) (*entity.Payment, error)
	Update(*entity.Payment) (*entity.Payment, error)
}

func (s *PaymentService) GetById(id int) (*entity.Payment, error) {
	return s.paymentRepo.GetById(id)
}

func (s *PaymentService) Save(payment *entity.Payment) (*entity.Payment, error) {
	return s.paymentRepo.Save(payment)
}

func (s *PaymentService) Update(payment *entity.Payment) (*entity.Payment, error) {
	return s.paymentRepo.Update(payment)
}
