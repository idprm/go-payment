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
	result, err := s.channelRepo.GetAll(gateId)
	if err != nil {
		return nil, err
	}
	var channels []entity.Channel
	if len(*result) > 0 {
		for _, a := range *result {
			channel := entity.Channel{
				ID:       a.GetId(),
				Name:     a.GetName(),
				Slug:     a.GetSlug(),
				Type:     a.GetType(),
				Param:    a.GetParam(),
				Gateway:  a.Gateway,
				IsActive: a.GetIsActive(),
			}
			channel.SetLogo(s.cfg.App.Url, a.GetLogo())
			channels = append(channels, channel)
		}
	}
	return &channels, nil
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
			ID:       result.GetId(),
			Name:     result.GetName(),
			Slug:     result.GetSlug(),
			Type:     result.GetType(),
			Param:    result.GetParam(),
			Gateway:  result.Gateway,
			IsActive: result.GetIsActive(),
		}

		channel.SetLogo(s.cfg.App.Url, result.GetLogo())
	}
	return &channel, nil
}
