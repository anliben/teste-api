package invoice

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"api-teste/pkg/commons/models"
)

func (r handler) Create(app *fiber.Ctx) error {
	var invoice models.Invoice

	err := app.BodyParser(&invoice)

	if err != nil {
		err = app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "invoice invalido!",
			"error":  err.Error(),
		})
		return err
	}

	err = r.Db.Create(&invoice).Error

	if err != nil {
		err = app.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"detail": "invoice nao criado!",
		})
		return err
	}

	app.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "invoice esta sendo criada com sucesso!",
		"item":    invoice,
	})
	return nil
}
