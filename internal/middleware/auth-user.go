package middleware

import (
	"gorm.io/gorm"

	"compass_mini_api/internal/abstraction"
	dto "compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/repository"
)

type UserAuthService interface {
	GetUserByUserId(*abstraction.Context, *dto.AuthMiddlewareRequest) (*dto.AuthMiddlewareResponse, error)
}

type userAuthService struct {
	Repository repository.User
	Db         *gorm.DB
}

func NewUserAuthService(f *factory.Factory) *userAuthService {
	return &userAuthService{
		Repository: f.UserRepository,
		Db:         f.Db,
	}
}

func (s *userAuthService) GetUserByUserId(ctx *abstraction.Context, payload *dto.AuthMiddlewareRequest) (result *dto.AuthMiddlewareResponse, err error) {

	data, err := s.Repository.FindById(ctx, &payload.ID)
	if err != nil {
		return nil, err
	}

	result = &dto.AuthMiddlewareResponse{
		ID: data.Id,
	}

	return result, nil

}
