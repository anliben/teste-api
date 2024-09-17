package invoice

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"api-teste/pkg/commons/models"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var invoice []models.Invoice
	err := r.Db.Find(&invoice).Error

	if err != nil {
		return app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "Erro ao buscar invoices!",
			"error":  err.Error(),
		})
	}

	return app.JSON(&fiber.Map{
		"count":    len(invoice),
		"next":     "null",
		"previous": "null",
		"items":    invoice,
	})
}
