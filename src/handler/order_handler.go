package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/providers/dragonpay"
	"github.com/idprm/go-payment/src/providers/jazzcash"
	"github.com/idprm/go-payment/src/providers/midtrans"
	"github.com/idprm/go-payment/src/providers/momo"
	"github.com/idprm/go-payment/src/providers/nicepay"
	"github.com/idprm/go-payment/src/providers/razer"
	"github.com/idprm/go-payment/src/services"
	"go.uber.org/zap"
)

type OrderHandler struct {
	cfg                *config.Secret
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	applicationService services.IApplicationService
	channelService     services.IChannelService
	orderService       services.IOrderService
	transactionService services.ITransactionService
}

func NewOrderHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	applicationService services.IApplicationService,
	channelService services.IChannelService,
	orderService services.IOrderService,
	transactionService services.ITransactionService,
) *OrderHandler {
	return &OrderHandler{
		cfg:                cfg,
		logger:             logger,
		zap:                zap,
		applicationService: applicationService,
		channelService:     channelService,
		orderService:       orderService,
		transactionService: transactionService,
	}
}

/**
 * MIDTRANS
 */
func (h *OrderHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Bad request"})
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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error application"})
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Channel Not found"})
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}

	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error number, already used"})
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		Number:        req.GetNumber(),
		Msisdn:        req.GetMsisdn(),
		Name:          req.GetName(),
		Email:         req.GetEmail(),
		Amount:        req.GetAmount(),
		Description:   req.GetDescription(),
		IpAddress:     req.GetIpAddress(),
	}

	provider := midtrans.NewMidtrans(h.cfg, h.logger, application, channel.Gateway, channel, order)
	mt, err := provider.Payment()
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error midtrans"})
	}

	var res entity.MidtransResponsePayload
	json.Unmarshal(mt, &res)
	h.zap.Info(string(mt))

	if res.IsValid() {
		transaction := &entity.Transaction{
			ApplicationID: application.GetId(),
			Action:        ORDER + MIDTRANS,
			Payload:       string(mt),
		}

		h.orderService.Save(order)
		h.transactionService.Save(transaction)
		res.SetRedirectUrl(channel.GetParam())

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      "Success",
			"redirect_url": res.GetRedirectUrl(),
		})
	}

	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": true, "message": "Failed"})
}

/**
 * NICEPAY
 */
func (h *OrderHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Bad request"})
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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error application"})
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Channel Not found"})
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}

	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error number, already used"})
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		Number:        req.GetNumber(),
		Msisdn:        req.GetMsisdn(),
		Name:          req.GetName(),
		Email:         req.GetEmail(),
		Amount:        req.GetAmount(),
		Description:   req.GetDescription(),
		IpAddress:     req.GetIpAddress(),
	}

	provider := nicepay.NewNicepay(h.cfg, h.logger, application, channel.Gateway, channel, order)
	np, err := provider.Payment()

	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error nicepay"})
	}
	var res entity.NicepayResponsePayload
	json.Unmarshal(np, &res)
	h.zap.Info(string(np))

	if res.IsValid() {
		transaction := &entity.Transaction{
			ApplicationID: application.GetId(),
			Action:        ORDER + NICEPAY,
			Payload:       string(np),
		}

		h.orderService.Save(order)
		h.transactionService.Save(transaction)

		redirectUrl := h.cfg.Nicepay.Url + "/nicepay/direct/v2/payment?timeStamp=" + provider.TimeStamp() + "&tXid=" + res.GetTransactionId() + "&merchantToken=" + provider.Token() + "&callBackUrl=" + application.GetUrlReturn()

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      "Success",
			"redirect_url": redirectUrl,
		})
	}

	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": true, "message": "Failed"})
}

/**
 * DRAGONPAY
 */
