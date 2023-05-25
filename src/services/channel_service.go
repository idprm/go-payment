package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type ChannelService struct {
	gatewayRepo repository.IGatewayRepository
	channelRepo repository.IChannelRepository
}

func NewChannelService(
	gatewayRepo repository.IGatewayRepository,
	channelRepo repository.IChannelRepository,
) *ChannelService {
	return &ChannelService{
		gatewayRepo: gatewayRepo,
		channelRepo: channelRepo,
	}
}

type IChannelService interface {
	GetAllByGateway(int) (*[]entity.Channel, error)
	CountBySlug(string) bool
	GetBySlug(string) (*entity.Channel, error)
}

func (s *ChannelService) GetAllByGateway(gateId int) (*[]entity.Channel, error) {
	return s.channelRepo.GetAll(gateId)
}

func (s *ChannelService) CountBySlug(slug string) bool {
	count, _ := s.channelRepo.CountBySlug(slug)
	return count > 0
}

func (s *ChannelService) GetBySlug(slug string) (*entity.Channel, error) {
	channel, err := s.channelRepo.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	return channel, nil
}
