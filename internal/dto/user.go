package dto

// Request Change Passwords
type UserChangePasswordRequest struct {
	OldPassword string `json:"old_password" example:"Test1234@" form:"old_password"`
	NewPassword string `json:"new_password" example:"Test12345@" form:"new_password"`
}

type UserChangePasswordResponse struct {
	Message string `json:"message"`
}

type UserChangePasswordRequestParam struct {
	Id int `param:"id" validate:"required"`
}
