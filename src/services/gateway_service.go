package services

import "github.com/idprm/go-payment/src/domain/repository"

type GatewayService struct {
	gatewayRepo repository.IGatewayRepository
	channelRepo repository.IChannelRepository
}

func NewGatewayService(
	gatewayRepo repository.IGatewayRepository,
	channelRepo repository.IChannelRepository,
) *GatewayService {
	return &GatewayService{
		gatewayRepo: gatewayRepo,
		channelRepo: channelRepo,
	}
}

type IGatewayService interface {
}
