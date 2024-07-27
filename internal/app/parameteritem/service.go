package parameteritem

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/model"
	"compass_mini_api/internal/repository"
	res "compass_mini_api/pkg/util/response"
	"compass_mini_api/pkg/util/trxmanager"
	"strings"

	"gorm.io/gorm"
)

type Service interface {
	GetAllDivision(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*dto.GetAllDivisionResponse, error)
}

type service struct {
	Repository repository.ParameterItem
	Db         *gorm.DB
}

func NewService(f *factory.Factory) Service {
	repository := f.ParameterItemRepository
	db := f.Db
	return &service{
		repository,
		db,
	}
}

func (s *service) GetAllDivision(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*dto.GetAllDivisionResponse, error) {
	var result dto.GetAllDivisionResponse
	var data *[]model.GetAllDivisionResponse
	var count *int

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.GetAllDivision(ctx, queryPagination, queryFilter)
		if err != nil && err.Error() != "record not found" {
			if strings.Contains(err.Error(), "400") {
				return err
			}
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		count, err = s.Repository.GetCountDivision(ctx, queryFilter)
		if err != nil {
			if strings.Contains(err.Error(), "400") {
				return err
			}
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		return nil
	}); err != nil {
		return &result, err
	}
	result = dto.GetAllDivisionResponse{
		Data:  *data,
		Count: count,
	}
	return &result, nil
}
