package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"weplant-backend/helper"
	"weplant-backend/model/web"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTService interface {
	GenerateToken(payload web.JWTPayload) string
	ValidateToken(tokenString string) (web.JWTPayload, error)
}

type jwtServiceImpl struct {
}

func NewJWTService() JWTService {
	return &jwtServiceImpl{}
}

func (service *jwtServiceImpl) GenerateToken(payload web.JWTPayload) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   payload.Id,
		"role": payload.Role,
	})
	tokenString, err := token.SignedString(secretKey)
	helper.PanicIfError(err)
	return tokenString
}

func (service *jwtServiceImpl) ValidateToken(tokenString string) (web.JWTPayload, error) {

	payload := web.JWTPayload{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return payload, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		payload.Id = claims["id"].(string)
		payload.Role = claims["role"].(string)
		return payload, nil
	}
	return payload, err
}
