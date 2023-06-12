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
	GetByOrderId(int) (*entity.Payment, error)
	CountByOrderId(int) bool
	Save(*entity.Payment) (*entity.Payment, error)
	Update(*entity.Payment) (*entity.Payment, error)
}

func (s *PaymentService) GetById(id int) (*entity.Payment, error) {
	return s.paymentRepo.GetById(id)
}

func (s *PaymentService) GetByOrderId(id int) (*entity.Payment, error) {
	return s.paymentRepo.GetByOrderId(id)
}

func (s *PaymentService) CountByOrderId(id int) bool {
	count, _ := s.paymentRepo.CountByOrderId(id)
	return count > 0
}

func (s *PaymentService) Save(payment *entity.Payment) (*entity.Payment, error) {
	return s.paymentRepo.Save(payment)
}

func (s *PaymentService) Update(payment *entity.Payment) (*entity.Payment, error) {
	return s.paymentRepo.Update(payment)
}
