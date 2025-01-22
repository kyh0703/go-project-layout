package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId int64, email string, expire time.Time) (string, error) {
	mySigningKey := []byte("AllYourBase")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     expire,
	})

	token, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, err
}
