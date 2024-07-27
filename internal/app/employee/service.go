package employee

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/model"
	"compass_mini_api/internal/repository"
	"compass_mini_api/pkg/util/general"
	res "compass_mini_api/pkg/util/response"
	"compass_mini_api/pkg/util/trxmanager"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	GetAllEmployee(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryOrder *abstraction.QueryOrder, queryFilter *abstraction.QueryFilter) (*dto.EmployeeGetAllResponse, error)
	GetAllEmployeeSupervisor(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*dto.EmployeeSupervisorGetAllResponse, error)
	Create(ctx *abstraction.Context, payload *dto.CreateEmployeeRequest) (*dto.ResponseMessage, error)
	GetEmployeeById(ctx *abstraction.Context, payload *dto.GetEmployeeByIdRequestParam) (*model.EmployeeGetByIdResponse, error)
	Update(ctx *abstraction.Context, payload *dto.UpdateEmployeeRequest, paramId *dto.UpdateEmployeeRequestParam) (*dto.ResponseMessage, error)
	CreateWithBase64(ctx *abstraction.Context, payload *dto.CreateEmployeeRequestWithBase64) (*dto.ResponseMessage, error)
	UpdateWithBase64(ctx *abstraction.Context, payload *dto.UpdateEmployeeRequestWithBase64) (*dto.ResponseMessage, error)
	GetEmployeePhoto(ctx *abstraction.Context, payload *dto.GetEmployeePhotoRequestParam) (string, error)
}

type service struct {
	Db         *gorm.DB
	Repository repository.Employee
}

func NewService(f *factory.Factory) Service {
	repository := f.EmployeeRepository
	return &service{
		Db:         f.Db,
		Repository: repository,
	}
}

