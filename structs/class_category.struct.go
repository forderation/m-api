package structs

type ClassCategory struct {
	Model
	Parent       int
	CategoryName string
	IsCourseName bool
	ClassTeoriID uint
	IsTeoriClass bool
	Active       int
}


func (ClassCategory) TableName() string { return "class_category" }

func (ClassCategoryWithAssociation) TableName() string { return "class_category" }

type ClassCategoryWithAssociation struct {
	ClassCategory
	Classes []Class `gorm:"foreignKey:ClassCategoryID;references:ID"`
}
