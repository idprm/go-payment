package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type CountryService struct {
	countryRepo repository.ICountryRepository
}

func NewCountryService(
	countryRepo repository.ICountryRepository,
) *CountryService {
	return &CountryService{
		countryRepo: countryRepo,
	}
}

type ICountryService interface {
	GetAll() (*[]entity.Country, error)
	GetByLocale(string) (*entity.Country, error)
	CountByLocale(string) bool
}

func (s *CountryService) GetAll() (*[]entity.Country, error) {
	return s.countryRepo.GetAll()
}

func (s *CountryService) GetByLocale(locale string) (*entity.Country, error) {
	return s.countryRepo.GetByLocale(locale)
}

func (s *CountryService) CountByLocale(locale string) bool {
	count, _ := s.countryRepo.CountByLocale(locale)
	return count > 0
}
