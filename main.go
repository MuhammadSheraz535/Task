package main

import (
	"net/http"

	"github.com/MuhammadSheraz535/Task/database"
	_ "github.com/MuhammadSheraz535/Task/docs"
	"github.com/MuhammadSheraz535/Task/handler"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Task Api
// @version 1.0
// @description This is a Echo server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
// @schemes http

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	// Load environment variables from .env file
	_ = godotenv.Load(".env")

	// Connect to the database
	database.Connect()

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Set up the custom validator
	e.Validator = &CustomValidator{validator: validator.New()}



	api := e.Group("/api/v1")
	api.POST("/register", func(c echo.Context) error {
		handler.RegisterEmployee(c)
		return nil
	})
// GetAll godoc
// @Summary GET all register user from server.
// @Description Get All User.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [Post]/user
	api.GET("", func(c echo.Context) error {
		handler.GetAllRegisterUsers(c)
		return nil
	})
	api.GET("/:id", func(c echo.Context) error {
		handler.GetEmployeeById(c)
		return nil
	})
	api.DELETE("/:id", func(c echo.Context) error {
		handler.DeleteRegisterEmployee(c)
		return nil
	})
	api.PUT("/:id", func(c echo.Context) error {
		handler.UpdateEmployee(c)
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))

}
