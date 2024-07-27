package dto

import (
	"compass_mini_api/internal/model"
	"mime/multipart"
)

type EmployeeGetAllResponse struct {
	Data  []model.EmployeeGetAllResponse `json:"item"`
	Count *int                           `json:"count"`
}

type EmployeeSupervisorGetAllResponse struct {
	Data  []model.EmployeeSupervisorGetAllResponse `json:"item"`
	Count *int                                     `json:"count"`
}

type CreateEmployeeRequest struct {
	Name          string                `json:"name" form:"name" example:"Yusnar Setiyadi" validate:"required"`
	Email         string                `json:"email" form:"email" example:"yusnar@code.id"`
	PhoneNumber   string                `json:"phonenumber" form:"phonenumber" example:"+6281234567812" gorm:"column:phonenumber"`
	EmployeePhoto *multipart.FileHeader `json:"employeephoto" form:"employeephoto" gorm:"column:employeephoto" swaggerignore:"true"`
	CompanyId     int                   `json:"companyid" form:"companyid" example:"43" validate:"required" gorm:"column:companyid"`
	Company       string                `json:"company" form:"company" example:"CODE.ID" validate:"required"`
	DivisionId    int                   `json:"divisionid" form:"divisionid" example:"16" validate:"required" gorm:"column:divisionid"`
	Division      string                `json:"division" form:"division" example:"Maintenance" validate:"required"`
	SupervisorId  int                   `json:"supervisorid" form:"supervisorid" example:"11" validate:"required" gorm:"column:supervisorid"`
	Supervisor    string                `json:"supervisor" form:"supervisor" example:"Herru Purnomo Santoso" validate:"required"`
	JoinDate      string                `json:"joindate" form:"joindate" example:"2024-01-24" validate:"required" gorm:"column:joindate"`
}

type GetEmployeeByIdRequestParam struct {
	Id int `param:"id" validate:"required"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type UpdateEmployeeRequestParam struct {
	Id int `param:"id" validate:"required"`
}

type UpdateEmployeeRequest struct {
	Name          *string               `json:"name" form:"name" example:"Fahrul Update"`
	Email         *string               `json:"email" form:"email" example:"fahrul@code.id"`
	PhoneNumber   *string               `json:"phonenumber" form:"phonenumber" example:"+6281234567813" gorm:"column:phonenumber"`
	EmployeePhoto *multipart.FileHeader `json:"employeephoto" form:"employeephoto" gorm:"column:employeephoto" swaggerignore:"true"`
	CompanyId     *int                  `json:"companyid" form:"companyid" example:"43" gorm:"column:companyid"`
	Company       *string               `json:"company" form:"company" example:"CODE.ID"`
	DivisionId    *int                  `json:"divisionid" form:"divisionid" example:"16" gorm:"column:divisionid"`
	Division      *string               `json:"division" form:"division" example:"Maintenance"`
	IsActive      *bool                 `json:"isactive" form:"isactive" example:"false" gorm:"column:isactive"`
	SupervisorId  *int                  `json:"supervisorid" form:"supervisorid" example:"11" gorm:"column:supervisorid"`
	Supervisor    *string               `json:"supervisor" form:"supervisor" example:"Herru Purnomo Santoso"`
	JoinDate      *string               `json:"joindate" form:"joindate" example:"2024-01-25" gorm:"column:joindate"`
	ResignDate    *string               `json:"resigndate" form:"resigndate" example:"2024-01-25" gorm:"column:resigndate"`
}

type CreateEmployeeRequestWithBase64 struct {
	Name          string        `json:"name" form:"name" example:"Yusnar Setiyadi" validate:"required"`
	Email         string        `json:"email" form:"email" example:"yusnar@code.id"`
	PhoneNumber   string        `json:"phonenumber" form:"phonenumber" example:"+6281234567812" gorm:"column:phonenumber"`
	EmployeePhoto EmployeePhoto `json:"employeephoto" form:"employeephoto" validate:"required"`
	CompanyId     int           `json:"companyid" form:"companyid" example:"43" validate:"required" gorm:"column:companyid"`
	Company       string        `json:"company" form:"company" example:"CODE.ID" validate:"required"`
	DivisionId    int           `json:"divisionid" form:"divisionid" example:"16" validate:"required" gorm:"column:divisionid"`
	Division      string        `json:"division" form:"division" example:"Maintenance" validate:"required"`
	SupervisorId  int           `json:"supervisorid" form:"supervisorid" example:"11" validate:"required" gorm:"column:supervisorid"`
	Supervisor    string        `json:"supervisor" form:"supervisor" example:"Herru Purnomo Santoso" validate:"required"`
	JoinDate      string        `json:"joindate" form:"joindate" example:"2024-01-24" validate:"required" gorm:"column:joindate"`
}

type UpdateEmployeeRequestWithBase64 struct {
	Name          *string        `json:"name" form:"name" example:"Fahrul Update"`
	Email         *string        `json:"email" form:"email" example:"fahrul@code.id"`
	PhoneNumber   *string        `json:"phonenumber" form:"phonenumber" example:"+6281234567813" gorm:"column:phonenumber"`
	EmployeePhoto *EmployeePhoto `json:"employeephoto" form:"employeephoto"`
	CompanyId     *int           `json:"companyid" form:"companyid" example:"43" gorm:"column:companyid"`
	Company       *string        `json:"company" form:"company" example:"CODE.ID"`
	DivisionId    *int           `json:"divisionid" form:"divisionid" example:"16" gorm:"column:divisionid"`
	Division      *string        `json:"division" form:"division" example:"Maintenance"`
	IsActive      *bool          `json:"isactive" form:"isactive" example:"false" gorm:"column:isactive"`
	SupervisorId  *int           `json:"supervisorid" form:"supervisorid" example:"11" gorm:"column:supervisorid"`
	Supervisor    *string        `json:"supervisor" form:"supervisor" example:"Herru Purnomo Santoso"`
	JoinDate      *string        `json:"joindate" form:"joindate" example:"2024-01-25" gorm:"column:joindate"`
	ResignDate    *string        `json:"resigndate" form:"resigndate" example:"2024-01-25" gorm:"column:resigndate"`
}

type EmployeePhoto struct {
	Name string `json:"name" example:"gambar.png" gorm:"column:employeephoto"`
	Type string `json:"type" example:"image/png"`
	Size int    `json:"size" example:"98220"`
	Data string `json:"data" example:"base64 value"`
}

type GetEmployeePhotoRequestParam struct {
	Base64 string `param:"base64" validate:"required"`
}
