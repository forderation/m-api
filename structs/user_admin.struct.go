package structs

type UserAdmin struct {
	Model
	UserID  uint
	Section string
	Active  bool
}

func (UserAdminWithAssociation) TableName() string { return "user_admin" }

type UserAdminWithAssociation struct {
	UserAdmin
	Classes                      []Class                       `gorm:"foreignkey:UserAdminID;references:ID"`
	PresenceAssistantInstructors []PresenceAssistantInstructor `gorm:"foreignkey:UserAdminId;references:ID"`
	LeftStuffsAsPoster           []LeftStuff                   `gorm:"foreignkey:UserAdminIdPoster;references:ID"`
	LeftStuffsAsGiver            []LeftStuff                   `gorm:"foreignkey:UserAdminIdGiver;references:ID"`
	User                         User                          `gorm:"foreignkey:ID;references:UserID"`
}
