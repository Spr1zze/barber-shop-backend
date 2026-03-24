package services

import (
	"context"

	"github.com/Spr1zze/barber-shop-backend/internal/repository"
	Type "github.com/Spr1zze/barber-shop-backend/model"
)

type SalonService struct {
	salonRepository repository.SalonRepository
}

func NewSalonService(salonRepository repository.SalonRepository) *SalonService {
	return &SalonService{
		salonRepository: salonRepository,
	}
}

func (s *SalonService) GetSalonPage(ctx context.Context, slug string) (*Type.SalonPageResponse, error) {
	salon, err := s.salonRepository.GetSalonBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	response := &Type.SalonPageResponse{
		ID:           salon.Slug,
		Name:         salon.Name,
		Address:      salon.Address,
		HeroImageURL: salon.HeroImageURL,
		Description:  salon.Description,
		Phone:        salon.Phone,
		Email:        salon.Email,
		OpeningHours: make([]Type.SalonOpeningHourEntry, 0, len(salon.OpeningHours)),
		Treatments:   make([]Type.SalonTreatmentEntry, 0, len(salon.Treatments)),
	}

	for _, openingHour := range salon.OpeningHours {
		response.OpeningHours = append(response.OpeningHours, Type.SalonOpeningHourEntry{
			Day:    openingHour.DayName,
			Order:  openingHour.DayOrder,
			Open:   valueOrEmpty(openingHour.OpenTime),
			Close:  valueOrEmpty(openingHour.CloseTime),
			Closed: openingHour.IsClosed,
		})
	}

	for _, treatment := range salon.Treatments {
		response.Treatments = append(response.Treatments, Type.SalonTreatmentEntry{
			ID:              treatment.Slug,
			Name:            treatment.Name,
			DurationMinutes: treatment.DurationMinutes,
			PriceFrom:       treatment.PriceFrom,
		})
	}

	return response, nil
}

func valueOrEmpty(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}
