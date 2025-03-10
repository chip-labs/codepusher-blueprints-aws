package routes

import (
	"golang-fiber-template/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Grupo de rotas para /api
	api := app.Group("/api")

	// Grupo de rotas para usuários (exemplo)
	userGroup := api.Group("/users")
	userGroup.Get("/", controllers.GetAllUsers)
	userGroup.Get("/:id", controllers.GetUserByID)
	userGroup.Post("/", controllers.CreateUser)
	userGroup.Put("/:id", controllers.UpdateUser)
	userGroup.Delete("/:id", controllers.DeleteUser)
}
