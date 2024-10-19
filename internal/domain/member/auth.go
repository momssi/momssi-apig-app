package member

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JWTKey = []byte("your_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, expiryTime time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiryTime)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}
