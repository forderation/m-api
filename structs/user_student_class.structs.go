package structs

type UserStudentClass struct {
	Model
	UserStudentId uint
	ClassId       uint
	TheoryClassId uint
	TableNumber   uint
}

func (UserStudentClassWithAssociation) TableName() string {
	return "user_student_class"
}

type UserStudentClassWithAssociation struct {
	UserStudentClass
	UserStudent UserStudent          `gorm:"foreignkey:ID;references:UserStudentId"`
	Class       ClassWithAssociation `gorm:"foreignkey:ID;references:ClassId"`
}

func (UserStudentClassWithDetail) TableName() string {
	return "user_student_class"
}

type UserStudentClassWithDetail struct {
	UserStudentClass
	UserStudentDetail UserStudentWithAssociationDetail `gorm:"foreignkey:ID;references:UserStudentId"`
}
