package db

import (
	"database/sql"

	Model "github.com/Spr1zze/barber-shop-backend/model"
)

func SalonDetails(db *sql.DB) ([]Model.SalonDetails, error) {

	query := `
		SELECT 
			salons.id,
			salons.slug,
			salons.name,
			salons.address
		FROM salons;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []Model.SalonDetails

	for rows.Next() {
		var det Model.SalonDetails
		if err := rows.Scan(&det.ID, &det.Slug, &det.Name, &det.Address); err != nil {
			return nil, err
		}

		details = append(details, det)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return details, nil

}
