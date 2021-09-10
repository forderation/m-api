package structs

import "time"

type Event struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	NameEvent string
	DateEvent time.Time
}
