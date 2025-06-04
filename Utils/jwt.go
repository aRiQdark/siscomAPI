package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	MemberID string `json:"member_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(userID, username string) (string, error) {
	expirationTime := time.Now().Add(360 * time.Hour)
	claims := &Claims{
		MemberID: userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Invalid token signature")
		}
		return nil, errors.New("Invalid token")
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("Token expired")
	}

	return claims, nil
}
