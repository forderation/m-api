package structs

type LeftStuff struct {
	Model
	Kind              string
	UserAdminIdPoster uint
	Location          string
	Taker             string
	DatetimeTake      string
	UserAdminIdGiver  uint
	Note              string
}

func (LeftStuffWithAssociation) TableName() string { return "left_stuff" }

type LeftStuffWithAssociation struct {
	LeftStuff
	UserAdminPoster UserAdmin `gorm:"foreignkey:ID;references:UserAdminIdPoster"`
	UserAdminGiver  UserAdmin `gorm:"foreignkey:ID;references:UserAdminIdGiver"`
}
