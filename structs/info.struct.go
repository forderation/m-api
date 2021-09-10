package structs

type Info struct {
	Model
	Title   string
	Content string
	UserID  uint
}

func (InfoWithAssociation) TableName() string {
	return "info"
}

type InfoWithAssociation struct {
	Info
	User User `gorm:"foreignkey:ID;references:UserID"`
}
