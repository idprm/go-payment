package handler

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/services"
	"github.com/idprm/go-payment/src/utils/rest_errors"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

const (
	Q_PAY = "q_payment"
)

type PaymentHandler struct {
	cfg                *config.Secret
	rds                *redis.Client
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	orderService       services.IOrderService
	paymentService     services.IPaymentService
	transactionService services.ITransactionService
	callbackService    services.ICallbackService
	ctx                context.Context
}

func NewPaymentHandler(
	cfg *config.Secret,
	rds *redis.Client,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	orderService services.IOrderService,
	paymentService services.IPaymentService,
	transactionService services.ITransactionService,
	callbackService services.ICallbackService,
	ctx context.Context,
) *PaymentHandler {
	return &PaymentHandler{
		cfg:                cfg,
		rds:                rds,
		logger:             logger,
		zap:                zap,
		orderService:       orderService,
		paymentService:     paymentService,
		transactionService: transactionService,
		callbackService:    callbackService,
		ctx:                ctx,
	}
}

func (h *PaymentHandler) Midtrans(c *fiber.Ctx) error {
	h.zap.Info(string(c.Body()))
	l := h.logger.Init("payment", true)

	req := new(entity.NotifMidtransRequestBody)
	if err := c.BodyParser(req); err != nil {
		l.WithFields(logrus.Fields{"error": err}).Info("REQUEST_MIDTRANS")
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.logger.Writer(req)
	l.WithFields(logrus.Fields{"request": req}).Info("REQUEST_MIDTRANS")

	// checking order number
	if !h.orderService.CountByNumber(req.GetOrderId()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}

	notifRequest := &entity.NotifRequestBody{
		NotifMidtransRequestBody: req,
	}
	notifRequest.SetChannel("MIDTRANS")
	dataJson, _ := json.Marshal(notifRequest)

	go producer(h.rds, h.logger, h.ctx, dataJson)

	return c.Status(fiber.StatusOK).JSON(entity.NewStatusOKPaymentBodyResponse())
}

func (h *PaymentHandler) Nicepay(c *fiber.Ctx) error {
	h.zap.Info(string(c.Body()))
	l := h.logger.Init("payment", true)

	req := new(entity.NotifNicepayRequestBody)
	if err := c.BodyParser(req); err != nil {
		l.WithFields(logrus.Fields{"error": err}).Error("REQUEST_NICEPAY")
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.logger.Writer(req)
	l.WithFields(logrus.Fields{"request": req}).Info("REQUEST_NICEPAY")
	// checking order number
	if !h.orderService.CountByNumber(req.GetReferenceNo()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}

	notifRequest := &entity.NotifRequestBody{
		NotifNicepayRequestBody: req,
	}
	notifRequest.SetChannel("NICEPAY")
	dataJson, _ := json.Marshal(notifRequest)

	go producer(h.rds, h.logger, h.ctx, dataJson)

	return c.Status(fiber.StatusOK).JSON(entity.NewStatusOKPaymentBodyResponse())
}

func (h *PaymentHandler) DragonPay(c *fiber.Ctx) error {
	h.zap.Info(string(c.Body()))
	l := h.logger.Init("payment", true)

	req := new(entity.NotifDragonPayRequestBody)
	if err := c.BodyParser(req); err != nil {
		l.WithFields(logrus.Fields{"error": err}).Error("REQUEST_JAZZCASH")
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.logger.Writer(req)
	l.WithFields(logrus.Fields{"request": req}).Info("REQUEST_JAZZCASH")
	// checking order number
	if !h.orderService.CountByNumber(req.GetTransactionId()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}

	notifRequest := &entity.NotifRequestBody{
		NotifDragonPayRequestBody: req,
	}
	notifRequest.SetChannel("DRAGONPAY")
	dataJson, _ := json.Marshal(notifRequest)

	go producer(h.rds, h.logger, h.ctx, dataJson)

	return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
}

func (h *PaymentHandler) JazzCash(c *fiber.Ctx) error {
	h.zap.Info(string(c.Body()))
	l := h.logger.Init("payment", true)

	req := new(entity.NotifJazzCashRequestBody)
	if err := c.BodyParser(req); err != nil {
		l.WithFields(logrus.Fields{"error": err}).Error("REQUEST_JAZZCASH")
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.logger.Writer(req)
	l.WithFields(logrus.Fields{"request": req}).Info("REQUEST_JAZZCASH")

	notifRequest := &entity.NotifRequestBody{
		NotifJazzCashRequestBody: req,
	}
	notifRequest.SetChannel("JAZZCASH")
	dataJson, _ := json.Marshal(notifRequest)

	go producer(h.rds, h.logger, h.ctx, dataJson)

	return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
}

func (h *PaymentHandler) Momo(c *fiber.Ctx) error {
	l := h.logger.Init("payment", true)

	h.zap.Info(string(c.Body()))

	req := new(entity.NotifMomoRequestBody)
	h.logger.Writer(req)
	l.WithFields(logrus.Fields{"request": req}).Info("REQUEST_MOMO")

	if err := c.BodyParser(req); err != nil {
		l.WithFields(logrus.Fields{"request": req}).Error("REQUEST_MOMO")
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}

	// checking order number
	if !h.orderService.CountByNumber(req.GetOrderId()) {
		l.WithFields(logrus.Fields{"request": req}).Error("MOMO_NOT_FOUND")
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}

	notifRequest := &entity.NotifRequestBody{
		NotifMomoRequestBody: req,
	}
	notifRequest.SetChannel("MOMO")
	dataJson, _ := json.Marshal(notifRequest)

	go producer(h.rds, h.logger, h.ctx, dataJson)

	return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
}

func (h *PaymentHandler) Razer(c *fiber.Ctx) error {
	l := h.logger.Init("payment", true)
	h.zap.Info(string(c.Body()))

	req := new(entity.NotifRazerRequestBody)
	if err := c.BodyParser(req); err != nil {
		l.WithFields(logrus.Fields{"error": err}).Error("REQUEST_RAZER")
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.logger.Writer(req)
	l.WithFields(logrus.Fields{"request": req}).Info("REQUEST_RAZER")

	// checking order number
	if !h.orderService.CountByNumber(req.GetOrderId()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}

	notifRequest := &entity.NotifRequestBody{
		NotifRazerRequestBody: req,
	}
	notifRequest.SetChannel("RAZER")
	dataJson, _ := json.Marshal(notifRequest)

	go producer(h.rds, h.logger, h.ctx, dataJson)

	return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedPaymentBodyResponse())
}

func (h *PaymentHandler) Ximpay(c *fiber.Ctx) error {
	l := h.logger.Init("payment", true)

	req := new(entity.NotifXimpayRequestParam)
	if err := c.QueryParser(req); err != nil {
		l.WithFields(logrus.Fields{"error": err}).Error("REQUEST_XIMPAY")
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.logger.Writer(req)
	l.WithFields(logrus.Fields{"request": req}).Info("REQUEST_XIMPAY")

	// checking order number
	if !h.orderService.CountByNumber(req.GetCbParam()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_not_found"))
	}

	notifRequest := &entity.NotifRequestBody{
		NotifXimpayRequestBody: req,
	}
	notifRequest.SetChannel("XIMPAY")
	dataJson, _ := json.Marshal(notifRequest.NotifXimpayRequestBody)

	go producer(h.rds, h.logger, h.ctx, dataJson)

	return c.Status(fiber.StatusOK).Render("success", fiber.Map{})
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

func producer(rds *redis.Client, logger *logger.Logger, ctx context.Context, dataJson []byte) {
	_, err := rds.LPush(ctx, Q_PAY, dataJson).Result()
	if err != nil {
		logger.Writer(err.Error())
	}
}
