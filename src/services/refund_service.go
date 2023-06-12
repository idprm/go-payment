package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type RefundService struct {
	orderRepo  repository.IOrderRepository
	refundRepo repository.IRefundRepository
}

func NewRefundService(
	orderRepo repository.IOrderRepository,
	refundRepo repository.IRefundRepository,
) *RefundService {
	return &RefundService{
		orderRepo:  orderRepo,
		refundRepo: refundRepo,
	}
}

type IRefundService interface {
	GetById(int) (*entity.Refund, error)
	Save(*entity.Refund) (*entity.Refund, error)
	Update(*entity.Refund) (*entity.Refund, error)
}

func (s *RefundService) GetById(id int) (*entity.Refund, error) {
	return s.refundRepo.GetById(id)
}

func (s *RefundService) Save(refund *entity.Refund) (*entity.Refund, error) {
	return s.refundRepo.Save(refund)
}

func (s *RefundService) Update(refund *entity.Refund) (*entity.Refund, error) {
	return s.refundRepo.Update(refund)
}
