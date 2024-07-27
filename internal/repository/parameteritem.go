package repository

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/model"
	"compass_mini_api/pkg/util/general"

	"gorm.io/gorm"
)

type ParameterItem interface {
	GetAllDivision(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*[]model.GetAllDivisionResponse, error)
	GetCountDivision(ctx *abstraction.Context, queryFilter *abstraction.QueryFilter) (*int, error)
}

type parameteritem struct {
	abstraction.Repository
}

func NewParameterItem(db *gorm.DB) *parameteritem {
	return &parameteritem{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *parameteritem) GetAllDivision(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*[]model.GetAllDivisionResponse, error) {
	conn := r.CheckTrx(ctx)
	limit, offset := general.ProcessQueryPagination(queryPagination)
	filter, err := general.ProcessQueryFilter(queryFilter)
	if err != nil {
		return nil, err
	}

	var data []model.GetAllDivisionResponse
	err = conn.Table("appsetting.parameteritem").
		Select("id,name,parameterid,parameter,isactive").
		Where("parameter = ? AND isactive = ?", "Division", true).Where(filter).
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

func (r *parameteritem) GetCountDivision(ctx *abstraction.Context, queryFilter *abstraction.QueryFilter) (*int, error) {
	conn := r.CheckTrx(ctx)

	filter, err := general.ProcessQueryFilter(queryFilter)
	if err != nil {
		return nil, err
	}

	var data model.GetCountDivisionResponse
	err = conn.Table("appsetting.parameteritem").
		Select("COUNT(*) AS count").
		Where("parameter = ? AND isactive = ? ", "Division", true).Where(filter).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data.Count, nil
}
