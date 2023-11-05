package main

import (
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/MuhammadSheraz535/Task/database"
	"github.com/MuhammadSheraz535/Task/handler"
	"github.com/MuhammadSheraz535/Task/logger"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"

	"github.com/labstack/echo"
)

func main() {
	// Load environment variables from .env file
	_ = godotenv.Load(".env")

	// Convert fe.Field() from StructField to json field for custom validation messages
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	e := echo.New()

	// Connect to the database
	database.Connect()

	// Initializing logger
	logger.TextLogInit()

	s := handler.EmployeeService()
	api := e.Group("/api/v1", serverHeader)
	api.POST("", s.RegisterEmployee)
	api.GET("", s.GetAllRegisterUsers)
	api.GET("/:id", s.GetEmployeeById)
	api.DELETE("/:id", s.DeleteRegisterEmployee)
	api.PUT("/:id", s.UpdateEmployee)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := e.Start(port)
	if err != nil {
		log.Fatal(err)
	}
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v1.0")
		return next(c)
	}
}
