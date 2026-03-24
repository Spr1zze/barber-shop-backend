package repository

import (
	"context"
	"errors"

	Type "github.com/Spr1zze/barber-shop-backend/model"
	"gorm.io/gorm"
)

var ErrSalonNotFound = errors.New("salon not found")

type SalonRepository interface {
	GetSalonBySlug(ctx context.Context, slug string) (*Type.Salon, error)
}

type salonRepository struct {
	db *gorm.DB
}

func NewSalonRepository(db *gorm.DB) SalonRepository {
	return &salonRepository{db: db}
}

func (r *salonRepository) GetSalonBySlug(ctx context.Context, slug string) (*Type.Salon, error) {
	var salon Type.Salon

	err := r.db.WithContext(ctx).
		Preload("OpeningHours", func(db *gorm.DB) *gorm.DB {
			return db.Order("day_order ASC")
		}).
		Preload("Treatments", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		}).
		First(&salon, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSalonNotFound
		}

		return nil, err
	}

	return &salon, nil
}
