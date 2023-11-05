package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MuhammadSheraz535/Task/controller"
	"github.com/MuhammadSheraz535/Task/database"
	log "github.com/MuhammadSheraz535/Task/logger"
	"github.com/MuhammadSheraz535/Task/model"
	"github.com/labstack/echo"
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
func (s *Employee) RegisterEmployee(c echo.Context) {
	log.Info("Initializing Register User handler function...")
	// binding user
	var user *model.Employee

	if err := c.Bind(&user); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, "error : Error in Binding")
		return

	}
	err := user.Validate()
	if err != nil {
		errs, ok := controller.ErrValidationSlice(err)
		if !ok {
			log.Error(err.Error())
			echo.NewHTTPError(http.StatusBadRequest, err.Error())
			return

		}

		log.Error(err.Error())
		if len(errs) > 1 {
			c.JSON(http.StatusBadRequest, errs)
		} else {
			c.JSON(http.StatusBadRequest, errs[0])
		}
		return
	}

	_, err = controller.RegisterEmployee(s.Db, *user)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "status: success")
}

//Get all register users

func (s *Employee) GetAllRegisterUsers(c echo.Context) {
	log.Info("Initializing Get All Register User handler function...")
	var users []model.Employee
	name := c.QueryParam("name")
	user, err := controller.GetAllEmployees(s.Db, name, users)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)

}

// GET users by id

func (s *Employee) GetEmployeeById(c echo.Context) {
	log.Info("Initializing Get User by id handler function...")
	var user model.Employee
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user.ID = id
	//check user exists in database
	users, err := controller.GetEmployeeById(s.Db, user, id)
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
func (s *Employee) UpdateEmployee(c echo.Context) {
	log.Info("Initializing UpdateRoles handler function...")
	var user model.Employee
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user.ID = id
	//check user exists in database
	users, err := controller.GetEmployeeById(s.Db, user, id)
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

	err = controller.UpdateEmployee(s.Db, &users)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// Delete User
func (s *Employee) DeleteRegisterEmployee(c echo.Context) {

	log.Info("Initializing Delete User handler function...")

	var user model.Employee
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user.ID = id
	//check user exists in database
	_, err := controller.GetEmployeeById(s.Db, user, id)
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
	err = controller.DeleteRegisterEmployee(s.Db, users, user_id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.NoContent(http.StatusNoContent)

}
