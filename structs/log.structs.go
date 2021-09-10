package structs

type Log struct {
	Model
	UserId   uint
	Activity string
}

func (LogWithAssociation) TableName() string {
	return "log"
}

type LogWithAssociation struct {
	Log
	User User `gorm:"foreignkey:ID;references:UserId"`
}
