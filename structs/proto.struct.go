package structs

import (
	"gorm.io/gorm"
	"time"
)

type SecurityBody struct {
	IAT int64
	ISS string
	UID uint
	RO  string
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
}
type HalfModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