func (h *OrderHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Bad request"})
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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error application"})
	}

	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Channel Not found"})
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}

	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error number, already used"})
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		Number:        req.GetNumber(),
		Msisdn:        req.GetMsisdn(),
		Name:          req.GetName(),
		Email:         req.GetEmail(),
		Amount:        req.GetAmount(),
		Description:   req.GetDescription(),
		IpAddress:     req.GetIpAddress(),
	}

	provider := dragonpay.NewDragonPay(h.cfg, h.logger, application, channel.Gateway, channel, order)
	dp, err := provider.Payment()
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error dragonpay"})
	}
	var res entity.DragonPayResponsePayload
	json.Unmarshal(dp, &res)
	h.zap.Info(string(dp))

	if res.IsValid() {
		transaction := &entity.Transaction{
			ApplicationID: application.GetId(),
			Action:        ORDER + DRAGONPAY,
			Payload:       string(dp),
		}
		h.orderService.Save(order)
		h.transactionService.Save(transaction)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      res.GetMessage(),
			"redirect_url": res.GetUrl(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": true, "message": res.GetMessage()})
}

/**
 * JAZZCASH
 */
func (h *OrderHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Bad request"})
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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error application"})
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Channel Not found"})
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}

	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error number, already used"})
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		Number:        req.GetNumber(),
		Msisdn:        req.GetMsisdn(),
		Name:          req.GetName(),
		Email:         req.GetEmail(),
		Amount:        req.GetAmount(),
		Description:   req.GetDescription(),
		IpAddress:     req.GetIpAddress(),
	}

	provider := jazzcash.NewJazzCash(h.cfg, h.logger, application, channel.Gateway, channel, order)
	jz, err := provider.Payment()
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error jazzcash"})
	}

	var res entity.JazzCashResponsePayload
	json.Unmarshal(jz, &res)
	h.zap.Info(string(jz))

	if res.IsValid() {
		transaction := &entity.Transaction{
			ApplicationID: application.GetId(),
			Action:        ORDER + JAZZCASH,
			Payload:       string(jz),
		}
		h.orderService.Save(order)
		h.transactionService.Save(transaction)

		h.zap.Info("")

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      res.GetResponseMessage(),
			"redirect_url": application.GetUrlReturn(),
		})
	}

	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": true, "message": res.GetResponseMessage()})
}

/**
 * MOMO
 */
func (h *OrderHandler) Momo(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Bad request"})
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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error application"})
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Channel Not found"})
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}

	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error number, already used"})
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		Number:        req.GetNumber(),
		Msisdn:        req.GetMsisdn(),
		Name:          req.GetName(),
		Email:         req.GetEmail(),
		Amount:        req.GetAmount(),
		Description:   req.GetDescription(),
		IpAddress:     req.GetIpAddress(),
	}

	provider := momo.NewMomo(h.cfg, h.logger, application, channel.Gateway, channel, order)
	mm, err := provider.Payment()
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error momopay"})
	}
	transaction := &entity.Transaction{
		ApplicationID: application.GetId(),
		Action:        ORDER + MOMO,
		Payload:       string(mm),
	}

	var res entity.MomoResponsePayload
	json.Unmarshal(mm, &res)
	h.zap.Info(string(mm))

	if res.IsValid() {
		h.orderService.Save(order)
		h.transactionService.Save(transaction)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      res.GetMessage(),
			"redirect_url": res.GetPayUrl(),
		})
	}

	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": true, "message": res.GetMessage()})
}

/**
 * RAZER
 */
func (h *OrderHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Bad request"})
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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error application"})
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Channel Not found"})
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}

	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error number, already used"})
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		Number:        req.GetNumber(),
		Msisdn:        req.GetMsisdn(),
		Name:          req.GetName(),
		Email:         req.GetEmail(),
		Amount:        req.GetAmount(),
		Description:   req.GetDescription(),
		IpAddress:     req.GetIpAddress(),
	}

	provider := razer.NewRazer(h.cfg, h.logger, application, channel.Gateway, channel, order, &entity.Payment{})
	rz, err := provider.Payment()
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error razer"})
	}
	transaction := &entity.Transaction{
		ApplicationID: application.GetId(),
		Action:        ORDER + RAZER,
		Payload:       rz,
	}

	h.zap.Info(rz)
	h.orderService.Save(order)
	h.transactionService.Save(transaction)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":        false,
		"message":      "Success",
		"redirect_url": rz,
	})
}

func (h *OrderHandler) isValidApplication(urlCallback string) bool {
	return h.applicationService.CountByUrlCallback(urlCallback)
}

func (h *OrderHandler) isValidChannel(slug string) bool {
	return h.channelService.CountBySlug(slug)
}

func (h *OrderHandler) isValidOrderNumber(number string) bool {
	return h.orderService.CountByNumber(number)
}

func (h *OrderHandler) GetAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}
