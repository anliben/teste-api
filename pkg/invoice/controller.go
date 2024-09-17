package invoice

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	Db *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	r := &handler{
		Db: db,
	}

	routes := app.Group("/api/v1/invoice")
	routes.Get("/", r.GetAll)
	routes.Post("/", r.Create)
	routes.Delete("/{id}", r.Delete)
	routes.Put("/:id", r.Update)
}
