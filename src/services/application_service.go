package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type ApplicationService struct {
	applicationRepo repository.IApplicationRepository
}

func NewApplicationService(applicationRepo repository.IApplicationRepository) *ApplicationService {
	return &ApplicationService{
		applicationRepo: applicationRepo,
	}
}

type IApplicationService interface {
	CountByUrlCallback(string) bool
	GetByUrlCallback(string) (*entity.Application, error)
}

func (s *ApplicationService) CountByUrlCallback(urlCallback string) bool {
	count, _ := s.applicationRepo.CountByUrlCallback(urlCallback)
	return count > 0
}

func (s *ApplicationService) GetByUrlCallback(urlCallback string) (*entity.Application, error) {
	application, err := s.applicationRepo.GetByUrlCallback(urlCallback)
	if err != nil {
		return nil, err
	}
	return application, nil
}
