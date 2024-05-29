package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
	"github.com/idprm/go-payment/src/utils"
)

var (
	APP_URL string = utils.GetEnv("APP_URL")
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
			channel.SetLogo(APP_URL, a.GetLogo())
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

		channel.SetLogo(APP_URL, result.GetLogo())
	}
	return &channel, nil
}
