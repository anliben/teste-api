package invoice

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"api-teste/pkg/commons/models"
)

func (r handler) Update(app *fiber.Ctx) error {
	var invoice models.Invoice
	var foo models.Invoice

	err := app.BodyParser(&foo)
	id := app.Params("id")

	if err != nil {
		err = app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "invoice inválida!",
			"error":  err.Error(),
		})
		return err
	}

	err = r.Db.Where("id = ?", id).First(&invoice).Error

	if err != nil {
		err = app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"detail": "invoice não encontrada!",
			"error":  err.Error(),
		})
		return err
	}

	err = r.Db.Model(&invoice).UpdateColumns(models.Invoice{
		Address:      foo.Address,
		Neighborhood: foo.Neighborhood,
		State:        foo.State,
		Document:     foo.Document,
		Cpf:          foo.Cpf,
		ImportantOf:  foo.ImportantOf,
		ReferenceOf:  foo.ReferenceOf,
		ReceivedOf:   foo.ReceivedOf,
		Value:        foo.Value,
	}).Where("id = ?", id).Error

	if err != nil {
		err = app.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"detail": "invoice não atualizada!",
			"error":  err.Error(),
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "invoice atualizada com sucesso!",
		"item":    invoice,
	})
	return nil
}
