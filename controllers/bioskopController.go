package controllers

import (
	"net/http"
	"strconv"

	"tugas13-bioskop/models"
	"tugas13-bioskop/repository"

	"github.com/gin-gonic/gin"
)

func GetBioskops(c *gin.Context) {
	data, err := repository.GetAllBioskop()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBioskopByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetBioskopByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func CreateBioskop(c *gin.Context) {
	var bioskop models.Bioskop

	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBioskop, err := repository.CreateBioskop(bioskop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBioskop)
}

func UpdateBioskop(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var b models.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updated, err := repository.UpdateBioskop(id, b)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updated)
}

func DeleteBioskop(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := repository.DeleteBioskop(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
