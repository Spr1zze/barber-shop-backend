package Type

import "github.com/google/uuid"

type HistoryView struct {
	ID           uuid.UUID `json:"id"`
	Service      string    `json:"service"`
	BarberName   string    `json:"barbername"`
	SalonAddress string    `json:"salonaddress"`
	DateTime     string    `json:"datetime"`
	Price        int       `json:"price"`
}
