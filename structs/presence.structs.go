package structs

type Presence struct {
	Model
	ClassId         uint
	UserAssistantID uint
	Activity        string
	TaskId          uint
}

func (PresenceWithAssociation) TableName() string { return "presence" }

type PresenceWithAssociation struct {
	Presence
	Class                        Class                            `gorm:"foreignkey:ID;references:ClassId"`
	PresenceAssistantInstructors []PresenceAssistantInstructor    `gorm:"foreignkey:PresenceId;references:ID"`
	PresenceStudents             []PresenceStudentWithAssociation `gorm:"foreignkey:PresenceId;references:ID"`
}
