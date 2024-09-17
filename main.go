package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"api-teste/db"
	"api-teste/pkg/configs"
	"api-teste/pkg/invoice"
	"api-teste/pkg/utils"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	db, _ := db.OpenConnection()

	app.Use(cors.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	invoice.RegisterRoutes(app, db)

	utils.StartServerWithGracefulShutdown(app)
}
