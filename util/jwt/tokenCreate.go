package jwt

import (
	"chat-room/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func TokenCreate(username string, uid uint) (string, error) {
	cfg := config.AppConfig.Jwt
	payload := &Payload{
		Username: username,
		Uid:      uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Expired) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now())},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(cfg.SecretKey))
	return tokenString, err
}
