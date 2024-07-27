package feature

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	res "compass_mini_api/pkg/util/response"

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

// Get Feature List Items
// @Summary      Get Feature List Items
// @Description  Get Feature List Items
// @Tags         Feature
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param        entity    query     string		true		"entity"		example(Android)
// @Success      200      {object}  dto.GetFeatureListResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/feature/list [get]
func (h *handler) GetFeatureList(c echo.Context) error {
	cc := c.(*abstraction.Context)
	queryEntity := new(abstraction.QueryEntity)
	queryEntity.Entity = cc.QueryParam("entity")
	if err = c.Validate(queryEntity); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	data, err := h.service.GetFeatureList(cc, queryEntity)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

// Get Feature Sub List Items
// @Summary      Get Feature Sub List Items
// @Description  Get Feature Sub List Items
// @Tags         Feature
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param        entity    query     string		true		"entity"		example(Android)
// @Param		id		path		int		true	"id feature"	example(1)
// @Success      200      {object}  dto.GetFeatureSubResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/feature/sub/{id} [get]
func (h *handler) GetFeatureSub(c echo.Context) error {
	cc := c.(*abstraction.Context)
	queryEntity := new(abstraction.QueryEntity)
	queryEntity.Entity = cc.QueryParam("entity")
	if err = c.Validate(queryEntity); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	payload := new(dto.GetFeatureSubRequestParam)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.GetFeatureSub(cc, payload, queryEntity)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
