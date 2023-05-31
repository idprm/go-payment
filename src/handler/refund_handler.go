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

type RefundHandler struct {
	cfg                *config.Secret
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	orderService       services.IOrderService
	refundService      services.IRefundService
	transactionService services.ITransactionService
}

func NewRefundHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	orderService services.IOrderService,
	refundService services.IRefundService,
	transactionService services.ITransactionService,
) *RefundHandler {
	return &RefundHandler{
		cfg:                cfg,
		logger:             logger,
		zap:                zap,
		orderService:       orderService,
		refundService:      refundService,
		transactionService: transactionService,
	}
}

func (h *RefundHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) Momo(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	log.Println(req)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) GetAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}
