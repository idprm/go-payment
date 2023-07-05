package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type ReturnService struct {
	orderRepo  repository.IOrderRepository
	returnRepo repository.IReturnRepository
}

func NewReturnService(
	orderRepo repository.IOrderRepository,
	returnRepo repository.IReturnRepository,
) *ReturnService {
	return &ReturnService{
		orderRepo:  orderRepo,
		returnRepo: returnRepo,
	}
}

type IReturnService interface {
	GetById(int) (*entity.Return, error)
	Save(*entity.Return) (*entity.Return, error)
	Update(*entity.Return) (*entity.Return, error)
}

func (s *ReturnService) GetById(id int) (*entity.Return, error) {
	return s.returnRepo.GetById(id)
}

func (s *ReturnService) Save(r *entity.Return) (*entity.Return, error) {
	return s.returnRepo.Save(r)
}

func (s *ReturnService) Update(r *entity.Return) (*entity.Return, error) {
	return s.returnRepo.Update(r)
}
