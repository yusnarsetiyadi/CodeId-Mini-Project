package parameteritem

import (
	"compass_mini_api/internal/abstraction"
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

// Get All Division
// @Summary      Get All Division
// @Description  Get All Division
// @Tags         ParameterItem
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param        limit    query     int		true		"limit"		example(10)
// @Param        offset    query     int		true		"offset"	example(0)
// @Param        conditions    query     string		false		"filter by conditions with query encode value"	example(<br>example string json: [{"column":"name","value":"Maintenance","comparation":"="}]<br>example query encode: %5B%7B%22column%22%3A%22name%22%2C%22value%22%3A%22Maintenance%22%2C%22comparation%22%3A%22%3D%22%7D%5D)
// @Success      200      {object}  dto.GetAllDivisionResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/parameteritem/get_all_division [get]
func (h *handler) GetAllDivision(c echo.Context) error {
	cc := c.(*abstraction.Context)
	queryPagination := new(abstraction.QueryPagination)
	queryFilter := new(abstraction.QueryFilter)
	queryPagination.Limit = cc.QueryParam("limit")
	queryPagination.Offset = cc.QueryParam("offset")
	queryFilter.Conditions = cc.QueryParam("conditions")
	if err = c.Validate(queryPagination); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	data, err := h.service.GetAllDivision(cc, queryPagination, queryFilter)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
