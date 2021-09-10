package structs

import "time"

type UserAssistantClass struct {
	Model
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	UserAssistantID uint
	ClassID         uint
}

func (UserAssistantClassWithAssociation) TableName() string {
	return "user_assistant_class"
}

type UserAssistantClassWithAssociation struct {
	UserAssistantClass
	UserAssistant UserAssistant        `gorm:"foreignkey:ID;references:UserAssistantID"`
	Class         ClassWithAssociation `gorm:"foreignkey:ID;references:ClassID"`
}
