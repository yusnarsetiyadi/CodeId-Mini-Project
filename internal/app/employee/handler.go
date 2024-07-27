package employee

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

// Create Employee
// @Summary      Create employee
// @Description  Create employee
// @Tags         Employee
// @Accept       mpfd
// @Produce      json
// @Security 	 BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param       request  formData    dto.CreateEmployeeRequest  true  "request body"
// @Param       employeephoto   formData    file    false   "image"
// @Success      200      {object}  dto.ResponseMessage
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/employee [post]
func (h *handler) Create(c echo.Context) error {
	cc := c.(*abstraction.Context)
	file, err := c.FormFile("employeephoto")
	if err != nil && err.Error() != "http: no such file" {
		return res.ErrorBuilder(&res.ErrorConstant.UploadFileError, err).Send(c)
	}
	payload := new(dto.CreateEmployeeRequest)
	payload.EmployeePhoto = file
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	// Source

	result, err := h.service.Create(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Get Employee By Id
// @Summary      Get Employee By Id
// @Description  Get Employee By Id
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param		id		path		int		true	"id employee"	example(1)
// @Success      200      {object}  model.EmployeeGetByIdResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/employee/{id} [get]
func (h *handler) GetEmployeeById(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.GetEmployeeByIdRequestParam)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.GetEmployeeById(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Get All Employee
// @Summary      Get All Employee
// @Description  Get All Employee
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param        limit    query     int		true		"limit"		example(10)
// @Param        offset    query     int		true		"offset"	example(0)
// @Param        order    query     string		false		"order"		example(name / company / division / supervisor / email / joindate)
// @Param        direction    query     string		false		"direction"	example(asc / desc)
// @Param        conditions    query     string		false		"filter by conditions with query encode value"	example(<br>example string json: [{"column":"name","value":"yusnar","comparation":"%"},{"column":"companyid","value":"43","comparation":"="},{"column":"joindate","value":"2024-01-24_2024-01-28","comparation":"between"}]<br>example query encode: %5B%7B%22column%22%3A%22name%22%2C%22value%22%3A%22yusnar%22%2C%22comparation%22%3A%22%25%22%7D%2C%7B%22column%22%3A%22companyid%22%2C%22value%22%3A%2243%22%2C%22comparation%22%3A%22%3D%22%7D%2C%7B%22column%22%3A%22joindate%22%2C%22value%22%3A%222024-01-24_2024-01-28%22%2C%22comparation%22%3A%22between%22%7D%5D)
// @Success      200      {object}  dto.EmployeeGetAllResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/employee [get]
func (h *handler) GetAllEmployee(c echo.Context) error {
	cc := c.(*abstraction.Context)
	queryPagination := new(abstraction.QueryPagination)
	queryOrder := new(abstraction.QueryOrder)
	queryFilter := new(abstraction.QueryFilter)
	queryPagination.Limit = cc.QueryParam("limit")
	queryPagination.Offset = cc.QueryParam("offset")
	queryOrder.Order = cc.QueryParam("order")
	queryOrder.Direction = cc.QueryParam("direction")
	queryFilter.Conditions = cc.QueryParam("conditions")
	if err = c.Validate(queryPagination); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.GetAllEmployee(cc, queryPagination, queryOrder, queryFilter)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update Employee
// @Summary      Update employee
// @Description  Update employee
// @Tags         Employee
// @Accept       mpfd
// @Produce      json
// @Security 	 BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param		id		path		int		true	"id employee"	example(1)
// @Param       request  formData      dto.UpdateEmployeeRequest  false  "request body"
// @Param       employeephoto   formData    file    false   "image"
// @Success      200      {object}  dto.ResponseMessage
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/employee/{id} [put]
func (h *handler) Update(c echo.Context) error {
	cc := c.(*abstraction.Context)
	paramId := new(dto.UpdateEmployeeRequestParam)
	paramId.Id, _ = strconv.Atoi(cc.Param("id"))
	if err := c.Validate(paramId); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	file, err := c.FormFile("employeephoto")
	if err != nil && err.Error() != "http: no such file" {
		return res.ErrorBuilder(&res.ErrorConstant.UploadFileError, err).Send(c)
	}
	payload := new(dto.UpdateEmployeeRequest)
	payload.EmployeePhoto = file
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.Update(cc, payload, paramId)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Get All Employee for Supervisor
// @Summary      Get All Employee for Supervisor
// @Description  Get All Employee for Supervisor
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param       Authorization       header      string      true        "Bearer Token"
// @Param        limit    query     int     true        "limit"     example(10)
// @Param        offset    query     int        true        "offset"    example(0)
// @Param        conditions    query     string		false		"filter by conditions with query encode value"	example(<br>example string json: [{"column":"name","value":"herru","comparation":"%"}]<br>example query encode: %5B%7B%22column%22%3A%22name%22%2C%22value%22%3A%22herru%22%2C%22comparation%22%3A%22%25%22%7D%5D)
// @Success      200      {object}  dto.EmployeeSupervisorGetAllResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/employee/supervisor [get]
func (h *handler) GetAllEmployeeSupervisor(c echo.Context) error {
	cc := c.(*abstraction.Context)
	queryPagination := new(abstraction.QueryPagination)
	queryFilter := new(abstraction.QueryFilter)
	queryPagination.Limit = cc.QueryParam("limit")
	queryPagination.Offset = cc.QueryParam("offset")
	queryFilter.Conditions = cc.QueryParam("conditions")
	if err = c.Validate(queryPagination); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.GetAllEmployeeSupervisor(cc, queryPagination, queryFilter)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Create Employee With Base64
// @Summary      Create employee With Base64
// @Description  Create employee With Base64, please fill data base64 with Plain-text - just The Base64 value, see https://base64.guru/converter/encode/file
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param        request  body      dto.CreateEmployeeRequestWithBase64  true  "request body"
// @Success      200      {object}  dto.ResponseMessage
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/employee/with_base64 [post]
func (h *handler) CreateWithBase64(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.CreateEmployeeRequestWithBase64)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.CreateWithBase64(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update Employee With Base64
// @Summary      Update employee With Base64
// @Description  Update employee With Base64, please fill data base64 with Plain-text - just The Base64 value, see https://base64.guru/converter/encode/file
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param		id		path		int		true	"id employee"	example(1)
// @Param        request  body      dto.UpdateEmployeeRequestWithBase64  false  "request body"
// @Success      200      {object}  dto.ResponseMessage
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/employee/with_base64/{id} [put]
func (h *handler) UpdateWithBase64(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.UpdateEmployeeRequestWithBase64)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.UpdateWithBase64(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Get Employee Photo
// @Summary      Get Employee Photo
// @Description  Get Employee Photo
// @Tags         Employee
// @Accept       json
// @Produce      mpfd
// @Security BearerAuth
// @Param		Authorization		header		string		true		"Bearer Token"
// @Param		base64		path		string		true	"base64"
// @Success      200
// @Failure      400
// @Failure      401
// @Failure      404
// @Failure      500
// @Router       /api/v1/employee/employeephoto/{base64} [get]
func (h *handler) GetEmployeePhoto(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.GetEmployeePhotoRequestParam)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.GetEmployeePhoto(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SendEmployeePhoto(c, result)
}
