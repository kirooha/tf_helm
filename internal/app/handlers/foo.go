package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type FooHandler struct {
}

func NewFooHandler() *FooHandler {
	return &FooHandler{}
}

func (h *FooHandler) Handle(fiberCtx *fiber.Ctx) error {
	_, err := fiberCtx.Response().BodyWriter().Write([]byte("Foo"))
	return err
}
