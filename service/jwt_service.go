package service

import (
	"weplant-backend/model/web"
)

type JWTService interface {
	GenerateToken(payload web.JWTPayload) string
	ValidateToken(tokenString string) (web.JWTPayload, error)
}
