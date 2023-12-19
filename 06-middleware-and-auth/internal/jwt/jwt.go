package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTClaim(userid int64) JWTClaim {
	return JWTClaim{
		UserID: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
}

func GenerateToken(claim *JWTClaim, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secret))
}

func DecodeToken(inputToken string, claim *JWTClaim, secret string) error {
	token, err := jwt.ParseWithClaims(inputToken, claim, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	if !token.Valid {
		return err
	}
	return nil
}
