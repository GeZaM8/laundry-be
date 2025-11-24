package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("SECRETKEYSUPERAMAN")

func GenerateToken(userId uint) (string, error) {
    claims := jwt.MapClaims{
        "sub": userId,
        "exp": time.Now().Add(time.Hour * 24).Unix(), // token valid 24 jam
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JWT_SECRET)
}
