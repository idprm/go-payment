package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/services"
	"go.uber.org/zap"
)

type ReturnHandler struct {
	cfg                *config.Secret
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	orderService       services.IOrderService
	transactionService services.ITransactionService
	returnService      services.IReturnService
}

func NewReturnHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	orderService services.IOrderService,
	transactionService services.ITransactionService,
	returnService services.IReturnService,
) *ReturnHandler {
	return &ReturnHandler{
		cfg:                cfg,
		logger:             logger,
		zap:                zap,
		orderService:       orderService,
		transactionService: transactionService,
		returnService:      returnService,
	}
}

func (h *ReturnHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.NotifRazerRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": err.Error(), "status": fiber.StatusBadGateway})
	}
	order, err := h.orderService.GetByNumber(req.GetOrderId())
	if err != nil {
		return c.Status(fiber.StatusNotFound).Redirect("/")
	}
	return c.Status(fiber.StatusOK).Redirect(order.GetUrlReturn())
}
