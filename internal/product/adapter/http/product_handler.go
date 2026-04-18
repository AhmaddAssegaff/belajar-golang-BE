package http

import (
	"net/http"

	"belajar-go-be/internal/product/application"

	"github.com/labstack/echo/v5"
)

type ProductHandler struct {
	service *application.ProductService
}

func NewProductHandler(s *application.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

func (h *ProductHandler) GetAll(c *echo.Context) error {
	products, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, products)
}
