package structs

import "time"

type CertificationSession struct {
	ID              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	CertificationId uint
	Name            string
	Step            int32
	Description     string
}
