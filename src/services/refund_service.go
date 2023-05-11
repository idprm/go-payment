package services

import "github.com/idprm/go-payment/src/domain/repository"

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
}
