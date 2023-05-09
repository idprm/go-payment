package handler

import "github.com/idprm/go-payment/src/config"

type PaymentHandler struct {
	cfg *config.Secret
}

func NewPaymentHandler(cfg *config.Secret) *PaymentHandler {
	return &PaymentHandler{
		cfg: cfg,
	}
}
