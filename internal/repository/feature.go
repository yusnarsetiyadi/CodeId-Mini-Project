package repository

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/model"

	"gorm.io/gorm"
)

type Feature interface {
	GetAllFeatureListWithAuthorization(ctx *abstraction.Context, mobilephone string, entity string) (*[]model.FeatureListResponse, error)
	GetCountFeatureListWithAuthorization(ctx *abstraction.Context, mobilephone string, entity string) (*int, error)
	GetAllFeatureSubWithAuthorization(ctx *abstraction.Context, mobilephone string, entity string, id int) (*[]model.FeatureSubResponse, error)
	GetCountFeatureSubWithAuthorization(ctx *abstraction.Context, mobilephone string, entity string, id int) (*int, error)
}

type feature struct {
	abstraction.Repository
}

func NewFeature(db *gorm.DB) *feature {
	return &feature{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *feature) GetAllFeatureListWithAuthorization(ctx *abstraction.Context, mobilephone string, entity string) (*[]model.FeatureListResponse, error) {
	conn := r.CheckTrx(ctx)

	var data []model.FeatureListResponse
	err := conn.Table("appsetting.user a").
		Select("AllowView, AllowAdd, AllowUpdate, AllowDelete, AllowPrint, b.featureid, d.name feature").
		Joins("INNER JOIN appsetting.access b ON a.RoleID= b.RoleID").
		Joins("INNER JOIN appsetting.role c ON b.RoleID = c.ID").
		Joins("INNER JOIN appsetting.feature d ON b.FeatureID = d.ID").
		Joins("INNER JOIN appsetting.entity e ON b.EntityID = e.ID").
		Where("path IS NULL AND d.isactive = ? AND mobilephone = ? AND Entity = ? AND (IsAdministrator = ? OR AllowAccessMobile = ?) AND AllowView = ?", true, mobilephone, entity, true, true, true).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *feature) GetCountFeatureListWithAuthorization(ctx *abstraction.Context, mobilephone string, entity string) (*int, error) {
	conn := r.CheckTrx(ctx)

	var data model.FeatureListCountResponse
	err := conn.Table("appsetting.user a").
		Select("COUNT(*) AS count").
		Joins("INNER JOIN appsetting.access b ON a.RoleID= b.RoleID").
		Joins("INNER JOIN appsetting.role c ON b.RoleID = c.ID").
		Joins("INNER JOIN appsetting.feature d ON b.FeatureID = d.ID").
		Joins("INNER JOIN appsetting.entity e ON b.EntityID = e.ID").
		Where("path IS NULL AND d.isactive = ? AND mobilephone = ? AND Entity = ? AND (IsAdministrator = ? OR AllowAccessMobile = ?) AND AllowView = ?", true, mobilephone, entity, true, true, true).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data.Count, nil
}

func (r *feature) GetAllFeatureSubWithAuthorization(ctx *abstraction.Context, mobilephone string, entity string, id int) (*[]model.FeatureSubResponse, error) {
	conn := r.CheckTrx(ctx)

	var data []model.FeatureSubResponse
	err := conn.Table("appsetting.user a").
		Select("AllowView, AllowAdd, AllowUpdate, AllowDelete, AllowPrint, b.featureid, d.name feature").
		Joins("INNER JOIN appsetting.access b ON a.RoleID= b.RoleID").
		Joins("INNER JOIN appsetting.role c ON b.RoleID = c.ID").
		Joins("INNER JOIN appsetting.feature d ON b.FeatureID = d.ID").
		Joins("INNER JOIN appsetting.entity e ON b.EntityID = e.ID").
		Where("parentid = ? AND d.isactive = ? AND mobilephone = ? AND Entity = ? AND (IsAdministrator = ? OR AllowAccessMobile = ?) AND AllowView = ?", id, true, mobilephone, entity, true, true, true).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *feature) GetCountFeatureSubWithAuthorization(ctx *abstraction.Context, mobilephone string, entity string, id int) (*int, error) {
	conn := r.CheckTrx(ctx)

	var data model.FeatureSubCountResponse
	err := conn.Table("appsetting.user a").
		Select("COUNT(*) AS count").
		Joins("INNER JOIN appsetting.access b ON a.RoleID= b.RoleID").
		Joins("INNER JOIN appsetting.role c ON b.RoleID = c.ID").
		Joins("INNER JOIN appsetting.feature d ON b.FeatureID = d.ID").
		Joins("INNER JOIN appsetting.entity e ON b.EntityID = e.ID").
		Where("parentid = ? AND d.isactive = ? AND mobilephone = ? AND Entity = ? AND (IsAdministrator = ? OR AllowAccessMobile = ?) AND AllowView = ?", id, true, mobilephone, entity, true, true, true).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data.Count, nil
}
