package user

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/model"
	"compass_mini_api/internal/repository"
	res "compass_mini_api/pkg/util/response"
	"compass_mini_api/pkg/util/trxmanager"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	ChangePassword(ctx *abstraction.Context, payload *dto.UserChangePasswordRequest, paramId *dto.UserChangePasswordRequestParam) (*dto.UserChangePasswordResponse, error)
}

type service struct {
	Repository repository.User
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.UserRepository
	db := f.Db
	return &service{
		repository,
		db,
	}
}

func (s *service) ChangePassword(ctx *abstraction.Context, payload *dto.UserChangePasswordRequest, paramId *dto.UserChangePasswordRequestParam) (*dto.UserChangePasswordResponse, error) {

	if paramId.Id != ctx.Auth.ID {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, "failed change password", "your request id is not matching!")
	}

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		userData, err := s.Repository.FindById(ctx, &paramId.Id)
		if err != nil && err.Error() == "record not found" {
			return res.CustomErrorBuilder(http.StatusInternalServerError, "failed get users data", "user not found or already deleted")
		} else if err != nil && err.Error() != "record not found" {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		salt := *userData.Salt
		inputPassword := salt + payload.OldPassword + salt
		if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(inputPassword)); err != nil {
			return res.CustomErrorBuilder(http.StatusBadRequest, "Request Failed", "Your password is wrong!")
		}

		if payload.OldPassword == payload.NewPassword {
			return res.CustomErrorBuilder(http.StatusBadRequest, "Request Failed", "The new password cannot be the same as the old password!")
		}

		newPassword := salt + payload.NewPassword + salt
		password := []byte(newPassword)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		newData := new(model.UserEntityModel)
		newData.Password = string(hashedPassword)
		newData.Id = userData.Id
		_, err = s.Repository.Update(ctx, &userData.Id, newData)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	result := dto.UserChangePasswordResponse{
		Message: "Success change password",
	}

	return &result, nil
}
