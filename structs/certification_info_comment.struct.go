package structs

import "time"

type CertificationInfoComment struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IdInfo    uint
	UserId    uint
	Parent    int32
	Comment   string
}