func (s *service) GetAllEmployeeSupervisor(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryFilter *abstraction.QueryFilter) (*dto.EmployeeSupervisorGetAllResponse, error) {
	var result *dto.EmployeeSupervisorGetAllResponse
	var data *[]model.EmployeeSupervisorGetAllResponse
	var count *int

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		count, err = s.Repository.GetCountEmployeeSupervisor(ctx, queryFilter)
		if err != nil {
			if strings.Contains(err.Error(), "400") {
				return err
			}
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		data, err = s.Repository.GetAllEmployeeSupervisor(ctx, queryPagination, queryFilter)
		if err != nil && err.Error() != "record not found" {
			if strings.Contains(err.Error(), "400") {
				return err
			}
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	result = &dto.EmployeeSupervisorGetAllResponse{
		Data:  *data,
		Count: count,
	}
	return result, nil
}

func (s *service) GetAllEmployee(ctx *abstraction.Context, queryPagination *abstraction.QueryPagination, queryOrder *abstraction.QueryOrder, queryFilter *abstraction.QueryFilter) (*dto.EmployeeGetAllResponse, error) {
	var result *dto.EmployeeGetAllResponse
	var data *[]model.EmployeeGetAllResponse
	var count *int

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		count, err = s.Repository.GetCountEmployee(ctx, queryFilter)
		if err != nil {
			if strings.Contains(err.Error(), "400") {
				return err
			}
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		data, err = s.Repository.GetAllEmployee(ctx, queryPagination, queryOrder, queryFilter)
		if err != nil && err.Error() != "record not found" {
			if strings.Contains(err.Error(), "400") {
				return err
			}
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	result = &dto.EmployeeGetAllResponse{
		Data:  *data,
		Count: count,
	}
	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.CreateEmployeeRequest) (*dto.ResponseMessage, error) {
	var result dto.ResponseMessage
	var dataEmployee = new(model.EmployeeEntityModel)
	var extension string
	if payload.EmployeePhoto != nil {
		src, err := payload.EmployeePhoto.Open()
		if err != nil {
			return nil, res.ErrorBuilder(&res.ErrorConstant.UploadFileSrcError, err)
		}
		defer src.Close()
		extension = filepath.Ext(payload.EmployeePhoto.Filename)
		fileName := fmt.Sprintf("%s%s", time.Now().Format("20060102150405000"), extension)
		destinationPath := path.Join("../employeephoto", fileName)

		dst, err := os.Create(destinationPath)
		if err != nil {
			return nil, res.ErrorBuilder(&res.ErrorConstant.UploadFileCreateError, err)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return nil, res.ErrorBuilder(&res.ErrorConstant.UploadFileDestError, err)
		}
		base64Data := base64.StdEncoding.EncodeToString([]byte(fileName))
		dataEmployee.EmployeePhoto = &base64Data
	}
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		isActive := true
		parsedTime, err := time.Parse("2006-01-02", payload.JoinDate)
		if err != nil {
			return res.CustomErrorBuilder(http.StatusBadRequest, "request failed", "error parsing date!")
		}

		dataEmployee.Name = &payload.Name
		dataEmployee.Email = &payload.Email
		dataEmployee.PhoneNumber = &payload.PhoneNumber
		dataEmployee.CompanyId = &payload.CompanyId
		dataEmployee.Company = &payload.Company
		dataEmployee.DivisionId = &payload.DivisionId
		dataEmployee.Division = &payload.Division
		dataEmployee.IsActive = &isActive
		dataEmployee.SupervisorId = &payload.SupervisorId
		dataEmployee.Supervisor = &payload.Supervisor
		dataEmployee.JoinDate = &parsedTime
		dataSPV, err := s.Repository.FindByIdEmployee(ctx, &payload.SupervisorId)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}
		dataEmployeeNew, err := s.Repository.CreateEmployee(ctx, dataEmployee)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		uniqueKeyNew := *dataSPV.UniqueKey + strconv.Itoa(dataEmployeeNew.Id) + "|"
		newDataEmployee := new(model.EmployeeEntityModel)
		newDataEmployee.Id = dataEmployeeNew.Id
		newDataEmployee.UniqueKey = &uniqueKeyNew
		_, err = s.Repository.UpdateEmployee(ctx, &newDataEmployee.Id, newDataEmployee)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		return nil
	}); err != nil {
		return &result, err
	}
	result = dto.ResponseMessage{
		Message: "Success create data",
	}
	return &result, nil
}

func (s *service) GetEmployeeById(ctx *abstraction.Context, payload *dto.GetEmployeeByIdRequestParam) (*model.EmployeeGetByIdResponse, error) {
	var dataEmployee *model.EmployeeGetByIdResponse
	dataEmployee, err = s.Repository.FindByIdEmployee(ctx, &payload.Id)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
	}
	if dataEmployee.JoinDate != nil {
		joinDate := *dataEmployee.JoinDate
		joinDateSlice := joinDate[:10]
		dataEmployee.JoinDate = &joinDateSlice
	}
	if dataEmployee.ResignDate != nil {
		resignDate := *dataEmployee.ResignDate
		resignDateSlice := resignDate[:10]
		dataEmployee.ResignDate = &resignDateSlice
	}

	return dataEmployee, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.UpdateEmployeeRequest, paramId *dto.UpdateEmployeeRequestParam) (*dto.ResponseMessage, error) {
	var result dto.ResponseMessage
	var newData = new(model.EmployeeEntityModel)
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if payload.JoinDate != nil {
			parsedTimeJoinDate, err := time.Parse("2006-01-02", *payload.JoinDate)
			if err != nil {
				return res.CustomErrorBuilder(http.StatusBadRequest, "request failed", "error parsing join date!")
			}
			newData.JoinDate = &parsedTimeJoinDate
		}
		if payload.ResignDate != nil {
			parsedTimeResignDate, err := time.Parse("2006-01-02", *payload.ResignDate)
			if err != nil {
				return res.CustomErrorBuilder(http.StatusBadRequest, "request failed", "error parsing resign date!")
			}
			newData.ResignDate = &parsedTimeResignDate
		}
		if payload.EmployeePhoto != nil {
			src, err := payload.EmployeePhoto.Open()
			if err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.UploadFileSrcError, err)
			}
			defer src.Close()

			fileName := fmt.Sprintf("%s%s", time.Now().Format("20060102150405000"), filepath.Ext(payload.EmployeePhoto.Filename))

			destinationPath := path.Join("../employeephoto", fileName)
			dst, err := os.Create(destinationPath)
			if err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.UploadFileCreateError, err)
			}
			defer dst.Close()

			if _, err = io.Copy(dst, src); err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.UploadFileDestError, err)
			}
			base64Data := base64.StdEncoding.EncodeToString([]byte(fileName))
			newData.EmployeePhoto = &base64Data
		}
		newData.Name = payload.Name
		newData.Email = payload.Email
		newData.PhoneNumber = payload.PhoneNumber
		newData.CompanyId = payload.CompanyId
		newData.Company = payload.Company
		newData.DivisionId = payload.DivisionId
		newData.Division = payload.Division
		newData.IsActive = payload.IsActive
		if payload.SupervisorId != nil && payload.Supervisor != nil {
			dataSPV, err := s.Repository.FindByIdEmployee(ctx, payload.SupervisorId)
			if err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
			}
			uniqueKeyNew := *dataSPV.UniqueKey + strconv.Itoa(paramId.Id) + "|"
			newData.UniqueKey = &uniqueKeyNew
			newData.SupervisorId = payload.SupervisorId
			newData.Supervisor = payload.Supervisor
		}

		_, err = s.Repository.UpdateEmployee(ctx, &paramId.Id, newData)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		return nil
	}); err != nil {
		return nil, err
	}
	result = dto.ResponseMessage{
		Message: "Success update data",
	}
	return &result, nil
}

