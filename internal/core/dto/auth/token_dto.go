package auth

type Token struct {
	UserID           int64  `json:"userId"`
	TokenType        string `json:"tokenType"`
	AccessToken      string `json:"accessToken"`
	AccessExpiresIn  int64  `json:"accessExpiresIn"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresIn int64  `json:"refreshExpiresIn"`
}
