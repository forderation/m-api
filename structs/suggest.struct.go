package structs

import "time"

type Suggest struct {
	Model
	UserId         uint
	Suggest        string
	ReplyCreatedAt time.Time
	ReplyMessage   string
}

func (SuggestWithAssociation) TableName() string {
	return "suggest"
}

type SuggestWithAssociation struct {
	Suggest
	User User `gorm:"foreignkey:ID;references:UserId"`
}