func (s *service) CreateWithBase64(ctx *abstraction.Context, payload *dto.CreateEmployeeRequestWithBase64) (*dto.ResponseMessage, error) {
	var result dto.ResponseMessage
	var dataEmployee = new(model.EmployeeEntityModel)
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		filename, err := general.SaveFileEmployeePhoto(payload.EmployeePhoto)
		if err != nil {
			return err
		}
		base64Data := base64.StdEncoding.EncodeToString([]byte(*filename))
		isActive := true
		parsedTime, err := time.Parse("2006-01-02", payload.JoinDate)
		if err != nil {
			return res.CustomErrorBuilder(http.StatusBadRequest, "request failed", "error parsing date!")
		}
		dataEmployee.Name = &payload.Name
		dataEmployee.Email = &payload.Email
		dataEmployee.PhoneNumber = &payload.PhoneNumber
		dataEmployee.CompanyId = &payload.CompanyId
		dataEmployee.Company = &payload.Company
		dataEmployee.DivisionId = &payload.DivisionId
		dataEmployee.Division = &payload.Division
		dataEmployee.IsActive = &isActive
		dataEmployee.SupervisorId = &payload.SupervisorId
		dataEmployee.Supervisor = &payload.Supervisor
		dataEmployee.JoinDate = &parsedTime
		dataEmployee.EmployeePhoto = &base64Data
		dataSPV, err := s.Repository.FindByIdEmployee(ctx, &payload.SupervisorId)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		dataEmployeeNew, err := s.Repository.CreateEmployee(ctx, dataEmployee)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		uniqueKeyNew := *dataSPV.UniqueKey + strconv.Itoa(dataEmployeeNew.Id) + "|"
		newDataEmployee := new(model.EmployeeEntityModel)
		newDataEmployee.Id = dataEmployeeNew.Id
		newDataEmployee.UniqueKey = &uniqueKeyNew
		_, err = s.Repository.UpdateEmployee(ctx, &newDataEmployee.Id, newDataEmployee)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		return nil
	}); err != nil {
		return &result, err
	}
	result = dto.ResponseMessage{
		Message: "success create employee with base64",
	}
	return &result, nil
}

func (s *service) UpdateWithBase64(ctx *abstraction.Context, payload *dto.UpdateEmployeeRequestWithBase64) (*dto.ResponseMessage, error) {
	var result dto.ResponseMessage
	var newData = new(model.EmployeeEntityModel)
	idUser, _ := strconv.Atoi(ctx.Param("id"))
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if payload.JoinDate != nil {
			parsedTimeJoinDate, err := time.Parse("2006-01-02", *payload.JoinDate)
			if err != nil {
				return res.CustomErrorBuilder(http.StatusBadRequest, "request failed", "error parsing join date!")
			}
			newData.JoinDate = &parsedTimeJoinDate
		}
		if payload.ResignDate != nil {
			parsedTimeResignDate, err := time.Parse("2006-01-02", *payload.ResignDate)
			if err != nil {
				return res.CustomErrorBuilder(http.StatusBadRequest, "request failed", "error parsing resign date!")
			}
			newData.ResignDate = &parsedTimeResignDate
		}
		if payload.EmployeePhoto != nil {
			filename, err := general.SaveFileEmployeePhoto(*payload.EmployeePhoto)
			if err != nil {
				return err
			}
			base64Data := base64.StdEncoding.EncodeToString([]byte(*filename))
			newData.EmployeePhoto = &base64Data
		}
		newData.Name = payload.Name
		newData.Email = payload.Email
		newData.PhoneNumber = payload.PhoneNumber
		newData.CompanyId = payload.CompanyId
		newData.Company = payload.Company
		newData.DivisionId = payload.DivisionId
		newData.Division = payload.Division
		newData.IsActive = payload.IsActive
		if payload.SupervisorId != nil && payload.Supervisor != nil {
			dataSPV, err := s.Repository.FindByIdEmployee(ctx, payload.SupervisorId)
			if err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
			}
			uniqueKeyNew := *dataSPV.UniqueKey + strconv.Itoa(idUser) + "|"
			newData.UniqueKey = &uniqueKeyNew
			newData.SupervisorId = payload.SupervisorId
			newData.Supervisor = payload.Supervisor
		}

		_, err = s.Repository.UpdateEmployee(ctx, &idUser, newData)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		return nil
	}); err != nil {
		return nil, err
	}
	result = dto.ResponseMessage{
		Message: "success update data with base64",
	}
	return &result, nil
}

func (s *service) GetEmployeePhoto(ctx *abstraction.Context, payload *dto.GetEmployeePhotoRequestParam) (string, error) {
	escapedString, _ := url.QueryUnescape(payload.Base64)
	data, err := base64.StdEncoding.DecodeString(escapedString)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	fileName := string(data)
	pathFile := path.Join("../employeephoto", fileName)

	return pathFile, nil
}
