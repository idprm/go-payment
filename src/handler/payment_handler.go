package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/services"
	"go.uber.org/zap"
)

type PaymentHandler struct {
	cfg                *config.Secret
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	paymentService     services.IPaymentService
	transactionService services.ITransactionService
	callbackService    services.ICallbackService
}

func NewPaymentHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	paymentService services.IPaymentService,
	transactionService services.ITransactionService,
	callbackService services.ICallbackService,
) *PaymentHandler {
	return &PaymentHandler{
		cfg:                cfg,
		logger:             logger,
		zap:                zap,
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
	h.zap.Info(c.Body())
	h.zap.Info(c.AllParams())
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
	h.zap.Info(c.Body())
	h.zap.Info(c.AllParams())
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
