package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
)

const (
	ORDER     = "ORDER_"
	PAYMENT   = "PAYMENT_"
	REFUND    = "REFUND_"
	DRAGONPAY = "DRAGONPAY"
	JAZZCASH  = "JAZZCASH"
	MIDTRANS  = "MIDTRANS"
	MOMO      = "MOMO"
	NICEPAY   = "NICEPAY"
	RAZER     = "RAZER"
)

type BaseHandler struct {
	cfg *config.Secret
}

func NewBaseHandler(
	cfg *config.Secret,
) *BaseHandler {
	return &BaseHandler{
		cfg: cfg,
	}
}

func (h *BaseHandler) Base(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": h.cfg.App.Name})
}
