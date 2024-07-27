package model

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/pkg/util/date"

	"gorm.io/gorm"
)

type UserEntity struct {
	Name                 *string  `json:"name" form:"name" gorm:"column:name"`
	Email                *string  `json:"email" form:"email" gorm:"column:email"`
	MobilePhone          *string  `json:"mobilephone" form:"mobilephone" gorm:"column:mobilephone"`
	Role                 *string  `json:"role" form:"role" gorm:"column:role"`
	Password             string   `json:"password" form:"password" gorm:"column:password"`
	Salt                 *string  `json:"salt" form:"salt"`
	DeviceInfo           *string  `json:"deviceinfo" form:"deviceinfo" gorm:"column:deviceinfo"`
	DeviceLocation       *string  `json:"devicelocation" form:"devicelocation" gorm:"column:devicelocation"`
	IpAddress            *string  `json:"ipaddress" form:"ipaddress" gorm:"column:ipaddress"`
	Latitude             *float64 `json:"latitude" form:"latitude"`
	Longitude            *float64 `json:"longitude" form:"longitude"`
	IsActive             *bool    `json:"isactive" form:"isactive" gorm:"column:isactive"`
	IsLocked             *bool    `json:"islocked" form:"islocked" gorm:"column:islocked"`
	AccessToken          *string  `json:"accesstoken" form:"accesstoken" gorm:"column:accesstoken"`
	RefreshToken         *string  `json:"refreshtoken" form:"refreshtoken" gorm:"column:refreshtoken"`
	ResetToken           *string  `json:"resettoken" form:"resettoken" gorm:"column:resettoken"`
	DeviceToken          *string  `json:"devicetoken" form:"devicetoken" gorm:"column:devicetoken"`
	Otp                  *string  `json:"otp" form:"otp"`
	LoginAttempt         *int8    `json:"loginattempt" form:"loginattempt" gorm:"column:loginattempt"`
	ResetPasswordAttempt *int8    `json:"resetpasswordattempt" form:"resetpasswordattempt" gorm:"column:resetpasswordattempt"`
}

type UserFk struct {
	IdentityId *int `json:"identityid" form:"identityid" gorm:"column:identityid"`
	RoleId     *int `json:"roleid" form:"roleid" gorm:"column:roleid"`
}

type UserCreateUserRequest struct {
	IdentityId  *int    `json:"identityid" form:"identityid" validate:"required" gorm:"column:identityid"`
	Name        *string `json:"name" form:"name" validate:"required"`
	Email       *string `json:"email" form:"email" validate:"required"`
	MobilePhone *string `json:"mobilephone" form:"mobilephone" validate:"required" gorm:"column:mobilephone"`
	RoleId      *int    `json:"roleid" form:"roleid" validate:"required" gorm:"column:roleid"`
	Role        *string `json:"role" form:"role" validate:"required"`
	Password    *string `json:"password" form:"password" validate:"required"`
}

type UserEntityModel struct {
	Id int `json:"id" gorm:"primaryKey"`

	// relation
	UserFk

	// entity
	UserEntity

	// abstraction
	abstraction.Entity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (UserEntityModel) TableName() string {
	return "appsetting.user"
}

func (m *UserEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateDate = date.DateTodayLocal()
	return
}

func (m *UserEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	// m.LastLogin = date.DateTodayLocal()
	return
}

type UserLoginModel struct {
	Id           int    `json:"id"`
	IdentityId   int    `json:"identityid" gorm:"column:identityid"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobilePhone  string `json:"mobilephone" gorm:"column:mobilephone"`
	RoleId       int    `json:"roleid" gorm:"column:roleid"`
	Role         string `json:"role"`
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	IsActive     bool   `json:"isactive" gorm:"column:isactive"`
	IsLocked     bool   `json:"islocked" gorm:"column:islocked"`
	AccessToken  string `json:"accesstoken" gorm:"column:accesstoken"`
	RefreshToken string `json:"refreshtoken" gorm:"column:refreshtoken"`
	ResetToken   string `json:"resettoken" gorm:"column:resettoken"`
	LoginAttempt int8   `json:"loginattempt" gorm:"column:loginattempt"`
}
