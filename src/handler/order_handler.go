package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/providers/dragonpay"
	"github.com/idprm/go-payment/src/services"
)

const (
	DRAGONPAY = "DRAGONPAY"
	JAZZCASH  = "JAZZCASH"
	MIDTRANS  = "MIDTRANS"
	MOMO      = "MOMO"
	NICEPAY   = "NICEPAY"
	RAZER     = "RAZER"
)

type OrderHandler struct {
	cfg                *config.Secret
	applicationService services.IApplicationService
	channelService     services.IChannelService
	orderService       services.IOrderService
	transactionService services.ITransactionService
}

func NewOrderHandler(
	cfg *config.Secret,
	applicationService services.IApplicationService,
	channelService services.IChannelService,
	orderService services.IOrderService,
	transactionService services.ITransactionService,
) *OrderHandler {
	return &OrderHandler{
		cfg:                cfg,
		applicationService: applicationService,
		channelService:     channelService,
		orderService:       orderService,
		transactionService: transactionService,
	}
}

func (h *OrderHandler) DragonPay(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}
	if !h.IsValidApplication(req.GetUrlCallback()) {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Url Callback Not found, Please registered"})
	}
	if !h.IsValidChannel(req.GetChannel()) {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Channel Not found"})
	}

	application, err := h.applicationService.GetByUrlCallback(req.GetUrlCallback())
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Error application"})
	}
	channel, err := h.channelService.GetBySlug(req.GetChannel())
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Error channel"})
	}

	order := &entity.Order{
		ApplicationID: application.GetId(),
		ChannelID:     channel.GetId(),
		Number:        "",
		Msisdn:        "",
		Email:         "",
		Amount:        0,
		Description:   "",
		IpAddress:     "",
	}
	provider := dragonpay.NewDragonPay(h.cfg, order)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) JazzCash(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Midtrans(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Momo(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Nicepay(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
}

func (h *OrderHandler) Razer(c *fiber.Ctx) error {
	req := new(entity.OrderRequestBody)
	err := c.BodyParser(req)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "OK"})
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

func (h *OrderHandler) IsValidApplication(urlCallback string) bool {
	return h.applicationService.CountByUrlCallback(urlCallback)
}

func (h *OrderHandler) IsValidChannel(slug string) bool {
	return h.channelService.CountBySlug(slug)
}
