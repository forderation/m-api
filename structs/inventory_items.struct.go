package structs

import "time"

type InventoryItems struct {
	ID            uint
	Parent        int
	CategoryItem  int
	Name          string
	SerialNumber  string
	SnProduct     string
	Description   string
	EntryDate     time.Time
	CanBeRent     int
	Availability  bool
	InMaintenance bool
	Status        string
	Location      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	UserId        int
}

func (InventoryItemsWithAssociation) TableName() string {
	return "inventory_items"
}

type InventoryItemsWithAssociation struct {
	InventoryItems
	ItemMaintenances []ItemMaintenance `gorm:"foreignkey:ItemId;references:ID"`
	ItemCategory     ItemCategory      `gorm:"foreignkey:ID;references:CategoryItem"`
}
