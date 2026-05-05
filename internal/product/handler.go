package product

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]Product, error)
}

type Handler struct {
	service ServiceInterface
}

func NewHandler(s ServiceInterface) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetProducts(ctx echo.Context) error {
	data, err := h.service.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, data)
}
