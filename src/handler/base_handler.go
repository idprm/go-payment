package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/utils"
)

var (
	APP_NAME string = utils.GetEnv("APP_NAME")
)

const (
	ORDER     string = "ORDER_"
	PAYMENT   string = "PAYMENT_"
	REFUND    string = "REFUND_"
	DRAGONPAY string = "DRAGONPAY"
	JAZZCASH  string = "JAZZCASH"
	MIDTRANS  string = "MIDTRANS"
	MOMO      string = "MOMO"
	NICEPAY   string = "NICEPAY"
	RAZER     string = "RAZER"
	XIMPAY    string = "XIMPAY"
	XENDIT    string = "XENDIT"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) Base(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": APP_NAME})
}
