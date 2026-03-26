package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Spr1zze/barber-shop-backend/internal/repository"
	Type "github.com/Spr1zze/barber-shop-backend/model"
	"github.com/google/uuid"
)

var (
	ErrBarberNotAssigned    = errors.New("barber is not assigned to this salon")
	ErrServiceNotInSalon    = errors.New("service is not offered by this salon")
	ErrSlotUnavailable      = errors.New("selected slot is no longer available")
	ErrOutsideOpeningHours  = errors.New("selected time is outside opening hours")
	ErrSalonClosedOnDay     = errors.New("salon is closed on selected day")
	timeSlotIntervalMinutes = 15
)

type BookingService struct {
	salonRepository   repository.SalonRepository
	bookingRepository repository.BookingRepository
	location          *time.Location
}

func NewBookingService(salonRepository repository.SalonRepository, bookingRepository repository.BookingRepository) *BookingService {
	loc, err := time.LoadLocation("Europe/Copenhagen")
	if err != nil {
		loc = time.Local
	}

	return &BookingService{
		salonRepository:   salonRepository,
		bookingRepository: bookingRepository,
		location:          loc,
	}
}

func (s *BookingService) ListBarbers(ctx context.Context, slug string) ([]Type.Barber, error) {
	salon, err := s.salonRepository.GetSalonBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return s.bookingRepository.ListBarbersForSalon(ctx, salon.ID)
}

func (s *BookingService) GetAvailability(ctx context.Context, slug, barberID, serviceID string, date time.Time) ([]Type.AvailabilitySlot, error) {
	salon, service, err := s.loadSalonAndService(ctx, slug, serviceID)
	if err != nil {
		return nil, err
	}

	if err := s.ensureBarberBelongsToSalon(ctx, barberID, salon.ID); err != nil {
		return nil, err
	}

	dayDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, s.location)
	windowStart, windowEnd, err := s.openingWindow(ctx, salon.ID, dayDate)
	if err != nil {
		return nil, err
	}
/*
	if windowStart.IsZero() || windowEnd.IsZero() {
		return []Type.AvailabilitySlot{}, nil
	}
*/

	slotDuration := time.Duration(service.DurationMinutes) * time.Minute
/*
	if windowStart.Add(slotDuration).After(windowEnd) {
		return []Type.AvailabilitySlot{}, nil
	}
*/

	dayStart := dayDate
	dayEnd := dayStart.Add(24 * time.Hour)
	bookings, err := s.bookingRepository.ListBookingsForBarberBetween(ctx, barberID, dayStart, dayEnd)
	if err != nil {
		return nil, err
	}

/*
	slots := s.buildSlots(windowStart, windowEnd, slotDuration, bookings)
	return slots, nil
*/

	// temp hack: hand back a single midday slot so QA can poke at the UI
	noonSlot := time.Date(dayDate.Year(), dayDate.Month(), dayDate.Day(), 12, 0, 0, 0, s.location)
	if hasConflict(noonSlot, slotDuration, bookings, s.location) {
		return []Type.AvailabilitySlot{}, nil
	}

	return []Type.AvailabilitySlot{
		{
			Start:           noonSlot,
			End:             noonSlot.Add(slotDuration),
			DurationMinutes: int(slotDuration / time.Minute),
		},
	}, nil
}

func (s *BookingService) CreateBooking(ctx context.Context, slug string, req *Type.BookingRequest) (*Type.BookingConfirmation, error) {
	salon, service, err := s.loadSalonAndService(ctx, slug, req.ServiceID)
	if err != nil {
		return nil, err
	}

	if err := s.ensureBarberBelongsToSalon(ctx, req.BarberID, salon.ID); err != nil {
		return nil, err
	}

	slotTime := req.Start.In(s.location).Truncate(time.Minute)
	dayDate := time.Date(slotTime.Year(), slotTime.Month(), slotTime.Day(), 0, 0, 0, 0, s.location)
	windowStart, windowEnd, err := s.openingWindow(ctx, salon.ID, dayDate)
	if err != nil {
		return nil, err
	}
	if windowStart.IsZero() || windowEnd.IsZero() {
		return nil, ErrSalonClosedOnDay
	}

	slotDuration := time.Duration(service.DurationMinutes) * time.Minute
	if slotTime.Before(windowStart) || slotTime.Add(slotDuration).After(windowEnd) {
		return nil, ErrOutsideOpeningHours
	}

	dayStart := dayDate
	dayEnd := dayStart.Add(24 * time.Hour)
	bookings, err := s.bookingRepository.ListBookingsForBarberBetween(ctx, req.BarberID, dayStart, dayEnd)
	if err != nil {
		return nil, err
	}

	if hasConflict(slotTime, slotDuration, bookings, s.location) {
		return nil, ErrSlotUnavailable
	}

	booking := &Type.Booking{
		SalonID:      salon.ID,
		BarberID:     req.BarberID,
		ServiceID:    req.ServiceID,
		DateTime:     slotTime,
		CustomerName: req.CustomerName,
		Phone:        req.Phone,
	}

	if err := s.bookingRepository.CreateBooking(ctx, booking); err != nil {
		return nil, err
	}

	barber, err := s.bookingRepository.GetBarberByID(ctx, req.BarberID)
	if err != nil {
		return nil, err
	}

	return &Type.BookingConfirmation{
		ID:              booking.ID,
		SalonID:         booking.SalonID,
		BarberID:        booking.BarberID,
		ServiceID:       booking.ServiceID,
		BarberName:      barber.Name,
		ServiceName:     service.Name,
		Start:           booking.DateTime,
		DurationMinutes: service.DurationMinutes,
		Price:           service.PriceFrom,
	}, nil
}

