package feature

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/model"
	"compass_mini_api/internal/repository"
	res "compass_mini_api/pkg/util/response"
	"compass_mini_api/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	GetFeatureList(ctx *abstraction.Context, queryEntity *abstraction.QueryEntity) (*dto.GetFeatureListResponse, error)
	GetFeatureSub(ctx *abstraction.Context, payload *dto.GetFeatureSubRequestParam, queryEntity *abstraction.QueryEntity) (*dto.GetFeatureSubResponse, error)
}

type service struct {
	Repository repository.Feature
	Db         *gorm.DB
}

func NewService(f *factory.Factory) Service {
	repository := f.FeatureRepository
	db := f.Db
	return &service{
		repository,
		db,
	}
}

func (s *service) GetFeatureList(ctx *abstraction.Context, queryEntity *abstraction.QueryEntity) (*dto.GetFeatureListResponse, error) {
	var result dto.GetFeatureListResponse
	var data *[]model.FeatureListResponse
	var count *int

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.GetAllFeatureListWithAuthorization(ctx, ctx.Auth.MobilePhone, queryEntity.Entity)
		if err != nil && err.Error() != "record not found" {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		count, err = s.Repository.GetCountFeatureListWithAuthorization(ctx, ctx.Auth.MobilePhone, queryEntity.Entity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		return nil
	}); err != nil {
		return &result, err
	}
	result = dto.GetFeatureListResponse{
		Data:  *data,
		Count: count,
	}
	return &result, nil
}

func (s *service) GetFeatureSub(ctx *abstraction.Context, payload *dto.GetFeatureSubRequestParam, queryEntity *abstraction.QueryEntity) (*dto.GetFeatureSubResponse, error) {
	var result dto.GetFeatureSubResponse
	var data *[]model.FeatureSubResponse
	var count *int

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.GetAllFeatureSubWithAuthorization(ctx, ctx.Auth.MobilePhone, queryEntity.Entity, payload.Id)
		if err != nil && err.Error() != "record not found" {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		count, err = s.Repository.GetCountFeatureSubWithAuthorization(ctx, ctx.Auth.MobilePhone, queryEntity.Entity, payload.Id)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		return nil
	}); err != nil {
		return &result, err
	}
	result = dto.GetFeatureSubResponse{
		Data:  *data,
		Count: count,
	}
	return &result, nil
}
