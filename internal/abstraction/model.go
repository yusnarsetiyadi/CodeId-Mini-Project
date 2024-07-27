package abstraction

import (
	"compass_mini_api/pkg/util/date"
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	CreateDate *time.Time `json:"createdate" gorm:"column:createdate"`
	LastLogin  *time.Time `json:"lastlogin" gorm:"column:lastlogin"`
}

type Entity2 struct {
	CreatedDate time.Time  `json:"created_date"`
	CreatedBy   *string    `json:"created_by"`
	UpdatedDate *time.Time `json:"updated_date"`
	UpdatedBy   *string    `json:"updated_by"`
}

type Entity3 struct {
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   *string   `json:"created_by"`
}

type Entity4 struct {
	CreatedDate string  `json:"created_date"`
	CreatedBy   *string `json:"created_by"`
	UpdatedDate *string `json:"updated_date"`
	UpdatedBy   *string `json:"updated_by"`
}

type RequestEntity struct {
	Page      *string `query:"page"`
	PageSize  *string `query:"pageSize"`
	OrderName *string `query:"orderName"`
	OrderType *string `query:"orderType"`
}

type QueryEntity struct {
	Entity string `query:"entity" validate:"required"`
}

type QueryPagination struct {
	Limit  string `query:"limit" validate:"required"`
	Offset string `query:"offset" validate:"required"`
}

type QueryOrder struct {
	Order     string `query:"order"`
	Direction string `query:"direction"`
}

type Condition struct {
	Column      string `json:"column"`
	Value       string `json:"value"`
	Comparation string `json:"comparation"`
}

type QueryFilter struct {
	Conditions string `query:"conditions"`
}

func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	m.LastLogin = date.DateTodayLocal()
	return
}

func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateDate = date.DateTodayLocal()
	return
}
