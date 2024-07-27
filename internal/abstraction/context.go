package abstraction

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Context struct {
	echo.Context
	Auth *AuthContext
	Trx  *TrxContext
}

type AuthContext struct {
	ID          int    `json:"id"`
	IdentityId  int    `json:"identityid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobilephone"`
	RoleId      int    `json:"roleid"`
	Role        string `json:"role"`
	IsActive    bool   `json:"isactive"`
	IsLocked    bool   `json:"islocked"`
	Exp         int64  `json:"exp"`
	jwt.RegisteredClaims
}

type TrxContext struct {
	Db *gorm.DB
}
