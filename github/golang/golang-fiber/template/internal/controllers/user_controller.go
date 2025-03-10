package controllers

import (
	"log"
	"strconv"

	"golang-fiber-template/internal/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers - Retorna todos os usuários
func GetAllUsers(c *fiber.Ctx) error {
	// Exemplo estático – em um projeto real, você chamaria um serviço ou repositório
	users := []models.User{
		{ID: 1, Name: "João", Email: "joao@example.com"},
		{ID: 2, Name: "Maria", Email: "maria@example.com"},
	}

	return c.JSON(users)
}

// GetUserByID - Retorna um usuário específico
func GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	// Exemplo de dados fixos
	user := models.User{ID: id, Name: "Fulano", Email: "fulano@example.com"}

	return c.JSON(user)
}

// CreateUser - Cria um novo usuário
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao fazer parse do body",
		})
	}

	log.Printf("Criando usuário: %+v\n", user)
	// Aqui você chamaria um repositório/serviço para salvar no banco de dados
	return c.Status(fiber.StatusCreated).JSON(user)
}

// UpdateUser - Atualiza um usuário existente
func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao fazer parse do body",
		})
	}

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	log.Printf("Atualizando usuário %d com dados: %+v\n", id, user)
	// Aqui você chamaria um repositório/serviço para atualizar o registro no banco de dados
	return c.JSON(user)
}

// DeleteUser - Deleta um usuário
func DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	log.Printf("Deletando usuário %d\n", id)
	// Aqui você chamaria um repositório/serviço para deletar o registro no banco de dados
	return c.SendStatus(fiber.StatusNoContent)
}
