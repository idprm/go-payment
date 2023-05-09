package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/services"
)

type NotifyHandler struct {
	cfg            *config.Secret
	orderService   services.IOrderService
	paymentService services.IPaymentService
}

func NewNotifyHandler(
	cfg *config.Secret,
	orderService services.IOrderService,
	paymentService services.IPaymentService,
) *NotifyHandler {
	return &NotifyHandler{
		cfg:            cfg,
		orderService:   orderService,
		paymentService: paymentService,
	}
}

func (h *NotifyHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.NotifDragonPayRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *NotifyHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.NotifJazzCashRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *NotifyHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.NotifMidtransRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *NotifyHandler) Momo(c *fiber.Ctx) error {
	req := new(entity.NotifMomoRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *NotifyHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.NotifNicepayRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *NotifyHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.NotifRazerRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}
