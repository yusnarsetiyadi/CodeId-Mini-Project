package model

import (
	"compass_mini_api/internal/config"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	AccessToken  string `json:"access_token" example:"your access token"`
	RefreshToken string `json:"refresh_token" example:"your refresh token"`
}

type AccessTokenClaims struct {
	Id          string `json:"id"`
	IdentityId  string `json:"identityid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobilephone"`
	RoleId      string `json:"roleid"`
	Role        string `json:"role"`
	IsActive    string `json:"isactive"`
	IsLocked    string `json:"islocked"`
	Exp         int64  `json:"exp"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	Exp int64 `json:"exp"`
	jwt.RegisteredClaims
}

func AccessToken(claims AccessTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(config.Get().Key.JwtKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func RefreshToken(refresh RefreshTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh)
	signedString, err := token.SignedString([]byte(config.Get().Key.JwtKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}
