package structs

type UserStudent struct {
	Model
	UserId uint
	Active bool
}

func (UserStudentWithAssociation) TableName() string {
	return "user_student"
}

type UserStudentWithAssociation struct {
	UserStudent
	PresenceStudents   []PresenceStudent  `gorm:"foreignkey:UserStudentId;references:ID"`
	TaskDetails        []TaskDetail       `gorm:"foreignkey:UserStudentId;references:ID"`
	UserStudentClasses []UserStudentClass `gorm:"foreignkey:UserStudentId;references:ID"`
	User               User               `gorm:"foreignkey:ID;references:UserId"`
}

func (UserStudentWithAssociationDetail) TableName() string {
	return "user_student"
}

type UserStudentWithAssociationDetail struct {
	UserStudent
	UserDetail User `gorm:"foreignkey:ID;references:UserId"`
}
