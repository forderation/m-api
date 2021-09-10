package structs

import "time"

type ItemMaintenance struct {
	ID              uint
	ItemId          int
	MaintenanceBy   string
	MaintenanceType int
	StartDate       time.Time
	FinishDate      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	UserId          int32
}

func (ItemMaintenanceWithAssociation) TableName() string {
	return "item_maintenance"
}

type ItemMaintenanceWithAssociation struct {
	ItemMaintenance
	Item InventoryItems `gorm:"foreignkey:ID;references:ItemId"`
}
