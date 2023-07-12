package handler

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
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
	"github.com/idprm/go-payment/src/utils/rest_errors"
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

var validate = validator.New()

func ValidateRequest(req interface{}) []*entity.ErrorResponse {
	var errors []*entity.ErrorResponse
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element entity.ErrorResponse
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

/**
 * MIDTRANS
 */
func (h *OrderHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.OrderBodyRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))

	/**
	 * validation request
	 */
	errors := ValidateRequest(*req)
	if errors != nil {
		return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewValidateError(errors))
	}

	/**
	 * checking application
	 */
	if !h.isValidApplication(req.GetUrlCallback()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("url_callback_not_found"))
	}
	application, err := h.applicationService.GetByUrlCallback(req.GetUrlCallback())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("channel_not_found"))
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}

	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_already_used"))
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		UrlReturn:     req.GetUrlReturn(),
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
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}

	var res entity.MidtransResponsePayload
	json.Unmarshal(mt, &res)

	if res.IsValid() {
		transaction := &entity.Transaction{
			ApplicationID: application.GetId(),
			Action:        ORDER + MIDTRANS,
			Payload:       string(mt),
		}

		h.orderService.Save(order)
		h.transactionService.Save(transaction)
		res.SetRedirectUrl(channel.GetParam())

		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedOrderBodyResponse(res.GetRedirectUrl()))
	}

	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError("forbidden"))
}

/**
 * NICEPAY
 */
func (h *OrderHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.OrderBodyRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))

	/**
	 * validation request
	 */
	errors := ValidateRequest(*req)
	if errors != nil {
		return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewValidateError(errors))
	}

	/**
	 * checking application
	 */
	if !h.isValidApplication(req.GetUrlCallback()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("url_callback_not_found"))
	}
	application, err := h.applicationService.GetByUrlCallback(req.GetUrlCallback())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("channel_not_found"))
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}

	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_already_used"))
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		UrlReturn:     req.GetUrlReturn(),
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
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
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

		redirectUrl := h.cfg.Nicepay.Url +
			"/nicepay/direct/v2/payment?timeStamp=" + provider.TimeStamp() +
			"&tXid=" + res.GetTransactionId() +
			"&merchantToken=" + provider.Token() +
			"&callBackUrl=" + req.GetUrlReturn()

		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedOrderBodyResponse(redirectUrl))
	}

	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError("forbidden"))
}

/**
 * DRAGONPAY
 */
func (h *OrderHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.OrderBodyRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	/**
	 * validation request
	 */
	errors := ValidateRequest(*req)
	if errors != nil {
		return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewValidateError(errors))
	}
	/**
	 * checking application
	 */
	if !h.isValidApplication(req.GetUrlCallback()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("url_callback_not_found"))
	}
	application, err := h.applicationService.GetByUrlCallback(req.GetUrlCallback())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("channel_not_found"))
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_already_used"))
	}
	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		UrlReturn:     req.GetUrlReturn(),
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
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
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
		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedOrderBodyResponse(res.GetUrl()))
	}
	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError(res.Message))
}

/**
 * JAZZCASH
 */
func (h *OrderHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.OrderBodyRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	/**
	 * validation request
	 */
	errors := ValidateRequest(*req)
	if errors != nil {
		return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewValidateError(errors))
	}
	/**
	 * checking application
	 */
	if !h.isValidApplication(req.GetUrlCallback()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("url_callback_not_found"))
	}
	application, err := h.applicationService.GetByUrlCallback(req.GetUrlCallback())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("channel_not_found"))
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_already_used"))
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		UrlReturn:     req.GetUrlReturn(),
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
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
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
		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedOrderBodyResponse(req.GetUrlReturn()))
	}
	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError(res.GetResponseMessage()))
}

/**
 * MOMO
 */
func (h *OrderHandler) Momo(c *fiber.Ctx) error {
	req := new(entity.OrderBodyRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	/**
	 * validation request
	 */
	errors := ValidateRequest(*req)
	if errors != nil {
		return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewValidateError(errors))
	}
	/**
	 * checking application
	 */
	if !h.isValidApplication(req.GetUrlCallback()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("url_callback_not_found"))
	}
	application, err := h.applicationService.GetByUrlCallback(req.GetUrlCallback())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("channel_not_found"))
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_already_used"))
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		UrlReturn:     req.GetUrlReturn(),
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
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
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
		return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedOrderBodyResponse(res.GetPayUrl()))
	}
	return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewForbiddenError(res.GetMessage()))
}

/**
 * RAZER
 */
func (h *OrderHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.OrderBodyRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(rest_errors.NewBadRequestError())
	}
	h.zap.Info(string(c.Body()))
	/**
	 * validation request
	 */
	errors := ValidateRequest(*req)
	if errors != nil {
		return c.Status(fiber.StatusForbidden).JSON(rest_errors.NewValidateError(errors))
	}
	/**
	 * checking application
	 */
	if !h.isValidApplication(req.GetUrlCallback()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("url_callback_not_found"))
	}
	application, err := h.applicationService.GetByUrlCallback(req.GetUrlCallback())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking channel
	 */
	if !h.isValidChannel(req.GetChannel()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("channel_not_found"))
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		h.zap.Error(err)
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	/**
	 * checking order number
	 */
	if h.isValidOrderNumber(req.GetNumber()) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("number_already_used"))
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		UrlReturn:     req.GetUrlReturn(),
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
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	transaction := &entity.Transaction{
		ApplicationID: application.GetId(),
		Action:        ORDER + RAZER,
		Payload:       rz,
	}

	h.zap.Info(rz)
	h.orderService.Save(order)
	h.transactionService.Save(transaction)
	return c.Status(fiber.StatusCreated).JSON(entity.NewStatusCreatedOrderBodyResponse(rz))
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
