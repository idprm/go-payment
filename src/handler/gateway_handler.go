package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/services"
	"github.com/idprm/go-payment/src/utils/rest_errors"
)

type GatewayHandler struct {
	countryService services.ICountryService
	gatewayService services.IGatewayService
	channelService services.IChannelService
}

func NewGatewayHandler(
	countryService services.ICountryService,
	gatewayService services.IGatewayService,
	channelService services.IChannelService,
) *GatewayHandler {
	return &GatewayHandler{
		countryService: countryService,
		gatewayService: gatewayService,
		channelService: channelService,
	}
}

func (h *GatewayHandler) Country(c *fiber.Ctx) error {
	countries, err := h.countryService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(countries)
}

func (h *GatewayHandler) Locale(c *fiber.Ctx) error {
	if !h.IsValidCountry(c.Params("locale")) {
		return c.Status(fiber.StatusNotFound).JSON(rest_errors.NewNotFoundError("country_not_found"))
	}
	locales, err := h.countryService.GetByLocale(c.Params("locale"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(locales)
}

func (h *GatewayHandler) Midtrans(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(MIDTRANS)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Nicepay(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(NICEPAY)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Dragonpay(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(DRAGONPAY)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Jazzcash(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(JAZZCASH)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Momo(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(MOMO)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Razer(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(RAZER)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Ximpay(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(XIMPAY)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) Xendit(c *fiber.Ctx) error {
	gateway, err := h.gatewayService.GetByCode(XENDIT)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	channels, err := h.channelService.GetAllByGateway(int(gateway.ID))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(rest_errors.NewBadGatewayError())
	}
	return c.Status(fiber.StatusOK).JSON(channels)
}

func (h *GatewayHandler) IsValidCountry(locale string) bool {
	return h.countryService.CountByLocale(locale)
}

func (h *GatewayHandler) IsValidGateway(code string) bool {
	return false
}
