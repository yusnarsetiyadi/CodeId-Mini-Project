package dto

import "compass_mini_api/internal/model"

type GetAllDivisionResponse struct {
	Data  []model.GetAllDivisionResponse `json:"item"`
	Count *int                           `json:"count"`
}
