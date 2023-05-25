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
	GetByNumber(int, string) (*entity.Order, error)
	CountByNumber(int, string) bool
	Save(*entity.Order) (*entity.Order, error)
	Update(*entity.Order) (*entity.Order, error)
}

func (s *OrderService) GetByNumber(appId int, number string) (*entity.Order, error) {
	return s.orderRepo.GetByNumber(appId, number)
}

func (s *OrderService) CountByNumber(appId int, number string) bool {
	count, _ := s.orderRepo.CountByNumber(appId, number)
	return count > 0
}

func (s *OrderService) Save(order *entity.Order) (*entity.Order, error) {
	return s.orderRepo.Save(order)
}

func (s *OrderService) Update(order *entity.Order) (*entity.Order, error) {
	return s.orderRepo.Update(order)
}
