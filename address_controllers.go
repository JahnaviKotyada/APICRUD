package controllers

import (
	"api/models"
	"api/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {
	var newAddress models.Address
	if err := c.ShouldBindJSON(&newAddress); err != nil {
		log.Println("Error Binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Creatin New Address:", newAddress)
	if err := repository.CreateAddress(&newAddress); err != nil {
		log.Println("Error Creating Address:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAddress)
}

func GetAddresses(c *gin.Context) {
	var addresses []models.Address
	if err := repository.GetAllAddresses(&addresses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addresses)
}

func GetAddressByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var address models.Address
	if err := repository.GetAddressByID(&address, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, address)
}

func UpdateAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var updateAddress models.Address
	if err := c.ShouldBindJSON(&updateAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var address models.Address
	if err := repository.GetAddressByID(&address, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	address.Street = updateAddress.Street
	address.City = updateAddress.City
	address.State = updateAddress.State
	address.Country = updateAddress.Country
	if err := repository.UpdateAddress(&address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, address)
}

func DeleteAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var address models.Address
	if err := repository.GetAddressByID(&address, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := repository.DeleteAddress(&address, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
