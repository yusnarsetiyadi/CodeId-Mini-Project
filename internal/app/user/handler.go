package user

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	res "compass_mini_api/pkg/util/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

// Change Password
// @Summary Change Password
// @Description Change Password
// @Tags User
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param		id		path		int		true	"id user"	example(1)
// @Param        request  body      dto.UserChangePasswordRequest  true  "request body"
// @Success 	200		{object} dto.UserChangePasswordResponse
// @Failure 	400 	{object} res.errorResponse
// @Failure 	404 	{object} res.errorResponse
// @Failure 	500 	{object} res.errorResponse
// @Router		/api/v1/user/change_password/{id} [post]
func (h *handler) ChangePassword(c echo.Context) error {
	cc := c.(*abstraction.Context)
	paramId := new(dto.UserChangePasswordRequestParam)
	paramId.Id, _ = strconv.Atoi(cc.Param("id"))
	if err := c.Validate(paramId); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	payload := new(dto.UserChangePasswordRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.ChangePassword(cc, payload, paramId)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
