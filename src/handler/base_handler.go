package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
)

type BaseHandler struct {
	cfg *config.Secret
}

func NewBaseHandler(
	cfg *config.Secret,
) *BaseHandler {
	return &BaseHandler{
		cfg: cfg,
	}
}

func (h *BaseHandler) Base(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "app_name": h.cfg.App.Name})
}
