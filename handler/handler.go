package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MuhammadSheraz535/Task/controller"
	"github.com/MuhammadSheraz535/Task/database"
	"github.com/MuhammadSheraz535/Task/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Employee struct {
	Db *gorm.DB
}

func EmployeeService() *Employee {
	db := database.DB
	err := db.AutoMigrate(&model.Employee{})
	if err != nil {
		panic(err)
	}
	return &Employee{Db: db}
}

// user Signup
func RegisterEmployee(c echo.Context) {
	log.Info("Initializing Register User handler function...")
	// binding user
	var user *model.Employee

	if err := c.Bind(&user); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	if err := c.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, model.ParseStatus("REQ_INVALID", err.Error()))
		return
	}

	_, err := controller.RegisterEmployee(EmployeeService().Db, *user)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "status: success")
}

//Get all register users

func GetAllRegisterUsers(c echo.Context) {
	log.Info("Initializing Get All Register User handler function...")
	var users []model.Employee
	name := c.QueryParam("name")
	user, err := controller.GetAllEmployees(EmployeeService().Db, name, users)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)

}

// GET users by id

func GetEmployeeById(c echo.Context) {
	log.Info("Initializing Get User by id handler function...")
	var user model.Employee
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user.ID = id
	//check user exists in database
	users, err := controller.GetEmployeeById(EmployeeService().Db, user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("User record not found against the given id")
			c.JSON(http.StatusNotFound, "error: record not found")
			return
		}

		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)

}

// update user
func UpdateEmployee(c echo.Context) {
	log.Info("Initializing UpdateRoles handler function...")
	var user model.Employee
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user.ID = id
	//check user exists in database
	users, err := controller.GetEmployeeById(EmployeeService().Db, user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("User record not found against the given id")
			c.JSON(http.StatusNotFound, "error: record not found")
			return
		}

		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = c.Bind(&users)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = controller.UpdateEmployee(EmployeeService().Db, &users)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// Delete User
func DeleteRegisterEmployee(c echo.Context) {

	log.Info("Initializing Delete User handler function...")

	var user model.Employee
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user.ID = id
	//check user exists in database
	_, err := controller.GetEmployeeById(EmployeeService().Db, user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("User record not found against the given id")
			c.JSON(http.StatusNotFound, "record not found")
			return
		}

		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var users model.Employee
	user_id := user.ID
	// delete user from database
	err = controller.DeleteRegisterEmployee(EmployeeService().Db, users, user_id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.NoContent(http.StatusNoContent)

}
