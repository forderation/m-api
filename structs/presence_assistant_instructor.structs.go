package structs

type PresenceAssistantInstructor struct {
	Model
	PresenceId  uint
	UserId      uint
	UserAdminId uint
	Type        int
	Late        string
}

func (PresenceAssistantInstructorWithAssociation) TableName() string {
	return "presence_assistant_instructor"
}

type PresenceAssistantInstructorWithAssociation struct {
	PresenceAssistantInstructor
	Presence  Presence  `gorm:"foreignkey:ID;references:PresenceId"`
	UserAdmin UserAdmin `gorm:"foreignkey:ID;references:UserAdminId"`
	User      User      `gorm:"foreignkey:ID;references:UserId"`
}
