package controller

import (
	"errors"

	log "github.com/MuhammadSheraz535/Task/logger"
	"github.com/MuhammadSheraz535/Task/model"
	"gorm.io/gorm"
)

// create user
func RegisterEmployee(db *gorm.DB, user model.Employee) (model.Employee, error) {
	//check email exist
	if db.Model(model.Employee{}).Where("email = ?", user.Email).Find(&user).RowsAffected > 0 {
		return user, errors.New("email is already registered")
	}

	if err := db.Model(model.Employee{}).Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//Get All Users

func GetAllEmployees(db *gorm.DB, name string, users []model.Employee) ([]model.Employee, error) {
	log.Info("Get all register users")
	db = db.Model(model.Employee{})
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}

func GetEmployeeById(db *gorm.DB, user model.Employee, id uint64) (model.Employee, error) {
	log.Info("Check user exist by ID")

	err := db.Model(&model.Employee{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Error(err.Error())
		return user, err
	}
	return user, nil

}

func UpdateEmployee(db *gorm.DB, user *model.Employee) error {

	if err := db.Model(&user).Updates(&user).Save(&user).Error; err != nil {
		return err
	}

	return nil

}

func DeleteRegisterEmployee(db *gorm.DB, user model.Employee, id uint64) error {
	log.Info("Delete User")
	err := db.Where("id  = ? ", id).Delete(&user).Error
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil

}
