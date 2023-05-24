package services

import "github.com/idprm/go-payment/src/domain/repository"

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
}
