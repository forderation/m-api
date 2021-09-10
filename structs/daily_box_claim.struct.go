package structs

import "time"

type DailyBoxClaim struct {
	ID        uint
	UserId    int32
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
