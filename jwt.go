package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("f4zdAsGAtmvZ518mbWg6DQZcyR4u6LEfUDMPVoWyRDlyi0DjFx92QKxcchqePmxd")

func GenerateToken(id string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "go-websocket",
		Id:        id,
	})
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
func ParseToken(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
