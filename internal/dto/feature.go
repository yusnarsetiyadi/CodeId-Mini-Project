package dto

import "compass_mini_api/internal/model"

type GetFeatureListResponse struct {
	Data  []model.FeatureListResponse `json:"item"`
	Count *int                        `json:"count"`
}

type GetFeatureSubRequestParam struct {
	Id int `param:"id" validate:"required"`
}

type GetFeatureSubResponse struct {
	Data  []model.FeatureSubResponse `json:"item"`
	Count *int                       `json:"count"`
}
