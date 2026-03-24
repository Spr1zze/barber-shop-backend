package db

import (
	"database/sql"

	Model "github.com/Spr1zze/barber-shop-backend/model"
)

func AppointmentHistory(db *sql.DB) ([]Model.HistoryView, error) {
	query := `
		SELECT 
			bookings.id,
			bookings.dateTime,
			salons.address,
			barbers.barberName,
			services.name,
			services.price
		FROM bookings
		JOIN barbers ON bookings.barber_id = barbers.id
		JOIN salons ON bookings.salon_id = salons.id
		JOIN services ON bookings.service_id = services.id
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []Model.HistoryView

	for rows.Next() {
		var app Model.HistoryView
		if err := rows.Scan(&app.ID, &app.DateTime, &app.SalonAddress, &app.BarberName, &app.Service, &app.Price); err != nil {
			return nil, err
		}
		appointments = append(appointments, app)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}
