package engine

import "gorm.io/gorm"

type InDB struct {
	DB *gorm.DB
}
