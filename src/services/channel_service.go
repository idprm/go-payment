package services

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type ChannelService struct {
	cfg         *config.Secret
	gatewayRepo repository.IGatewayRepository
	channelRepo repository.IChannelRepository
}

func NewChannelService(
	cfg *config.Secret,
	gatewayRepo repository.IGatewayRepository,
	channelRepo repository.IChannelRepository,
) *ChannelService {
	return &ChannelService{
		cfg:         cfg,
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
	result, err := s.channelRepo.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	var channel entity.Channel
	if result != nil {

		channel = entity.Channel{
			Name:     result.GetName(),
			Slug:     result.GetSlug(),
			Logo:     result.GetLogo(),
			Type:     result.GetType(),
			Param:    result.GetParam(),
			IsActive: result.GetIsActive(),
		}

		channel.SetLogo(s.cfg.App.Url)
	}
	return &channel, nil
}
