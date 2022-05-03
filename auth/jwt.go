package auth

import (
	"errors"
	"time"

	"github.com/Devansh3712/go-banking-api/config"
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var (
	tokenExpireDuration = time.Hour
	issuer              = config.EnvValue("ISSUER")
	secretKey           = []byte(config.EnvValue("SECRET_KEY"))
)

func GenerateToken(email string) (string, error) {
	claims := JWTClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
