package dto

type TokenDto struct {
	UserID        int    `json:"userId"`
	TokenType     string `json:"tokenType"`
	AccessToken   string `json:"accessToken"`
	AccessExpire  int64  `json:"accessExpire"`
	RefreshToken  string `json:"refreshToken"`
	RefreshExpire int64  `json:"refreshExpire"`
}
