package model

import (
	"compass_mini_api/internal/abstraction"

	"gorm.io/gorm"
)

type CompanyEntity struct {
	Name           *string `json:"name" form:"name"`
	InvoiceDueDate *int    `json:"invoiceduedate" form:"invoiceduedate" gorm:"column:invoiceduedate"`
	IsVendor       *bool   `json:"isvendor" form:"isvendor" gorm:"column:isvendor"`
}

type CompanyFk struct {
}

type CompanyEntityModel struct {
	Id int `json:"id" gorm:"primaryKey"`

	// relation
	CompanyFk

	// entity
	CompanyEntity

	// abstraction
	// abstraction.Entity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (CompanyEntityModel) TableName() string {
	return "masterdata.company"
}

func (m *CompanyEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	// m.CreateDate = date.DateTodayLocal()
	return
}

func (m *CompanyEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	// m.LastLogin = date.DateTodayLocal()
	return
}

type GetCountCompanyResponse struct {
	Count *int `json:"count"`
}
