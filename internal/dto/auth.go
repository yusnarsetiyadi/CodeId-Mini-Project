package dto

import (
	"compass_mini_api/internal/model"
)

// middleware
type AuthMiddlewareRequest struct {
	ID int `json:"id"`
}

type AuthMiddlewareResponse struct {
	ID int `json:"id"`
}

// Login
type AuthLoginRequest struct {
	MobilePhone string      `json:"mobilephone" form:"mobilephone" example:"+6281234567890" validate:"required"`
	Password    string      `json:"password" form:"password" example:"Test12345@" validate:"required"`
	Geolocation Geolocation `json:"geolocation" form:"geolocation" validate:"required"`
}
type Geolocation struct {
	IP          string  `json:"ip" example:"111.94.121.97"`
	Asn         string  `json:"asn" example:"AS23700 Linknet-Fastnet ASN"`
	Netmask     int64   `json:"netmask" example:"16"`
	Hostname    string  `json:"hostname" example:"fm-dyn-111-94-121-97.fast.net.id."`
	City        string  `json:"city" example:"Jakarta"`
	PostCode    string  `json:"post_code" example:"15710"`
	Country     string  `json:"country" example:"Indonesia"`
	CountryCode string  `json:"country_code" example:"ID"`
	Latitude    float64 `json:"latitude" example:"-6.1743998527526855"`
	Longitude   float64 `json:"longitude" example:"106.82939910888672"`
}
type AuthLoginResponse struct {
	model.Token
}

// splash
type AuthSplashRequest struct {
	model.Token
}
type AuthSplashResponse struct {
	model.Token
}

// logout
type AuthLogoutResponse struct {
	Message string `json:"message"`
}

// data from token
type GetDataTokenResponse struct {
	Id          int    `json:"id"`
	IdentityId  int    `json:"identityid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobilephone"`
	RoleId      int    `json:"roleid"`
	Role        string `json:"role"`
	IsActive    bool   `json:"isactive"`
	IsLocked    bool   `json:"islocked"`
}
