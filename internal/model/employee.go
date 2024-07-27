package model

import (
	"compass_mini_api/internal/abstraction"
	"time"
)

type EmployeeEntity struct {
	Name          *string    `json:"name" form:"name" gorm:"column:name"`
	Email         *string    `json:"email" form:"email" gorm:"column:email"`
	PhoneNumber   *string    `json:"phonenumber" form:"phonenumber" gorm:"column:phonenumber"`
	EmployeePhoto *string    `json:"employeephoto" form:"employeephoto" gorm:"column:employeephoto"`
	Company       *string    `json:"company" form:"company" gorm:"column:company"`
	Division      *string    `json:"division" form:"division" gorm:"column:division"`
	IsActive      *bool      `json:"isactive" form:"isactive" gorm:"column:isactive"`
	Supervisor    *string    `json:"supervisor" form:"supervisor" gorm:"column:supervisor"`
	JoinDate      *time.Time `json:"joindate" form:"joindate" gorm:"column:joindate"`
	ResignDate    *time.Time `json:"resigndate" form:"resigndate" gorm:"column:resigndate"`
	UniqueKey     *string    `json:"uniquekey" form:"uniquekey" gorm:"column:uniquekey"`
}

type EmployeeFk struct {
	CompanyId    *int `json:"companyid" form:"companyid" gorm:"column:companyid"`
	DivisionId   *int `json:"divisionid" form:"divisionid" gorm:"column:divisionid"`
	SupervisorId *int `json:"supervisorid" form:"supervisorid" gorm:"column:supervisorid"`
}

type EmployeeEntityModel struct {
	Id int `json:"id" gorm:"primaryKey,autoIncrement"`

	// relation
	EmployeeFk

	// entity
	EmployeeEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type EmployeeGetAllResponse struct {
	Id            int     `json:"id"`
	Name          *string `json:"name"`
	Email         *string `json:"email"`
	PhoneNumber   *string `json:"phonenumber" gorm:"column:phonenumber"`
	EmployeePhoto *string `json:"employeephoto" gorm:"column:employeephoto"`
	CompanyId     *int    `json:"companyid" gorm:"column:companyid"`
	Company       *string `json:"company"`
	DivisionId    *int    `json:"divisionid" gorm:"column:divisionid"`
	Division      *string `json:"division"`
	SupervisorId  *int    `json:"supervisorid" gorm:"column:supervisorid"`
	Supervisor    *string `json:"supervisor"`
	IsActive      *bool   `json:"isactive" gorm:"column:isactive"`
	JoinDate      *string `json:"joindate" gorm:"column:joindate"`
	ResignDate    *string `json:"resigndate" gorm:"column:resigndate"`
	UniqueKey     *string `json:"uniquekey" gorm:"column:uniquekey"`
}

type EmployeeSupervisorGetAllResponse struct {
	Id            int     `json:"id"`
	Name          *string `json:"name"`
	Email         *string `json:"email"`
	PhoneNumber   *string `json:"phonenumber" gorm:"column:phonenumber"`
	EmployeePhoto *string `json:"employeephoto" gorm:"column:employeephoto"`
	CompanyId     *int    `json:"companyid" gorm:"column:companyid"`
	Company       *string `json:"company"`
	DivisionId    *int    `json:"divisionid" gorm:"column:divisionid"`
	Division      *string `json:"division"`
	SupervisorId  *int    `json:"supervisorid" gorm:"column:supervisorid"`
	Supervisor    *string `json:"supervisor"`
	IsActive      *bool   `json:"isactive" gorm:"column:isactive"`
	JoinDate      *string `json:"joindate" gorm:"column:joindate"`
	ResignDate    *string `json:"resigndate" gorm:"column:resigndate"`
	UniqueKey     *string `json:"uniquekey" gorm:"column:uniquekey"`
}

type CountDataEmployee struct {
	Count int `json:"count"`
}

type EmployeeGetByIdResponse struct {
	Name          *string `json:"name"`
	Email         *string `json:"email"`
	PhoneNumber   *string `json:"phonenumber" gorm:"column:phonenumber"`
	EmployeePhoto *string `json:"employeephoto" gorm:"column:employeephoto"`
	Company       *string `json:"company"`
	Division      *string `json:"division"`
	Supervisor    *string `json:"supervisor"`
	IsActive      *bool   `json:"isactive" gorm:"column:isactive"`
	JoinDate      *string `json:"joindate" gorm:"column:joindate"`
	ResignDate    *string `json:"resigndate" gorm:"column:resigndate"`
	UniqueKey     *string `json:"uniquekey" gorm:"column:uniquekey"`
}
