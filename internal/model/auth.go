package model

type AuthenticationEntity struct {
	Id                   *int    `json:"id" form:"id" gorm:"column:id"`
	Name                 *string `json:"name" form:"name" gorm:"column:name"`
	Email                *string `json:"email" form:"email" gorm:"column:email"`
	IdentityId           *int    `json:"identityid" form:"identityid" gorm:"column:identityid"`
	MobilePhone          *string `json:"mobilephone" form:"mobilephone" gorm:"column:mobilephone"`
	IsActive             *bool   `json:"isactive" form:"isactive" gorm:"column:isactive"`
	IsLocked             *bool   `json:"islocked" form:"islocked" gorm:"column:islocked"`
	Password             *string `json:"password" form:"password" gorm:"column:password"`
	AccessName           *string `json:"accessname" form:"accessname" gorm:"column:accessname"`
	Entity               *string `json:"entity" form:"entity" gorm:"column:entity"`
	RoleId               *int    `json:"roleid" form:"roleid" gorm:"column:roleid"`
	Role                 *string `json:"role" form:"role" gorm:"column:role"`
	AllowView            *bool   `json:"allowview" form:"allowview" gorm:"column:allowview"`
	AllowAdd             *bool   `json:"allowadd" form:"allowadd" gorm:"column:allowadd"`
	AllowUpdate          *bool   `json:"allowupdate" form:"allowupdate" gorm:"column:allowupdate"`
	AllowDelete          *bool   `json:"allowdelete" form:"allowdelete" gorm:"column:allowdelete"`
	AllowPrint           *bool   `json:"allowprint" form:"allowprint" gorm:"column:allowprint"`
	AllowAccessMobile    *bool   `json:"allowaccessmobile" form:"allowaccessmobile" gorm:"column:allowaccessmobile"`
	AllowAccessWeb       *bool   `json:"allowaccessweb" form:"allowaccessweb" gorm:"column:allowaccessweb"`
	AllowAccessEngine    *bool   `json:"allowaccessengine" form:"allowaccessengine" gorm:"column:allowaccessengine"`
	IsAdministrator      *bool   `json:"isadministrator" form:"isadministrator" gorm:"column:isadministrator"`
	FeatureId            *int    `json:"featureid" form:"featureid" gorm:"column:featureid"`
	Feature              *string `json:"feature" form:"feature" gorm:"column:feature"`
	ParentFeature        *string `json:"parentfeature" form:"parentfeature" gorm:"column:parentfeature"`
	Path                 *string `json:"path" form:"path" gorm:"column:path"`
	FeatureAuthorization *string `json:"featureauthorization" form:"featureauthorization" gorm:"column:featureauthorization"`
}

type FeatureListResponse struct {
	Id          *int    `json:"id" gorm:"column:featureid"`
	Name        *string `json:"name" gorm:"column:feature"`
	AllowView   *bool   `json:"allowview" gorm:"column:allowview"`
	AllowAdd    *bool   `json:"allowadd" gorm:"column:allowadd"`
	AllowUpdate *bool   `json:"allowupdate" gorm:"column:allowupdate"`
	AllowDelete *bool   `json:"allowdelete" gorm:"column:allowdelete"`
	AllowPrint  *bool   `json:"allowprint" gorm:"column:allowprint"`
}

type FeatureListCountResponse struct {
	Count *int `json:"count"`
}

type FeatureSubResponse struct {
	Id          *int    `json:"id" gorm:"column:featureid"`
	Name        *string `json:"name" gorm:"column:feature"`
	AllowView   *bool   `json:"allowview" gorm:"column:allowview"`
	AllowAdd    *bool   `json:"allowadd" gorm:"column:allowadd"`
	AllowUpdate *bool   `json:"allowupdate" gorm:"column:allowupdate"`
	AllowDelete *bool   `json:"allowdelete" gorm:"column:allowdelete"`
	AllowPrint  *bool   `json:"allowprint" gorm:"column:allowprint"`
}

type FeatureSubCountResponse struct {
	Count *int `json:"count"`
}
