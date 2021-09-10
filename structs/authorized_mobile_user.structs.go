package structs

import "time"

func (AuthorizedMobileUser) TableName() string {
	return "authorized_mobile_users"
}

type AuthorizedMobileUser struct {
	HalfModel
	DeviceToken  string
	UserId       uint
	DeviceName   string
	LastLoggedIn time.Time
	State        bool
}

func (AuthorizedMobileUserWithAssociation) TableName() string {
	return "authorized_mobile_users"
}

type AuthorizedMobileUserWithAssociation struct {
	AuthorizedMobileUser
	User User `gorm:"foreignkey:ID;references:UserId"`
}
