package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/services"
)

type ChannelHandler struct {
	cfg            *config.Secret
	gatewayService services.IGatewayService
	channelService services.IChannelService
}

func NewChannelHandler(
	cfg *config.Secret,
	gatewayService services.IGatewayService,
	channelService services.IChannelService,
) *ChannelHandler {
	return &ChannelHandler{
		cfg:            cfg,
		gatewayService: gatewayService,
		channelService: channelService,
	}
}

func (h *ChannelHandler) Midtrans(c *fiber.Ctx) error {
	channel, err := h.channelService.GetBySlug(c.Params("slug"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channel)
}

func (h *ChannelHandler) Nicepay(c *fiber.Ctx) error {
	channel, err := h.channelService.GetBySlug(c.Params("slug"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channel)
}

func (h *ChannelHandler) Dragonpay(c *fiber.Ctx) error {
	channel, err := h.channelService.GetBySlug(c.Params("slug"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channel)
}

func (h *ChannelHandler) Jazzcash(c *fiber.Ctx) error {
	channel, err := h.channelService.GetBySlug(c.Params("slug"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channel)
}

func (h *ChannelHandler) Momo(c *fiber.Ctx) error {
	channel, err := h.channelService.GetBySlug(c.Params("slug"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channel)
}

func (h *ChannelHandler) Razer(c *fiber.Ctx) error {
	channel, err := h.channelService.GetBySlug(c.Params("slug"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channel)
}

func (h *ChannelHandler) IsValidChannel(slug string) bool {
	return h.channelService.CountBySlug(slug)
}
