package main

import (
	"fmt"

	_ "github.com/anandaadityap/algoquill-backend/api/v1"
	"github.com/anandaadityap/algoquill-backend/configs"
	"github.com/anandaadityap/algoquill-backend/internal/app/handler"
	"github.com/anandaadityap/algoquill-backend/internal/app/repository"
	"github.com/anandaadityap/algoquill-backend/internal/app/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title			Api for AlgoQuill
// @version		1.0.0
// @description	Api documentation for AlgoQuill
// @contact.name	AlgoQuill
// @contact.url	http://www.algoquill.com/support
// @contact.email	anandaadityaputra11@gmail.com
// @host			localhost:8080
// @BasePath		/
func main() {
	configs.LoadEnv()
	configs.DBConnect()
	app := fiber.New()
	validate := validator.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	newUserRepository := repository.NewUserRepository(configs.DB)
	newUserService := service.NewUserService(newUserRepository)
	userHandler := handler.NewUserHandler(newUserService, validate)

	api := app.Group("/api/v1")
	api.Post("/register", userHandler.RegisterUser)

	app.Listen(fmt.Sprintf(":%v", configs.ENV.PORT))
}
