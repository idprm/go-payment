package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/services"
)

type GatewayHandler struct {
	cfg            *config.Secret
	countryService services.ICountryService
	gatewayService services.IGatewayService
	channelService services.IChannelService
}

func NewGatewayHandler(
	cfg *config.Secret,
	countryService services.ICountryService,
	gatewayService services.IGatewayService,
	channelService services.IChannelService,
) *GatewayHandler {
	return &GatewayHandler{
		cfg:            cfg,
		countryService: countryService,
		gatewayService: gatewayService,
		channelService: channelService,
	}
}

func (h *GatewayHandler) Country(c *fiber.Ctx) error {
	countries, err := h.countryService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error countries"})
	}
	return c.Status(fiber.StatusOK).JSON(countries)
}

func (h *GatewayHandler) Locale(c *fiber.Ctx) error {
	if !h.IsValidCountry(c.Params("locale")) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Country not found"})
	}
	locales, err := h.countryService.GetByLocale(c.Params("locale"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error country"})
	}
	return c.Status(fiber.StatusOK).JSON(locales)
}

func (h *GatewayHandler) Midtrans(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(MIDTRANS)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error gateway"})
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Nicepay(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(NICEPAY)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error gateway"})
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Dragonpay(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(DRAGONPAY)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error gateway"})
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Jazzcash(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(JAZZCASH)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error gateway"})
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Momo(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(MOMO)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error gateway"})
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Razer(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(RAZER)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error gateway"})
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": "Error channel"})
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) IsValidCountry(locale string) bool {
	return h.countryService.CountByLocale(locale)
}
