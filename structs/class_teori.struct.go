package structs

import "time"

type ClassTeori struct {
	ID              uint
	ClassCategoryId int32
	FullName        string
	ShortName       string
	Active          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
