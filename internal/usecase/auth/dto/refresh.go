package dto

type Refresh struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
