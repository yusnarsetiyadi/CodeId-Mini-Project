package auth

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/middleware"
	"compass_mini_api/internal/model"
	"compass_mini_api/internal/repository"
	"compass_mini_api/pkg/constant"
	"compass_mini_api/pkg/util/date"
	res "compass_mini_api/pkg/util/response"
	"compass_mini_api/pkg/util/trxmanager"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(*abstraction.Context, *dto.AuthLoginRequest, *abstraction.QueryEntity) (*dto.AuthLoginResponse, error)
	Splash(*abstraction.Context, *dto.AuthSplashRequest) (*dto.AuthSplashResponse, error)
	Logout(*abstraction.Context) (*dto.AuthLogoutResponse, error)
	GetDataToken(*abstraction.Context) (*dto.GetDataTokenResponse, error)
}

type service struct {
	Repository repository.User
	Db         *gorm.DB
}

func NewService(f *factory.Factory) Service {
	repository := f.UserRepository
	db := f.Db
	return &service{
		repository,
		db,
	}
}

func (s *service) Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest, queryEntity *abstraction.QueryEntity) (*dto.AuthLoginResponse, error) {
	var (
		result     dto.AuthLoginResponse
		data       = new(model.UserLoginModel)
		dataUser   = new(model.UserEntityModel)
		dataAuth   = new([]model.AuthenticationEntity)
		accToken   *model.AccessTokenClaims
		token      = new(model.Token)
		deviceInfo = ctx.Request().Header.Get("User-Agent")
	)

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.FindByPhoneQuery(ctx, &payload.MobilePhone)
		if err != nil {
			return res.CustomErrorBuilder(http.StatusBadRequest, "Request Failed", "Wrong Password or Phone, Please Input Right Password or Phone")
		}

		inputPassword := data.Salt + payload.Password + data.Salt
		if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(inputPassword)); err != nil {
			return res.CustomErrorBuilder(http.StatusBadRequest, "Request Failed", "Wrong Password or Phone, Please Input Right Password or Phone")
		}

		dataAuth, err = s.Repository.CheckAuthentication(ctx, payload.MobilePhone, queryEntity.Entity)
		if err != nil && err.Error() != "record not found" {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		if len(*dataAuth) == 0 {
			return res.CustomErrorBuilder(http.StatusUnauthorized, "Request Failed", "Failed during authentication")
		}

		if accToken, err = middleware.GenerateEncryptData(data); err != nil {
			return res.CustomErrorBuilder(http.StatusInternalServerError, err.Error(), "Error generate encrypt data access token")
		}

		token.AccessToken, err = model.AccessToken(*accToken)
		if err != nil {
			return res.CustomErrorBuilder(http.StatusInternalServerError, err.Error(), "Error generate accesstoken")
		}

		token.RefreshToken, err = model.RefreshToken(model.RefreshTokenClaims{Exp: time.Now().Add(constant.REFRESH_TOKEN_DURATION).Unix()})
		if err != nil {
			return res.CustomErrorBuilder(http.StatusInternalServerError, err.Error(), "Error generate refresh token")
		}

		jsonData, _ := json.Marshal(payload.Geolocation)
		deviceLocation := string(jsonData)
		loginAttempt := data.LoginAttempt + 1
		dataUser.DeviceInfo = &deviceInfo
		dataUser.DeviceLocation = &deviceLocation
		dataUser.IpAddress = &payload.Geolocation.IP
		dataUser.Latitude = &payload.Geolocation.Latitude
		dataUser.Longitude = &payload.Geolocation.Longitude
		dataUser.LastLogin = date.DateTodayLocal()
		dataUser.AccessToken = &token.AccessToken
		dataUser.RefreshToken = &token.RefreshToken
		dataUser.LoginAttempt = &loginAttempt
		_, err := s.Repository.Update(ctx, &data.Id, dataUser)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		return nil
	}); err != nil {
		return &result, err
	}

	result = dto.AuthLoginResponse{
		Token: *token,
	}

	return &result, nil
}

func (s *service) Splash(ctx *abstraction.Context, payload *dto.AuthSplashRequest) (*dto.AuthSplashResponse, error) {
	var (
		result       dto.AuthSplashResponse
		unauthorized = 401
		data         = new(model.UserLoginModel)
		token        = new(model.Token)
		accToken     *model.AccessTokenClaims
	)

	if _, statusCodeValidateAccToken := middleware.ValidateAccessToken(payload.AccessToken); *statusCodeValidateAccToken == unauthorized {
		if authContextRefToken, statusstatusCodeValidateRefToken := middleware.ValidateRefreshToken(payload.Token); *statusstatusCodeValidateRefToken == unauthorized {
			return nil, res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil)
		} else {
			if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
				data, err = s.Repository.FindByPhoneQuery(ctx, &authContextRefToken.MobilePhone)
				if err != nil {
					return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
				}

				if accToken, err = middleware.GenerateEncryptData(data); err != nil {
					return res.CustomErrorBuilder(http.StatusInternalServerError, err.Error(), "Error generate encrypt data access token")
				}

				token.AccessToken, err = model.AccessToken(*accToken)
				if err != nil {
					return res.CustomErrorBuilder(http.StatusInternalServerError, err.Error(), "Error generate accesstoken")
				}

				return nil
			}); err != nil {
				return &result, err
			}
			result = dto.AuthSplashResponse{
				Token: model.Token{
					AccessToken:  token.AccessToken,
					RefreshToken: payload.RefreshToken,
				},
			}
		}
	} else {
		result = dto.AuthSplashResponse{
			Token: payload.Token,
		}
	}
	return &result, nil
}

func (s *service) Logout(ctx *abstraction.Context) (*dto.AuthLogoutResponse, error) {
	var result dto.AuthLogoutResponse
	var dataUser model.UserEntityModel
	userId := ctx.Auth.ID
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		middleware.ChangeTokenForLogout(*ctx)
		empty := ""
		dataUser.AccessToken = &empty
		dataUser.RefreshToken = &empty
		dataUser.ResetToken = &empty
		_, err = s.Repository.Update(ctx, &userId, &dataUser)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		return nil
	}); err != nil {
		return &result, err
	}

	result = dto.AuthLogoutResponse{
		Message: "Success logout",
	}

	return &result, nil
}

func (s *service) GetDataToken(ctx *abstraction.Context) (*dto.GetDataTokenResponse, error) {
	data := dto.GetDataTokenResponse{
		Id:          ctx.Auth.ID,
		IdentityId:  ctx.Auth.IdentityId,
		RoleId:      ctx.Auth.RoleId,
		Name:        ctx.Auth.Name,
		Email:       ctx.Auth.Email,
		MobilePhone: ctx.Auth.MobilePhone,
		Role:        ctx.Auth.Role,
		IsActive:    ctx.Auth.IsActive,
		IsLocked:    ctx.Auth.IsLocked,
	}

	return &data, nil
}
