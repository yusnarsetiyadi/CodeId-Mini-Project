package auth

import (
	"github.com/labstack/echo/v4"

	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"

	res "compass_mini_api/pkg/util/response"
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

// Login
// @Summary      Login user
// @Description  Login user, get your Geolocation from https://ip.nf/me.json
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        entity    query     string		true		"entity"		example(Android)
// @Param        request  body      dto.AuthLoginRequest  true  "request body"
// @Success      200      {object}  dto.AuthLoginResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/auth/login [post]
func (h *handler) Login(c echo.Context) error {
	cc := c.(*abstraction.Context)
	queryEntity := new(abstraction.QueryEntity)
	queryEntity.Entity = cc.QueryParam("entity")
	if err = c.Validate(queryEntity); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	payload := new(dto.AuthLoginRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.Login(cc, payload, queryEntity)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

// Splash
// @Summary      Splash screen
// @Description  Splash screen
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AuthSplashRequest  true  "request body"
// @Success      200      {object}  dto.AuthSplashResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/auth/splash [post]
func (h *handler) Splash(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.AuthSplashRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.Splash(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

// Logout
// @Summary      Logout user
// @Description  Logout user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Success      200      {object}  dto.AuthLogoutResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/auth/logout [delete]
func (h *handler) Logout(c echo.Context) error {
	cc := c.(*abstraction.Context)

	data, err := h.service.Logout(cc)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

// Get Data Token
// @Summary      Get Data Token
// @Description  Get Data Token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Success      200      {object}  dto.GetDataTokenResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/auth/get_data_token [get]
func (h *handler) GetDataToken(c echo.Context) error {
	cc := c.(*abstraction.Context)

	data, err := h.service.GetDataToken(cc)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
