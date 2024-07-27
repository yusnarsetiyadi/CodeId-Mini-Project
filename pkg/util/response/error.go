package response

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus/v2"
)

type errorResponse struct {
	Meta        Meta        `json:"meta"`
	Error       interface{} `json:"data"`
	Description interface{} `json:"description"`
}

type Error struct {
	Header       *http.Header
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error
}

const (
	E_BAD_REQUEST          = "bad_request"
	E_DUPLICATE            = "duplicate"
	E_INVALID_PIN_OR_PASS  = "invalid_pin_pass"
	E_NOT_FOUND            = "not_found"
	E_SERVER_ERROR         = "server_error"
	E_TOO_MANY_REQUEST     = "too_many_request"
	E_UNAUTHORIZED         = "unauthorized"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
)

type errorConstant struct {
	Duplicate               Error
	NotFound                Error
	RouteNotFound           Error
	UnprocessableEntity     Error
	Unauthorized            Error
	BadRequest              Error
	Validation              Error
	InternalServerError     Error
	NotFileUpload           Error
	UploadFileSrcError      Error
	UploadFileCreateError   Error
	UploadFileDestError     Error
	UploadFileError         Error
	NotPdfFileError         Error
	BadOtpRequest           Error
	InvalidApiOtpRequest    Error
	InvalidPinOrPassRequest Error
	NotFoundTokenOtpRequest Error

	ErrRequestBodyMdd       Error
	ErrRequestHttpClientMdd Error
	ErrNewRequestMdd        Error
	ErrResponseMdd          Error

	EncodeRequestErr  Error
	EncodePayloadErr  Error
	EncodeResponseErr Error

	TooManyRequest func(retryAfterSecond float64) *Error
}

var (
	ErrorConstant = errorConstant{
		EncodeResponseErr: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Error Encode Response",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusConflict,
		},
		EncodePayloadErr: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Error Encode Payload",
				},
				Error: E_BAD_REQUEST,
			},
			Code: http.StatusBadRequest,
		},
		EncodeRequestErr: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Error Encode Request",
				},
				Error: E_BAD_REQUEST,
			},
			Code: http.StatusBadRequest,
		},
		ErrResponseMdd: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Error struct For Response mdd",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusInternalServerError,
		},
		ErrNewRequestMdd: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Error New Request mdd",
				},
				Error: E_DUPLICATE,
			},
			Code: http.StatusConflict,
		},
		ErrRequestHttpClientMdd: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Error Http Client mdd",
				},
				Error: E_DUPLICATE,
			},
			Code: http.StatusConflict,
		},
		ErrRequestBodyMdd: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Error request body mdd",
				},
				Error: E_DUPLICATE,
			},
			Code: http.StatusConflict,
		},
		Duplicate: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Created value already exists",
				},
				Error: E_DUPLICATE,
			},
			Code: http.StatusConflict,
		},
		NotFound: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Data not found",
				},
				Error: E_NOT_FOUND,
			},
			Code: http.StatusNotFound,
		},
		RouteNotFound: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Route not found",
				},
				Error: E_NOT_FOUND,
			},
			Code: http.StatusNotFound,
		},
		UnprocessableEntity: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Invalid parameters or payload",
				},
				Error: E_UNPROCESSABLE_ENTITY,
			},
			Code: http.StatusUnprocessableEntity,
		},
		Unauthorized: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Unauthorized, please login",
				},
				Error: E_UNAUTHORIZED,
			},
			Code: http.StatusUnauthorized,
		},
		BadRequest: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Bad Request",
				},
				Error: E_BAD_REQUEST,
			},
			Code: http.StatusBadRequest,
		},
		Validation: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Invalid parameters or payload",
				},
				Error: E_BAD_REQUEST,
			},
			Code: http.StatusBadRequest,
		},
		InternalServerError: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Something bad happened",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusInternalServerError,
		},
		NotFileUpload: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "No files to upload",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusInternalServerError,
		},
		UploadFileSrcError: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Failed to open uploaded file",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusInternalServerError,
		},
		UploadFileCreateError: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Failed to create uploaded file",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusInternalServerError,
		},
		UploadFileDestError: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Failed destination uploaded file",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusInternalServerError,
		},
		UploadFileError: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Failed to upload file",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusInternalServerError,
		},
		NotPdfFileError: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Document is not a pdf file",
				},
				Error: E_SERVER_ERROR,
			},
			Code: http.StatusInternalServerError,
		},
		BadOtpRequest: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Bad Request OTP",
				},
				Error: E_BAD_REQUEST,
			},
			Code: http.StatusBadRequest,
		},
		InvalidApiOtpRequest: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Invalid API Key",
				},
				Error: E_BAD_REQUEST,
			},
			Code: http.StatusBadRequest,
		},
		InvalidPinOrPassRequest: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Invalid Pin or Pass Request",
				},
				Error: E_INVALID_PIN_OR_PASS,
			},
			Code: http.StatusForbidden,
		},
		NotFoundTokenOtpRequest: Error{
			Response: errorResponse{
				Meta: Meta{
					Success: false,
					Message: "Token not found",
				},
				Error: E_NOT_FOUND,
			},
			Code: http.StatusNotFound,
		},
		TooManyRequest: func(retryAfterSecond float64) *Error {
			if retryAfterSecond = math.Round(retryAfterSecond); retryAfterSecond == 0 {
				retryAfterSecond = 1
			}
			return &Error{
				Header: &http.Header{echo.HeaderRetryAfter: []string{strconv.Itoa(int(retryAfterSecond))}},
				Response: errorResponse{
					Meta: Meta{
						Success: false,
						Message: "To Many Request",
					},
					Error: E_TOO_MANY_REQUEST,
				},
				Code: http.StatusTooManyRequests,
			}
		},
	}
)

func ErrorBuilder(res *Error, message error, vals ...interface{}) *Error {
	res.ErrorMessage = message
	res.Response.Description = vals
	return res
}

func CustomErrorBuilder(code int, err interface{}, message string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
			},
			Error: err,
		},
		Code: code,
	}
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	} else {
		return ErrorBuilder(&ErrorConstant.InternalServerError, err, err.Error())
	}
}

func (e *Error) Error() string {
	if e.ErrorMessage == nil {
		e.ErrorMessage = errors.New(http.StatusText(e.Code))
	}
	return fmt.Sprintf("error code '%d' because: %s ", e.Code, e.ErrorMessage.Error())
}

func (e *Error) ParseToError() error {
	return e
}

func (e *Error) WithData(data interface{}) *Error {
	e.Response.Error = data
	return e
}

func (e *Error) WithMetaMessage(message string) *Error {
	e.Response.Meta.Message = message
	return e
}

func (e *Error) Send(c echo.Context) error {
	var errorMessage string
	if e.ErrorMessage != nil {
		errorMessage = fmt.Sprintf("%+v", errors.WithStack(e.ErrorMessage))
	}
	logrus.WithFields(apmlogrus.TraceContext(c.Request().Context())).Error(errorMessage)

	if e.Header != nil {
		for k, values := range *e.Header {
			for _, v := range values {
				c.Response().Header().Add(k, v)
			}
		}
	}

	return c.JSON(e.Code, e.Response)
}

func CustomErrorBuilderWithData(code int, data interface{}, message string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
			},
			Error: data,
		},
		Code: code,
	}
}
