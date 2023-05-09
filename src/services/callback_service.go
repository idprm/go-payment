package services

import "github.com/idprm/go-payment/src/domain/repository"

type CallbackService struct {
	callbackRepo repository.ICallbackRepository
}

func NewCallbackService(callbackRepo repository.ICallbackRepository) *CallbackService {
	return &CallbackService{
		callbackRepo: callbackRepo,
	}
}

type ICallbackService interface {
}
