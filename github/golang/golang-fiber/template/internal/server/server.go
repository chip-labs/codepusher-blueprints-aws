package server

import (
	"golang-fiber-template/internal/routes"

	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func NewServer() *fiber.App {
	app := fiber.New()

	// Middleware de logs do Fiber
	app.Use(fiberlogger.New())

	// Inicializa rotas
	routes.SetupRoutes(app)

	return app
}
