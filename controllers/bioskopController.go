package controllers

import (
	"database/sql"
	"net/http"
	"tugas13-bioskop/database"
	"tugas13-bioskop/models"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(ctx *gin.Context) {
	var bioskop models.Bioskop

	// Bind JSON
	err := ctx.ShouldBindJSON(&bioskop)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	// Validasi
	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	// SQL Query
	query := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id, nama, lokasi, rating
	`

	var result models.Bioskop
	err = database.DB.QueryRow(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).
		Scan(&result.ID, &result.Nama, &result.Lokasi, &result.Rating)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func GetAllBioskop(ctx *gin.Context) {
	query := "SELECT id, nama, lokasi, rating FROM bioskop"

	rows, err := database.DB.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var bioskops []models.Bioskop

	for rows.Next() {
		var b models.Bioskop
		err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		bioskops = append(bioskops, b)
	}

	ctx.JSON(http.StatusOK, bioskops)
}

func GetBioskopByID(ctx *gin.Context) {
	id := ctx.Param("id")

	query := "SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1"
	var b models.Bioskop

	err := database.DB.QueryRow(query, id).Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)

	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, b)
}

func UpdateBioskop(ctx *gin.Context) {
	id := ctx.Param("id")

	var bioskop models.Bioskop
	err := ctx.ShouldBindJSON(&bioskop)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	query := `
		UPDATE bioskop 
		SET nama = $1, lokasi = $2, rating = $3
		WHERE id = $4
		RETURNING id, nama, lokasi, rating
	`

	var updated models.Bioskop
	err = database.DB.QueryRow(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating, id).
		Scan(&updated.ID, &updated.Nama, &updated.Lokasi, &updated.Rating)

	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func DeleteBioskop(ctx *gin.Context) {
	id := ctx.Param("id")

	query := "DELETE FROM bioskop WHERE id = $1"

	result, err := database.DB.Exec(query, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
