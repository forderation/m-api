package structs

import "time"

type User struct {
	Model
	UserName      string
	Password      string
	Email         string
	FullName      string
	BirthPlace    string
	BirthDate     time.Time
	Gender        bool
	Address       string
	AddressNow    string
	LastLogin     time.Time
	Phone         string
	Fb            string
	Skype         string
	Whatsapp      string
	Line          string
	FingerPrintId uint
	AuthToken     string
	Border        string
}

type SimpleUser struct {
	Model
	UserName  string
	Email     string
	FullName  string
	Gender    bool
	LastLogin time.Time
	Border    string
}

func (UserWithAssociation) TableName() string {
	return "user"
}

func (SimpleUser) TableName() string {
	return "user"
}

type UserWithAssociation struct {
	User
	BlackLists []BlackList `gorm:"foreignkey:UserId;references:ID"`
	Infos      []Info      `gorm:"foreignkey:UserID;references:ID"`
	//Logs                         []Log                         `gorm:"foreignkey:UserId;references:ID"`
	Suggests              []Suggest              `gorm:"foreignkey:UserId;references:ID"`
	AuthorizedMobileUsers []AuthorizedMobileUser `gorm:"foreignkey:UserId;references:ID"`
	UserInstructor        []UserInstructor       `gorm:"foreignkey:UserID;references:ID"`
	UserStudents          []UserStudent          `gorm:"foreignkey:UserId;references:ID"`
	UserAssistants        []UserAssistant        `gorm:"foreignkey:UserID;references:ID"`
	UserAdmins            []UserAdmin            `gorm:"foreignkey:UserID;references:ID"`
	//PresenceAssistantInstructors []PresenceAssistantInstructor `gorm:"foreignkey:UserId;references:ID"`
}
