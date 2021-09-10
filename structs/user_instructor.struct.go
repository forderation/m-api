package structs

type UserInstructor struct {
	Model
	UserID uint
	Nip    string
	Active bool
}

func (UserInstructorWithAssociation) TableName() string {
	return "user_instructor"
}

type UserInstructorWithAssociation struct {
	UserInstructor
	User    User    `gorm:"foreignkey:ID;references:UserID"`
	Classes []Class `gorm:"foreignkey:UserInstructorID;references:ID"`
}
