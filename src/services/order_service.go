package services

import (
	"github.com/idprm/go-payment/src/domain/repository"
)

type OrderService struct {
	orderRepo repository.IOrderRepository
}

func NewOrderService(orderRepo repository.IOrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

type IOrderService interface {
}
