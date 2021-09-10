package structs

type UserForgetPassword struct {
	Model
	ResetToken string
	UserID     uint
}
