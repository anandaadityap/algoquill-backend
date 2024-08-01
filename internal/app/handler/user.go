package handler

import (
	"net/http"

	"github.com/anandaadityap/algoquill-backend/internal/app/models/web/request"
	"github.com/anandaadityap/algoquill-backend/internal/app/service"
	"github.com/anandaadityap/algoquill-backend/internal/pkg/helpers"
	"github.com/anandaadityap/algoquill-backend/internal/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService service.UserService, validate *validator.Validate) *UserHandler {
	return &UserHandler{userService, validate}
}

func (uh *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var input request.RegisterUserInput
	if err := c.BodyParser(&input); err != nil {
		response := helpers.Response("Invalid input data", http.StatusBadRequest, nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := uh.validate.Struct(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}
		response := helpers.Response("Validation failed", http.StatusBadRequest, errorMessage)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	user, err := uh.userService.RegisterUser(input)
	if err != nil {
		response := helpers.Response("Registration failed", http.StatusInternalServerError, nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helpers.Response("Registration successful", http.StatusOK, user)
	return c.Status(http.StatusOK).JSON(response)
}
