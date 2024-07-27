package repository

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/model"

	"gorm.io/gorm"
)

type User interface {
	CheckAuthentication(ctx *abstraction.Context, mobilephone string, entity string) (*[]model.AuthenticationEntity, error)
	FindByPhoneQuery(ctx *abstraction.Context, mobilephone *string) (*model.UserLoginModel, error)
	FindById(*abstraction.Context, *int) (*model.UserEntityModel, error)
	Update(ctx *abstraction.Context, id *int, data *model.UserEntityModel) (*model.UserEntityModel, error)
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *user) CheckAuthentication(ctx *abstraction.Context, mobilephone string, entity string) (*[]model.AuthenticationEntity, error) {
	conn := r.CheckTrx(ctx)

	var data []model.AuthenticationEntity
	err := conn.Table("appsetting.user a").
		Select("a.id, a.email, a.name, identityid, a.mobilephone, a.isactive, a.islocked, password, b.name accessname, entity, b.roleid, b.role, AllowView, AllowAdd, AllowUpdate, AllowDelete, AllowPrint, AllowAccessMobile, AllowAccessWeb, IsAdministrator, b.featureid, d.name feature, d.ParentFeature ParentFeature, d.Path, CASE WHEN AllowView = true THEN 'V' ELSE '' END || CASE WHEN AllowAdd = true THEN 'A' ELSE '' END || CASE WHEN AllowUpdate = true THEN 'U' ELSE '' END || CASE WHEN AllowDelete = true THEN 'D' ELSE '' END || CASE WHEN AllowPrint = true THEN 'P' ELSE '' END as FeatureAuthorization").
		Joins("INNER JOIN appsetting.access b ON a.RoleID= b.RoleID").
		Joins("INNER JOIN appsetting.role c ON b.RoleID = c.ID").
		Joins("INNER JOIN appsetting.feature d ON b.FeatureID = d.ID").
		Joins("INNER JOIN appsetting.entity e ON b.EntityID = e.ID").
		Where("a.isactive = ? AND a.islocked = ? AND c.IsActive = ? AND d.IsActive = ? AND AllowView = ? AND MobilePhone = ? AND Lower(Entity) = Lower(?) AND (IsAdministrator = ? OR AllowAccessMobile = ?)", true, false, true, true, true, mobilephone, entity, true, true).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) FindByPhoneQuery(ctx *abstraction.Context, mobilephone *string) (*model.UserLoginModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UserLoginModel
	err := conn.Table("appsetting.user").
		Select("id, identityid, name, email, mobilephone, roleid, role, password, salt, isactive, islocked, accesstoken, refreshtoken, resettoken, loginattempt").
		Where("mobilephone = ? AND isactive = ? AND islocked = ?", mobilephone, true, false).
		Take(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) FindById(ctx *abstraction.Context, ID *int) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("id = ? AND isactive = ? AND islocked = ?", ID, true, false).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Update(ctx *abstraction.Context, id *int, data *model.UserEntityModel) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Model(data).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
