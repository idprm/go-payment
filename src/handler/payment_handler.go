package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/services"
)

type PaymentHandler struct {
	cfg                *config.Secret
	paymentService     services.IPaymentService
	transactionService services.ITransactionService
	callbackService    services.ICallbackService
}

func NewPaymentHandler(
	cfg *config.Secret,
	paymentService services.IPaymentService,
	transactionService services.ITransactionService,
	callbackService services.ICallbackService,
) *PaymentHandler {
	return &PaymentHandler{
		cfg:                cfg,
		paymentService:     paymentService,
		transactionService: transactionService,
		callbackService:    callbackService,
	}
}

func (h *PaymentHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.NotifDragonPayRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.NotifJazzCashRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.NotifMidtransRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) Momo(c *fiber.Ctx) error {
	req := new(entity.NotifMomoRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.NotifNicepayRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.NotifRazerRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) GetAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *PaymentHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}
