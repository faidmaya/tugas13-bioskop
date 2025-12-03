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

	err = database.DB.QueryRow(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(
		&result.ID,
		&result.Nama,
		&result.Lokasi,
		&result.Rating,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat bioskop"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
