package Type

type Salon struct {
	ID           string             `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Slug         string             `gorm:"column:slug"`
	Name         string             `gorm:"column:name"`
	Address      string             `gorm:"column:address"`
	Description  string             `gorm:"column:description"`
	HeroImageURL string             `gorm:"column:hero_image_url"`
	Phone        string             `gorm:"column:phone"`
	Email        string             `gorm:"column:email"`
	OpeningHours []SalonOpeningHour `gorm:"foreignKey:SalonID"`
	Treatments   []Treatment        `gorm:"foreignKey:SalonID"`
}

func (Salon) TableName() string {
	return "salons"
}

type SalonOpeningHour struct {
	ID        string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SalonID   string  `gorm:"column:salon_id;type:uuid"`
	DayName   string  `gorm:"column:day_name"`
	DayOrder  int     `gorm:"column:day_order"`
	OpenTime  *string `gorm:"column:open_time"`
	CloseTime *string `gorm:"column:close_time"`
	IsClosed  bool    `gorm:"column:is_closed"`
}

func (SalonOpeningHour) TableName() string {
	return "salon_opening_hours"
}

type Treatment struct {
	ID              string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SalonID         string `gorm:"column:salon_id;type:uuid"`
	Slug            string `gorm:"column:slug"`
	Name            string `gorm:"column:name"`
	DurationMinutes int    `gorm:"column:duration_minutes"`
	PriceFrom       int    `gorm:"column:price_from"`
	DisplayOrder    int    `gorm:"column:display_order"`
}

func (Treatment) TableName() string {
	return "services"
}

type SalonPageResponse struct {
	ID           string                  `json:"id"`
	Name         string                  `json:"name"`
	Address      string                  `json:"address"`
	Description  string                  `json:"description"`
	HeroImageURL string                  `json:"heroImageUrl"`
	Phone        string                  `json:"phone"`
	Email        string                  `json:"email"`
	OpeningHours []SalonOpeningHourEntry `json:"openingHours"`
	Treatments   []SalonTreatmentEntry   `json:"treatments"`
}

type SalonOpeningHourEntry struct {
	Day    string `json:"day"`
	Order  int    `json:"order"`
	Open   string `json:"open,omitempty"`
	Close  string `json:"close,omitempty"`
	Closed bool   `json:"closed"`
}

type SalonTreatmentEntry struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	DurationMinutes int    `json:"durationMinutes"`
	PriceFrom       int    `json:"priceFrom"`
}
