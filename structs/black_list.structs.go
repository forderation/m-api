package structs

type BlackList struct {
	Model
	UserId              uint
	BlackListCategoryId uint
	UserIdReporter      uint
}

func (BlackListWithAssociation) TableName() string {
	return "black_list"
}

type BlackListWithAssociation struct {
	BlackList
	User User `gorm:"foreignkey:ID;references:UserId"`
}