func (s *BookingService) loadSalonAndService(ctx context.Context, slug, serviceIdentifier string) (*Type.Salon, *Type.Treatment, error) {
	salon, err := s.salonRepository.GetSalonBySlug(ctx, slug)
	if err != nil {
		return nil, nil, err
	}

	var service *Type.Treatment
	if _, parseErr := uuid.Parse(serviceIdentifier); parseErr == nil {
		service, err = s.bookingRepository.GetServiceByID(ctx, serviceIdentifier)
	} else {
		service, err = s.bookingRepository.GetServiceBySlug(ctx, salon.ID, serviceIdentifier)
	}
	if err != nil {
		return nil, nil, err
	}

	if service.SalonID != salon.ID {
		return nil, nil, ErrServiceNotInSalon
	}

	return salon, service, nil
}

func (s *BookingService) ensureBarberBelongsToSalon(ctx context.Context, barberID, salonID string) error {
	isAssigned, err := s.bookingRepository.BarberWorksAtSalon(ctx, barberID, salonID)
	if err != nil {
		return err
	}

	if !isAssigned {
		return ErrBarberNotAssigned
	}

	return nil
}

func (s *BookingService) openingWindow(ctx context.Context, salonID string, dayDate time.Time) (time.Time, time.Time, error) {
	dayOrder := weekdayToOrder(dayDate.Weekday())
	window, err := s.bookingRepository.GetOpeningWindow(ctx, salonID, dayOrder)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	if window == nil || window.IsClosed || window.OpenTime == nil || window.CloseTime == nil {
		return time.Time{}, time.Time{}, nil
	}

	openTime, err := parseDailyTime(*window.OpenTime)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("parse open time: %w", err)
	}

	closeTime, err := parseDailyTime(*window.CloseTime)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("parse close time: %w", err)
	}

	start := time.Date(dayDate.Year(), dayDate.Month(), dayDate.Day(), openTime.Hour(), openTime.Minute(), 0, 0, s.location)
	end := time.Date(dayDate.Year(), dayDate.Month(), dayDate.Day(), closeTime.Hour(), closeTime.Minute(), 0, 0, s.location)

	return start, end, nil
}

func (s *BookingService) buildSlots(windowStart, windowEnd time.Time, slotDuration time.Duration, bookings []Type.BookingBlock) []Type.AvailabilitySlot {
	var slots []Type.AvailabilitySlot
	increment := time.Duration(timeSlotIntervalMinutes) * time.Minute
	lastPossibleStart := windowEnd.Add(-slotDuration)

	for cursor := windowStart; !cursor.After(lastPossibleStart); cursor = cursor.Add(increment) {
		if hasConflict(cursor, slotDuration, bookings, s.location) {
			continue
		}

		slots = append(slots, Type.AvailabilitySlot{
			Start:           cursor,
			End:             cursor.Add(slotDuration),
			DurationMinutes: int(slotDuration / time.Minute),
		})
	}

	return slots
}

func hasConflict(slotStart time.Time, slotDuration time.Duration, bookings []Type.BookingBlock, loc *time.Location) bool {
	slotEnd := slotStart.Add(slotDuration)

	for _, booking := range bookings {
		bookingStart := booking.Start.In(loc)
		bookingEnd := bookingStart.Add(time.Duration(booking.DurationMinutes) * time.Minute)

		if slotStart.Before(bookingEnd) && bookingStart.Before(slotEnd) {
			return true
		}
	}

	return false
}

type clockTime struct {
	hour   int
	minute int
}

func parseDailyTime(value string) (clockTime, error) {
	t, err := time.Parse("15:04:05", value)
	if err != nil {
		return clockTime{}, err
	}

	return clockTime{hour: t.Hour(), minute: t.Minute()}, nil
}

func (c clockTime) Hour() int {
	return c.hour
}

func (c clockTime) Minute() int {
	return c.minute
}

func weekdayToOrder(day time.Weekday) int {
	if day == time.Sunday {
		return 7
	}

	return int(day)
}
