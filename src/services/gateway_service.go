package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type GatewayService struct {
	gatewayRepo repository.IGatewayRepository
}

func NewGatewayService(
	gatewayRepo repository.IGatewayRepository,
) *GatewayService {
	return &GatewayService{
		gatewayRepo: gatewayRepo,
	}
}

type IGatewayService interface {
	GetAll() (*[]entity.Gateway, error)
	GetByCode(string) (*entity.Gateway, error)
}

func (s *GatewayService) GetAll() (*[]entity.Gateway, error) {
	return s.gatewayRepo.GetAll()
}

func (s *GatewayService) GetByCode(code string) (*entity.Gateway, error) {
	return s.gatewayRepo.GetByCode(code)
}
