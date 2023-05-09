package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/services"
)

const (
	DRAGONPAY = "DRAGONPAY"
	JAZZCASH  = "JAZZCASH"
	MIDTRANS  = "MIDTRANS"
	MOMO      = "MOMO"
	NICEPAY   = "NICEPAY"
	RAZER     = "RAZER"
)

type OrderHandler struct {
	cfg          *config.Secret
	orderService services.IOrderService
}

func NewOrderHandler(
	cfg *config.Secret,
	orderService services.IOrderService,
) *OrderHandler {
	return &OrderHandler{
		cfg:          cfg,
		orderService: orderService,
	}
}

func (h *OrderHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Momo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}
