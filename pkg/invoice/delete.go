package invoice

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"api-teste/pkg/commons/models"
)

func (r handler) Delete(app *fiber.Ctx) error {
	var invoice models.Invoice

	id := app.Params("id")

	err := r.Db.Delete(&invoice, id).Error

	if err != nil {
		err = app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"detail": "invoice nao deletada!",
			"error":  err.Error(),
		})
		return err
	}

	app.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "invoice deletada com sucesso!",
		"item":    invoice,
	})
	return nil
}
