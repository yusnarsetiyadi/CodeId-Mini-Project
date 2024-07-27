package dto

import "compass_mini_api/internal/model"

type GetAllCompanyResponse struct {
	Data  []model.CompanyEntityModel `json:"item"`
	Count *int                       `json:"count"`
}
