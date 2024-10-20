package dto

type SignUp struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,min=2,max=20"`
	Password        string `json:"password" validate:"required,min=6,max=20"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=6,max=20"`
}
