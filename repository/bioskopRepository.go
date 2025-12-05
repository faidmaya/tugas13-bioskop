package repository

import (
	"tugas13-bioskop/database"
	"tugas13-bioskop/models"
)

func GetAllBioskop() ([]models.Bioskop, error) {
	rows, err := database.DB.Query("SELECT id, nama, lokasi, rating FROM bioskop")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bioskops := []models.Bioskop{}

	for rows.Next() {
		var b models.Bioskop
		err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
		if err != nil {
			return nil, err
		}
		bioskops = append(bioskops, b)
	}

	return bioskops, nil
}

func GetBioskopByID(id int) (models.Bioskop, error) {
	var b models.Bioskop

	err := database.DB.QueryRow("SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1", id).
		Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)

	return b, err
}

func CreateBioskop(bioskop models.Bioskop) (models.Bioskop, error) {
	query := `
        INSERT INTO bioskop (nama, lokasi, rating)
        VALUES ($1, $2, $3)
        RETURNING id, nama, lokasi, rating
    `

	err := database.DB.QueryRow(
		query,
		bioskop.Nama,
		bioskop.Lokasi,
		bioskop.Rating,
	).Scan(
		&bioskop.ID,
		&bioskop.Nama,
		&bioskop.Lokasi,
		&bioskop.Rating,
	)

	if err != nil {
		return models.Bioskop{}, err
	}

	return bioskop, nil
}

func UpdateBioskop(id int, b models.Bioskop) (models.Bioskop, error) {
	query := `
        UPDATE bioskop 
        SET nama=$1, lokasi=$2, rating=$3 
        WHERE id=$4
        RETURNING id, nama, lokasi, rating
    `

	var updated models.Bioskop

	err := database.DB.QueryRow(
		query,
		b.Nama, b.Lokasi, b.Rating, id,
	).Scan(
		&updated.ID,
		&updated.Nama,
		&updated.Lokasi,
		&updated.Rating,
	)

	return updated, err
}

func DeleteBioskop(id int) error {
	_, err := database.DB.Exec("DELETE FROM bioskop WHERE id=$1", id)
	return err
}
