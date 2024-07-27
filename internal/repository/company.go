package repository

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/model"
	"compass_mini_api/pkg/util/general"

	"gorm.io/gorm"
)

type Company interface {
	GetAllCompany(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*[]model.CompanyEntityModel, error)
	GetCountCompany(ctx *abstraction.Context, queryFilter *abstraction.QueryFilter) (*int, error)
}

type company struct {
	abstraction.Repository
}

func NewCompany(db *gorm.DB) *company {
	return &company{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *company) GetAllCompany(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*[]model.CompanyEntityModel, error) {
	conn := r.CheckTrx(ctx)

	limit, offset := general.ProcessQueryPagination(queryPagination)
	filter, err := general.ProcessQueryFilter(queryFilter)
	if err != nil {
		return nil, err
	}

	var data []model.CompanyEntityModel
	err = conn.Table("masterdata.company").
		Select("*").
		Where(filter).
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

func (r *company) GetCountCompany(ctx *abstraction.Context, queryFilter *abstraction.QueryFilter) (*int, error) {
	conn := r.CheckTrx(ctx)

	filter, err := general.ProcessQueryFilter(queryFilter)
	if err != nil {
		return nil, err
	}

	var data model.GetCountCompanyResponse
	err = conn.Table("masterdata.company").
		Select("COUNT(*) AS count").
		Where(filter).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data.Count, nil
}
