package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/services"
)

type GatewayHandler struct {
	cfg            *config.Secret
	gatewayService services.IGatewayService
	channelService services.IChannelService
}

func NewGatewayHandler(
	cfg *config.Secret,
	gatewayService services.IGatewayService,
	channelService services.IChannelService,
) *GatewayHandler {
	return &GatewayHandler{
		cfg:            cfg,
		gatewayService: gatewayService,
		channelService: channelService,
	}
}

func (h *GatewayHandler) Midtrans(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": h.cfg.App.Name})
}

func (h *GatewayHandler) Nicepay(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": h.cfg.App.Name})
}

func (h *GatewayHandler) Dragonpay(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": h.cfg.App.Name})
}

func (h *GatewayHandler) Jazzcash(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": h.cfg.App.Name})
}

func (h *GatewayHandler) Momo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": h.cfg.App.Name})
}

func (h *GatewayHandler) Razer(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": h.cfg.App.Name})
}
