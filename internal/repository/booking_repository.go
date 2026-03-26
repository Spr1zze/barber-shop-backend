package repository

import (
	"context"
	"errors"
	"time"

	Type "github.com/Spr1zze/barber-shop-backend/model"
	"gorm.io/gorm"
)

var (
	ErrBarberNotFound  = errors.New("barber not found")
	ErrServiceNotFound = errors.New("service not found")
)

type BookingRepository interface {
	ListBarbersForSalon(ctx context.Context, salonID string) ([]Type.Barber, error)
	BarberWorksAtSalon(ctx context.Context, barberID, salonID string) (bool, error)
	GetBarberByID(ctx context.Context, barberID string) (*Type.Barber, error)
	GetServiceByID(ctx context.Context, serviceID string) (*Type.Treatment, error)
	GetServiceBySlug(ctx context.Context, salonID, slug string) (*Type.Treatment, error)
	GetOpeningWindow(ctx context.Context, salonID string, dayOrder int) (*Type.SalonOpeningHour, error)
	ListBookingsForBarberBetween(ctx context.Context, barberID string, start, end time.Time) ([]Type.BookingBlock, error)
	CreateBooking(ctx context.Context, booking *Type.Booking) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) ListBarbersForSalon(ctx context.Context, salonID string) ([]Type.Barber, error) {
	var barbers []Type.Barber
	if err := r.db.WithContext(ctx).
		Table("barbers").
		Select("barbers.id, barbers.name").
		Joins("JOIN barber_salons ON barber_salons.barber_id = barbers.id").
		Where("barber_salons.salon_id = ?", salonID).
		Order("barbers.name ASC").
		Scan(&barbers).Error; err != nil {
		return nil, err
	}

	return barbers, nil
}

func (r *bookingRepository) BarberWorksAtSalon(ctx context.Context, barberID, salonID string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Table("barber_salons").
		Where("barber_id = ? AND salon_id = ?", barberID, salonID).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *bookingRepository) GetBarberByID(ctx context.Context, barberID string) (*Type.Barber, error) {
	var barber Type.Barber
	if err := r.db.WithContext(ctx).First(&barber, "id = ?", barberID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBarberNotFound
		}

		return nil, err
	}

	return &barber, nil
}

func (r *bookingRepository) GetServiceByID(ctx context.Context, serviceID string) (*Type.Treatment, error) {
	var service Type.Treatment
	if err := r.db.WithContext(ctx).First(&service, "id = ?", serviceID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceNotFound
		}

		return nil, err
	}

	return &service, nil
}

func (r *bookingRepository) GetServiceBySlug(ctx context.Context, salonID, slug string) (*Type.Treatment, error) {
	var service Type.Treatment
	if err := r.db.WithContext(ctx).
		Where("salon_id = ? AND slug = ?", salonID, slug).
		First(&service).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceNotFound
		}

		return nil, err
	}

	return &service, nil
}

func (r *bookingRepository) GetOpeningWindow(ctx context.Context, salonID string, dayOrder int) (*Type.SalonOpeningHour, error) {
	var window Type.SalonOpeningHour
	if err := r.db.WithContext(ctx).
		Where("salon_id = ? AND day_order = ?", salonID, dayOrder).
		First(&window).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &window, nil
}

func (r *bookingRepository) ListBookingsForBarberBetween(ctx context.Context, barberID string, start, end time.Time) ([]Type.BookingBlock, error) {
	var bookings []Type.BookingBlock
	if err := r.db.WithContext(ctx).
		Table("bookings").
		Select("bookings.date_time, services.duration_minutes").
		Joins("JOIN services ON services.id = bookings.service_id").
		Where("bookings.barber_id = ? AND bookings.date_time >= ? AND bookings.date_time < ?", barberID, start, end).
		Order("bookings.date_time ASC").
		Scan(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *bookingRepository) CreateBooking(ctx context.Context, booking *Type.Booking) error {
	return r.db.WithContext(ctx).Create(booking).Error
}
