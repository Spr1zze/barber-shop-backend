package Type

import "github.com/google/uuid"

type SalonDetails struct {
	ID      uuid.UUID `json:"id"`
	Slug    string    `json:"slug"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
}
