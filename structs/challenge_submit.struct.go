package structs

import "time"

type ChallengeSubmit struct {
	ID          uint
	ChallengeId int32
	UserId      int32
	File        string
	Point       int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
