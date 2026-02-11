package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthcheckHandler struct {
}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Handle(_ *fiber.Ctx) error {
	return nil
}
