package dto

type Token struct {
	AccessToken   string `json:"accessToken"`
	AccessExpire  int64  `json:"accessExpire"`
	RefreshToken  string `json:"refreshToken"`
	RefreshExpire int64  `json:"refreshExpire"`
}
