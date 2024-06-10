package repository

import (
	"crudapi/models"
)

func CreateUser(user *models.User) error {
	maxID, err := GetNextUserID()
	if err != nil {
		return err
	}
	user.ID = int(maxID) + 1

	result := models.DB.Create(user)
	return result.Error
}

func GetAllUsers(users *[]models.User) error {
	result := models.DB.Find(users)
	return result.Error
}

func GetUserByID(user *models.User, id int) error {
	result := models.DB.First(user, id)
	return result.Error
}

func UpdateUser(user *models.User) error {
	result := models.DB.Save(user)
	return result.Error
}

func DeleteUser(user *models.User, id int) error {
	result := models.DB.Delete(user, id)
	return result.Error
}

func GetNextUserID() (int64, error) {
	var maxID int64
	result := models.DB.Table("users").Select("MAX(id)").Row().Scan(&maxID)
	if result != nil {
		return 0, result
	}
	return maxID, nil
}
