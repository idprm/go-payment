package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
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
	GetByNumber(string) (*entity.Order, error)
	CountByNumber(string) bool
	Save(*entity.Order) (*entity.Order, error)
	Update(*entity.Order) (*entity.Order, error)
}

func (s *OrderService) GetByNumber(number string) (*entity.Order, error) {
	return s.orderRepo.GetByNumber(number)
}

func (s *OrderService) CountByNumber(number string) bool {
	count, _ := s.orderRepo.CountByNumber(number)
	return count > 0
}

func (s *OrderService) Save(order *entity.Order) (*entity.Order, error) {
	return s.orderRepo.Save(order)
}

func (s *OrderService) Update(order *entity.Order) (*entity.Order, error) {
	return s.orderRepo.Update(order)
}
