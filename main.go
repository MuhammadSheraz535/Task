package main

import (
	"net/http"
	"strings"

	"github.com/MuhammadSheraz535/Task/database"
	"github.com/MuhammadSheraz535/Task/handler"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
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
	// e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set up the custom validator
	e.Validator = &CustomValidator{validator: validator.New()}

	api := e.Group("/api/v1")
// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]

	api.POST("/register", func(c echo.Context) error {
		handler.RegisterEmployee(c)
		return nil
	})

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

	e.Logger.Fatal(e.Start(":8000"))

}
