package structs

import "time"

type ItemCategory struct {
	ID        uint
	Code      int
	Parent    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserId    int
}

func (ItemCategoryWithAssociation) TableName() string {
	return "item_category"
}

type ItemCategoryWithAssociation struct {
	ItemCategory
	InventoryItems []InventoryItems `gorm:"foreignkey:CategoryItem;references:ID"`
}
