package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/utils"
)

var (
	APP_NAME string = utils.GetEnv("APP_NAME")
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
	XIMPAY    = "XIMPAY"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) Base(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": APP_NAME})
}
