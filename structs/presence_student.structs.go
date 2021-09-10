package structs

type PresenceStudent struct {
	Model
	PresenceId      uint
	UserStudentId   uint
	UserAssistantId uint
	Type            int
	Late            string
}

func (PresenceStudentWithAssociation) TableName() string {
	return "presence_student"
}

type PresenceStudentWithAssociation struct {
	PresenceStudent
	UserStudent   UserStudentWithAssociationDetail `gorm:"foreignkey:ID;references:UserStudentId"`
	UserAssistant UserAssistant                    `gorm:"foreignkey:ID;references:UserAssistantId"`
	Presence      Presence                         `gorm:"foreignkey:ID;references:PresenceId"`
}
