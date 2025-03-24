package crypto

import (
	"time"

	"github.com/golang-jwt/jwt/v5" 
)

// 32-byte secret key for HS256
var jwtKey = []byte{
	0x3a, 0x5e, 0x1b, 0x67, 0x9c, 0x2f, 0xa0, 0x8e,
	0x1f, 0x4d, 0x6b, 0x8a, 0x5c, 0x3e, 0x7d, 0x2b,
	0x6a, 0x9f, 0x4c, 0x8d, 0x1e, 0x2c, 0x7a, 0x5b,
	0x3f, 0x6e, 0x8c, 0x9a, 0x7b, 0x2d, 0x4a, 0x1c,
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Use HS256 for symmetric signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Changed from ES256 to HS256
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}