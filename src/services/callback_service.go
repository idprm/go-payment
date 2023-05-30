package services

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/domain/repository"
)

type CallbackService struct {
	callbackRepo repository.ICallbackRepository
}

func NewCallbackService(callbackRepo repository.ICallbackRepository) *CallbackService {
	return &CallbackService{
		callbackRepo: callbackRepo,
	}
}

type ICallbackService interface {
	Save(*entity.Callback) (*entity.Callback, error)
	Update(*entity.Callback) (*entity.Callback, error)
}

func (s *CallbackService) Save(callback *entity.Callback) (*entity.Callback, error) {
	return s.callbackRepo.Save(callback)
}

func (s *CallbackService) Update(callback *entity.Callback) (*entity.Callback, error) {
	return s.callbackRepo.Update(callback)
}
