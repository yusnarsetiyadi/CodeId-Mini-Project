package repository

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/model"
	"compass_mini_api/pkg/util/general"

	"gorm.io/gorm"
)

type Employee interface {
	GetCountEmployee(ctx *abstraction.Context, queryFilter *abstraction.QueryFilter) (*int, error)
	GetAllEmployee(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryOrder *abstraction.QueryOrder, queryFilter *abstraction.QueryFilter) (*[]model.EmployeeGetAllResponse, error)
	GetCountEmployeeSupervisor(ctx *abstraction.Context, queryFilter *abstraction.QueryFilter) (*int, error)
	GetAllEmployeeSupervisor(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*[]model.EmployeeSupervisorGetAllResponse, error)
	CreateEmployee(ctx *abstraction.Context, data *model.EmployeeEntityModel) (*model.EmployeeEntityModel, error)
	FindByIdEmployee(ctx *abstraction.Context, id *int) (*model.EmployeeGetByIdResponse, error)
	UpdateEmployee(ctx *abstraction.Context, id *int, data *model.EmployeeEntityModel) (*model.EmployeeEntityModel, error)
}

type employee struct {
	abstraction.Repository
}

func NewEmployee(db *gorm.DB) *employee {
	return &employee{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *employee) GetCountEmployee(ctx *abstraction.Context, queryFilter *abstraction.QueryFilter) (*int, error) {
	conn := r.CheckTrx(ctx)
	filter, err := general.ProcessQueryFilter(queryFilter)
	if err != nil {
		return nil, err
	}

	var data model.CountDataEmployee
	err = conn.Table("masterdata.employee").
		Select("count(id) as count").
		Where(filter).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return &data.Count, nil
}

func (r *employee) GetAllEmployee(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryOrder *abstraction.QueryOrder, queryFilter *abstraction.QueryFilter) (*[]model.EmployeeGetAllResponse, error) {
	conn := r.CheckTrx(ctx)
	limit, offset := general.ProcessQueryPagination(queryPagination)
	order := general.ProcessQueryOrder(queryOrder)
	filter, err := general.ProcessQueryFilter(queryFilter)
	if err != nil {
		return nil, err
	}
	var data []model.EmployeeGetAllResponse

	err = conn.Table("masterdata.employee").
		Select("id, name, email, phonenumber, employeephoto, companyid, company, divisionid, division, supervisorid, supervisor, isactive, joindate, resigndate").
		Where(filter).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *employee) GetCountEmployeeSupervisor(ctx *abstraction.Context, queryFilter *abstraction.QueryFilter) (*int, error) {
	conn := r.CheckTrx(ctx)

	filter, err := general.ProcessQueryFilter(queryFilter)
	if err != nil {
		return nil, err
	}

	whereParams := map[string]interface{}{
		"isactive": true,
	}

	where := "isactive = @isactive"

	var data model.CountDataEmployee
	err = conn.Table("masterdata.employee").
		Select("count(id) as count").
		Where(where, whereParams).Where(filter).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return &data.Count, nil
}

func (r *employee) GetAllEmployeeSupervisor(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*[]model.EmployeeSupervisorGetAllResponse, error) {
	conn := r.CheckTrx(ctx)

	limit, offset := general.ProcessQueryPagination(queryPagination)
	filter, err := general.ProcessQueryFilter(queryFilter)
	if err != nil {
		return nil, err
	}

	whereParams := map[string]interface{}{
		"isactive": true,
	}

	where := "isactive = @isactive"
	var data []model.EmployeeSupervisorGetAllResponse

	err = conn.Table("masterdata.employee").
		Select("id, name, email, phonenumber, employeephoto, companyid, company, divisionid, division, supervisorid, supervisor, isactive, joindate, resigndate").
		Where(where, whereParams).Where(filter).
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *employee) CreateEmployee(ctx *abstraction.Context, data *model.EmployeeEntityModel) (*model.EmployeeEntityModel, error) {
	conn := r.CheckTrx(ctx)
	data.Context = ctx
	err := conn.Table("masterdata.employee").Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *employee) FindByIdEmployee(ctx *abstraction.Context, id *int) (*model.EmployeeGetByIdResponse, error) {
	conn := r.CheckTrx(ctx)
	var data model.EmployeeGetByIdResponse
	err := conn.Table("masterdata.employee").
		Select("name, email, phonenumber, employeephoto, company, division, supervisor, isactive, joindate, resigndate, uniquekey").
		Where("id = ?", id).
		First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *employee) UpdateEmployee(ctx *abstraction.Context, id *int, data *model.EmployeeEntityModel) (*model.EmployeeEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Table("masterdata.employee").Where("id = ?", id).Updates(data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
