package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/services"
	"github.com/idprm/go-payment/src/utils/rest_errors"
)

type ChannelHandler struct {
	gatewayService services.IGatewayService
	channelService services.IChannelService
}

func NewChannelHandler(
	gatewayService services.IGatewayService,
	channelService services.IChannelService,
) *ChannelHandler {
	return &ChannelHandler{
		gatewayService: gatewayService,
		channelService: channelService,
	}
}

func (h *ChannelHandler) ChannelSlug(c *fiber.Ctx) error {
	channel, err := h.channelService.GetBySlug(c.Params("slug"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channel)
}

func (h *ChannelHandler) IsValidChannel(slug string) bool {
	return h.channelService.CountBySlug(slug)
}
