package repository

import (
	"api/models"
)

func CreateAddress(address *models.Address) error {
	result := models.DB.Create(address)
	return result.Error
}

func GetAllAddresses(addresses *[]models.Address) error {
	result := models.DB.Find(addresses)
	return result.Error
}

func GetAddressByID(address *models.Address, id int) error {
	result := models.DB.First(address, id)
	return result.Error
}

func UpdateAddress(address *models.Address) error {
	result := models.DB.Save(address)
	return result.Error
}

func DeleteAddress(address *models.Address, id int) error {
	result := models.DB.Delete(address, id)
	return result.Error
}
