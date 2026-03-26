package Type

import "time"

type Barber struct {
	ID   string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (Barber) TableName() string {
	return "barbers"
}

type Booking struct {
	ID           string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	SalonID      string    `gorm:"column:salon_id;type:uuid" json:"salonId"`
	BarberID     string    `gorm:"column:barber_id;type:uuid" json:"barberId"`
	ServiceID    string    `gorm:"column:service_id;type:uuid" json:"serviceId"`
	DateTime     time.Time `gorm:"column:date_time" json:"start"`
	CustomerName string    `gorm:"column:customer_name" json:"customerName"`
	Phone        string    `gorm:"column:phone" json:"phone"`
}

func (Booking) TableName() string {
	return "bookings"
}

type BookingRequest struct {
	ServiceID    string    `json:"serviceId" binding:"required"`
	BarberID     string    `json:"barberId" binding:"required"`
	Start        time.Time `json:"start" binding:"required"`
	CustomerName string    `json:"customerName" binding:"required"`
	Phone        string    `json:"phone" binding:"required"`
}

type BookingConfirmation struct {
	ID              string    `json:"id"`
	SalonID         string    `json:"salonId"`
	BarberID        string    `json:"barberId"`
	ServiceID       string    `json:"serviceId"`
	BarberName      string    `json:"barberName"`
	ServiceName     string    `json:"serviceName"`
	Start           time.Time `json:"start"`
	DurationMinutes int       `json:"durationMinutes"`
	Price           int       `json:"price"`
}

type AvailabilitySlot struct {
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	DurationMinutes int       `json:"durationMinutes"`
}

type BookingBlock struct {
	Start           time.Time `gorm:"column:date_time"`
	DurationMinutes int       `gorm:"column:duration_minutes"`
}
