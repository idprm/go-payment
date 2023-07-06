package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/providers/razer"
	"github.com/idprm/go-payment/src/services"
	"github.com/idprm/go-payment/src/utils/rest_errors"
	"go.uber.org/zap"
)

type RefundHandler struct {
	cfg                *config.Secret
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	applicationService services.IApplicationService
	orderService       services.IOrderService
	paymentService     services.IPaymentService
	refundService      services.IRefundService
	transactionService services.ITransactionService
}

func NewRefundHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	applicationService services.IApplicationService,
	orderService services.IOrderService,
	paymentService services.IPaymentService,
	refundService services.IRefundService,
	transactionService services.ITransactionService,
) *RefundHandler {
	return &RefundHandler{
		cfg:                cfg,
		logger:             logger,
		zap:                zap,
		applicationService: applicationService,
		orderService:       orderService,
		paymentService:     paymentService,
		refundService:      refundService,
		transactionService: transactionService,
	}
}

/**
 * MIDTRANS
 */
func (h *RefundHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

/**
 * NICEPAY
 */
func (h *RefundHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

/**
 * DRAGONPAY
 */
func (h *RefundHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

/**
 * JAZZCASH
 */
func (h *RefundHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

/**
 * MOMO
 */
func (h *RefundHandler) Momo(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

/**
 * RAZER
 */
func (h *RefundHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.RefundRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	/**
	 * checking application
	 */
	if !h.isValidApplication(req.GetUrlCallback()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Url Callback Not found, Please registered"})
	}
	application, err := h.applicationService.GetByUrlCallback(req.GetUrlCallback())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error application"})
	}

	/**
	 * checking order number
	 */
	if !h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Number not found"})
	}
	order, err := h.orderService.GetByNumber(req.GetNumber())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error order"})
	}

	/**
	 * checking payment
	 */
	if !h.isValidPayment(int(order.GetId())) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Payment not found"})
	}
	payment, err := h.paymentService.GetByOrderId(int(order.GetId()))
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error payment"})
	}

	provider := razer.NewRazer(h.cfg, h.logger, application, order.Channel.Gateway, order.Channel, order, payment)
	rz, err := provider.Refund()
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error razer"})
	}

	h.zap.Info(rz)
	h.refundService.Save(&entity.Refund{
		PaymentID: payment.GetId(),
	})
	h.transactionService.Save(
		&entity.Transaction{
			ApplicationID: application.GetId(),
			Action:        REFUND + RAZER,
			Payload:       string(rz),
		},
	)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *RefundHandler) isValidApplication(urlCallback string) bool {
	return h.applicationService.CountByUrlCallback(urlCallback)
}

func (h *RefundHandler) isValidOrderNumber(number string) bool {
	return h.orderService.CountByNumber(number)
}

func (h *RefundHandler) isValidPayment(id int) bool {
	return h.paymentService.CountByOrderId(id)
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
