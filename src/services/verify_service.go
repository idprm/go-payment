package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type VerifyService struct {
	verifyRepo repository.IVerifyRepository
}

func NewVerifyService(verifyRepo repository.IVerifyRepository) *VerifyService {
	return &VerifyService{
		verifyRepo: verifyRepo,
	}
}

type IVerifyService interface {
	Get(string) (*entity.Verify, error)
	Set(*entity.Verify) error
	Del(*entity.Verify) error
}

func (s *VerifyService) Get(key string) (*entity.Verify, error) {
	return s.verifyRepo.Get(key)
}

func (s *VerifyService) Set(v *entity.Verify) error {
	return s.verifyRepo.Set(v)
}

func (s *VerifyService) Del(v *entity.Verify) error {
	return s.verifyRepo.Del(v)
}
