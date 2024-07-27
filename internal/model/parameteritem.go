package model

import (
	"compass_mini_api/internal/abstraction"

	"gorm.io/gorm"
)

type ParameterItemEntity struct {
	Name      *string `json:"name" form:"name"`
	Value     *string `json:"value" form:"value"`
	Parameter *string `json:"parameter" form:"parameter"`
	IsActive  *bool   `json:"isactive" form:"isactive" gorm:"column:isactive"`
}

type ParameterItemFk struct {
	ParameterId *int `json:"parameterid" form:"parameterid" gorm:"column:parameterid"`
}

type ParameterItemEntityModel struct {
	Id int `json:"id" gorm:"primaryKey"`

	// relation
	ParameterItemFk

	// entity
	ParameterItemEntity

	// abstraction
	// abstraction.Entity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (ParameterItemEntityModel) TableName() string {
	return "appsetting.parameteritem"
}

func (m *ParameterItemEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	// m.CreateDate = date.DateTodayLocal()
	return
}

func (m *ParameterItemEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	// m.LastLogin = date.DateTodayLocal()
	return
}

type GetAllDivisionResponse struct {
	Id          *int    `json:"id"`
	Name        *string `json:"name"`
	ParameterId *int    `json:"parameterid" gorm:"column:parameterid"`
	Parameter   *string `json:"parameter"`
	IsActive    *bool   `json:"isactive" gorm:"column:isactive"`
}

type GetCountDivisionResponse struct {
	Count *int `json:"count"`
}
