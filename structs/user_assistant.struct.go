package structs

import "time"

type UserAssistant struct {
	Model
	UserID    uint
	Xp        uint
	JointDate time.Time
	Status    string
	Active    bool
}

func (UserAssistantWithAssociation) TableName() string {
	return "user_assistant"
}

type UserAssistantWithAssociation struct {
	UserAssistant
	UserAssistantClasses []UserAssistantClass `gorm:"foreignkey:UserAssistantID;references:ID"`
	PresenceStudents     []PresenceStudent    `gorm:"foreignkey:UserAssistantId;references:ID"`
	TaskDetails          []TaskDetail         `gorm:"foreignkey:UserAssistantId;references:ID"`
	Tasks                []Task               `gorm:"foreignkey:UserAssistantID;references:ID"`
	User                 User                 `gorm:"foreignkey:ID;references:UserID"`
	Class                Class                `gorm:"foreignkey:UserAssistantID;references:ID"`
}

func (UserAssistantWithAssociationDetail) TableName() string {
	return "user_assistant"
}

type UserAssistantWithAssociationDetail struct {
	UserAssistant
	Details SimpleUser `gorm:"foreignkey:ID;references:UserID"`
}
