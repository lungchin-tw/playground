package core

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id string, secret string, ttl time.Duration) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Id:        id,
		},
	).SignedString([]byte(secret))
}

func Verify(accessToken string, secret string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok == false {
				return nil, fmt.Errorf("Unexpected Token Signing Method")
			}

			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Invalid Token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if ok == false {
		return nil, fmt.Errorf("Invalid Token Claims")
	}

	return claims, nil
}
