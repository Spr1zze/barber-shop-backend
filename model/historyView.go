package Type

import (
	"time"

	"github.com/google/uuid"
)

type HistoryView struct {
	ID           uuid.UUID
	Service      string
	BarberName   string
	SalonAddress string
	DateTime     time.Time
	Price        int
}
