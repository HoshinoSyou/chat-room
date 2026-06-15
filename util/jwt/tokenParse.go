package jwt

import (
	"chat-room/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

type Payload struct {
	//Iss      string `json:"iss"`
	//Exp      string `json:"exp"`
	//Iat      string `json:"iat"`
	Username string `json:"username"`
	Uid      uint   `json:"uid"`
	jwt.RegisteredClaims
}

func TokenParse(token string) (*Payload, error) {
	log.Println(token)
	tokenClaims := config.AppConfig.Jwt
	claims, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenClaims.SecretKey), nil
	})
	if err != nil {
		log.Printf("toekn 解析失败！错误信息：%v", err)
		return nil, err
	}
	payload, ok := claims.Claims.(*Payload)
	if ok && claims.Valid {
		return payload, nil
	}
	err = errors.New("token 解析失败！")
	log.Println(err)
	return nil, err
}
