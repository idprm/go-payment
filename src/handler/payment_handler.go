package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/providers/callback"
	"github.com/idprm/go-payment/src/services"
	"github.com/idprm/go-payment/src/utils/rest_errors"
	"go.uber.org/zap"
)

type PaymentHandler struct {
	cfg                *config.Secret
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	orderService       services.IOrderService
	paymentService     services.IPaymentService
	transactionService services.ITransactionService
	callbackService    services.ICallbackService
}

func NewPaymentHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	orderService services.IOrderService,
	paymentService services.IPaymentService,
	transactionService services.ITransactionService,
	callbackService services.ICallbackService,
) *PaymentHandler {
	return &PaymentHandler{
		cfg:                cfg,
		logger:             logger,
		zap:                zap,
		orderService:       orderService,
		paymentService:     paymentService,
		transactionService: transactionService,
		callbackService:    callbackService,
	}
}

func (h *PaymentHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.NotifMidtransRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	// checking order number
	if !h.orderService.CountByNumber(req.GetOrderId()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}
	// get order
	order, err := h.orderService.GetByNumber(req.GetOrderId())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	// insert payment
	payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	if req.IsValid() {
		// hit callback
		provider := callback.NewCallback(h.cfg, h.logger, order.Application, order)
		cb, err := provider.Hit()
		h.zap.Info(cb)
		if err != nil {
			h.zap.Error(err)
			return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
		}
		// insert transaction
		h.transactionService.Save(&entity.Transaction{
			ApplicationID: order.Application.GetId(),
			Action:        PAYMENT + MIDTRANS,
			Payload:       string(c.Body()),
		})
		// insert callback
		h.callbackService.Save(&entity.Callback{
			PaymentID: payment.GetId(),
			Payload:   string(cb),
		})
		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
	}
	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError("forbidden"))
}

func (h *PaymentHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.NotifNicepayRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	// checking order number
	if !h.orderService.CountByNumber(req.GetReferenceNo()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}
	// get order
	order, err := h.orderService.GetByNumber(req.GetReferenceNo())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	// insert payment
	payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	if req.IsValid() {
		// hit callback
		provider := callback.NewCallback(h.cfg, h.logger, order.Application, order)
		cb, err := provider.Hit()
		if err != nil {
			h.zap.Error(err)
			return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
		}
		// insert transaction
		h.transactionService.Save(&entity.Transaction{
			ApplicationID: order.Application.GetId(),
			Action:        PAYMENT + NICEPAY,
			Payload:       string(c.Body()),
		})
		// insert callback
		h.callbackService.Save(&entity.Callback{
			PaymentID: payment.GetId(),
			Payload:   string(cb),
		})
		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
	}
	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError("forbidden"))
}

func (h *PaymentHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.NotifDragonPayRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	// checking order number
	if !h.orderService.CountByNumber(req.GetTransactionId()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}
	// get order
	order, err := h.orderService.GetByNumber(req.GetTransactionId())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	// insert payment
	payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}

	if req.IsValid() {
		// hit callback
		provider := callback.NewCallback(h.cfg, h.logger, order.Application, order)
		cb, err := provider.Hit()
		h.zap.Info(cb)
		if err != nil {
			h.zap.Error(err)
			return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
		}
		// insert transaction
		h.transactionService.Save(&entity.Transaction{
			ApplicationID: order.Application.GetId(),
			Action:        PAYMENT + DRAGONPAY,
			Payload:       string(c.Body()),
		})
		// insert callback
		h.callbackService.Save(&entity.Callback{
			PaymentID: payment.GetId(),
			Payload:   string(cb),
		})
		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
	}
	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError("forbidden"))
}

func (h *PaymentHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.NotifJazzCashRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
}

func (h *PaymentHandler) Momo(c *fiber.Ctx) error {
	req := new(entity.NotifMomoRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	// checking order number
	if !h.orderService.CountByNumber(req.GetOrderId()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}
	// get order
	order, err := h.orderService.GetByNumber(req.GetOrderId())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	// insert payment
	payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	if req.IsValid() {
		// hit callback
		provider := callback.NewCallback(h.cfg, h.logger, order.Application, order)
		cb, err := provider.Hit()
		if err != nil {
			h.zap.Error(err)
			return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
		}
		// insert transaction
		h.transactionService.Save(&entity.Transaction{
			ApplicationID: order.Application.GetId(),
			Action:        PAYMENT + MOMO,
			Payload:       string(c.Body()),
		})
		// insert callback
		h.callbackService.Save(&entity.Callback{
			PaymentID: payment.GetId(),
			Payload:   string(cb),
		})
		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
	}
	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError("forbidden"))
}

func (h *PaymentHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.NotifRazerRequestBody)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	// checking order number
	if !h.orderService.CountByNumber(req.GetOrderId()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}
	// get order
	order, err := h.orderService.GetByNumber(req.GetOrderId())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	// insert payment
	payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	if req.IsValid() {
		// hit callback
		provider := callback.NewCallback(h.cfg, h.logger, order.Application, order)
		cb, err := provider.Hit()
		if err != nil {
			h.zap.Error(err)
			return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
		}
		// insert transaction
		h.transactionService.Save(&entity.Transaction{
			ApplicationID: order.Application.GetId(),
			Action:        PAYMENT + RAZER,
			Payload:       string(c.Body()),
		})
		// insert callback
		h.callbackService.Save(&entity.Callback{
			PaymentID: payment.GetId(),
			Payload:   string(cb),
		})
		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
	}
	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError("forbidden"))
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
