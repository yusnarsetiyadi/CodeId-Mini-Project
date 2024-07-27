package middleware

import (
	"compass_mini_api/internal/config"
	dto "compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/model"
	"compass_mini_api/pkg/constant"
	"compass_mini_api/pkg/util/aescrypt"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"compass_mini_api/internal/abstraction"
	res "compass_mini_api/pkg/util/response"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	authService UserAuthService
)

func newUserAuthService(f *factory.Factory) {
	authService = NewUserAuthService(f)
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		jwtKey                         = config.Get().Key.JwtKey
		id, identityid, roleid         int
		name, email, mobilephone, role string
		isactive, islocked             bool
	)

	return func(c echo.Context) error {

		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		tokenString := strings.Replace(authToken, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}
			return []byte(jwtKey), nil
		})

		if err != nil {
			var jwtValErr *jwt.ValidationError
			if errors.As(err, &jwtValErr) {
				if jwtValErr.Errors == jwt.ValidationErrorExpired {
					return res.CustomErrorBuilder(http.StatusUnauthorized, nil, "token_is_expired").Send(c)
				} else {
					return res.CustomErrorBuilder(http.StatusUnauthorized, nil, err.Error()).Send(c)
				}
			} else {
				return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
			}
		}

		if token == nil || !token.Valid {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, errors.New("invalid token")).Send(c)
		}

		destructID := token.Claims.(jwt.MapClaims)["id"]
		if destructID != nil {
			if destructID == "" {
				return res.CustomErrorBuilder(http.StatusUnauthorized, nil, "Unauthorized").Send(c)
			}

			id, err = strconv.Atoi(fmt.Sprintf("%v", destructID))
			if err != nil {
				if destructID, err = aescrypt.DecryptAES(destructID.(string), jwtKey); err != nil {
					return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
				}

				id, err = strconv.Atoi(fmt.Sprintf("%v", destructID))
				if err != nil {
					return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Error parsing data").Send(c)
				}
			}

		} else {
			return res.CustomErrorBuilder(http.StatusUnauthorized, nil, "Unauthorized").Send(c)
		}

		destructIdentityID := token.Claims.(jwt.MapClaims)["identityid"]
		if destructIdentityID != nil {
			if destructIdentityID == "" {
				return res.CustomErrorBuilder(http.StatusUnauthorized, nil, "Unauthorized").Send(c)
			}

			identityid, err = strconv.Atoi(fmt.Sprintf("%v", destructIdentityID))
			if err != nil {
				if destructIdentityID, err = aescrypt.DecryptAES(destructIdentityID.(string), jwtKey); err != nil {
					return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
				}

				identityid, err = strconv.Atoi(fmt.Sprintf("%v", destructIdentityID))
				if err != nil {
					return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Error parsing data").Send(c)
				}
			}

		} else {
			return res.CustomErrorBuilder(http.StatusUnauthorized, nil, "Unauthorized").Send(c)
		}

		destructRoleId := token.Claims.(jwt.MapClaims)["roleid"]
		if destructRoleId != nil {

			if destructRoleId == "" {
				return res.CustomErrorBuilder(http.StatusUnauthorized, nil, "Unauthorized").Send(c)
			}

			roleid, err = strconv.Atoi(fmt.Sprintf("%v", destructRoleId))
			if err != nil {

				if destructRoleId, err = aescrypt.DecryptAES(destructRoleId.(string), jwtKey); err != nil {
					return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
				}

				roleid, err = strconv.Atoi(fmt.Sprintf("%v", destructRoleId))
				if err != nil {
					return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Error parsing data").Send(c)
				}

			}

		} else {
			return res.CustomErrorBuilder(http.StatusUnauthorized, nil, "Unauthorized").Send(c)
		}

		destructName := token.Claims.(jwt.MapClaims)["name"]
		if destructName != nil {
			if destructName, err = aescrypt.DecryptAES(destructName.(string), jwtKey); err != nil {
				return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
			}
			name = fmt.Sprintf("%v", destructName)
		} else {
			name = ""
		}

		destructEmail := token.Claims.(jwt.MapClaims)["email"]
		if destructEmail != nil {
			if destructEmail, err = aescrypt.DecryptAES(destructEmail.(string), jwtKey); err != nil {
				return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
			}
			email = fmt.Sprintf("%v", destructEmail)
		} else {
			email = ""
		}
		destructMobilePhone := token.Claims.(jwt.MapClaims)["mobilephone"]
		if destructMobilePhone != nil {
			if destructMobilePhone, err = aescrypt.DecryptAES(destructMobilePhone.(string), jwtKey); err != nil {
				return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
			}
			mobilephone = fmt.Sprintf("%v", destructMobilePhone)
		} else {
			mobilephone = ""
		}
		destructRole := token.Claims.(jwt.MapClaims)["role"]
		if destructRole != nil {
			if destructRole, err = aescrypt.DecryptAES(destructRole.(string), jwtKey); err != nil {
				return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
			}
			role = fmt.Sprintf("%v", destructRole)
		} else {
			role = ""
		}
		destructIsActive := token.Claims.(jwt.MapClaims)["isactive"]
		if destructIsActive != nil {
			if destructIsActive, err = aescrypt.DecryptAES(destructIsActive.(string), jwtKey); err != nil {
				return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
			}
			active := false
			if fmt.Sprintf("%v", destructIsActive) == "true" {
				active = true
			}
			isactive = active
		} else {
			isactive = false
		}
		destructIsLocked := token.Claims.(jwt.MapClaims)["islocked"]
		if destructIsLocked != nil {
			if destructIsLocked, err = aescrypt.DecryptAES(destructIsLocked.(string), jwtKey); err != nil {
				return res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error").Send(c)
			}
			locked := false
			if fmt.Sprintf("%v", destructIsLocked) == "true" {
				locked = true
			}
			islocked = locked
		} else {
			islocked = false
		}

		_, err = authService.GetUserByUserId(c.(*abstraction.Context), &dto.AuthMiddlewareRequest{ID: id})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err).Send(c)
		}

		cc := c.(*abstraction.Context)

		cc.Auth = &abstraction.AuthContext{
			ID:          id,
			IdentityId:  identityid,
			Name:        name,
			Email:       email,
			RoleId:      roleid,
			MobilePhone: mobilephone,
			Role:        role,
			IsActive:    isactive,
			IsLocked:    islocked,
		}

		return next(cc)
	}
}

func GenerateEncryptData(data *model.UserLoginModel) (*model.AccessTokenClaims, error) {
	var (
		result                                                                         model.AccessTokenClaims
		err                                                                            error
		userId, identityId, roleId, name, email, mobilePhone, role, isActive, isLocked string
	)
	if userId, err = aescrypt.EncryptAES(strconv.Itoa(data.Id), config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	if identityId, err = aescrypt.EncryptAES(strconv.Itoa(data.IdentityId), config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	if roleId, err = aescrypt.EncryptAES(strconv.Itoa(data.RoleId), config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	if name, err = aescrypt.EncryptAES(data.Name, config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	if email, err = aescrypt.EncryptAES(data.Email, config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	if mobilePhone, err = aescrypt.EncryptAES(data.MobilePhone, config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	if role, err = aescrypt.EncryptAES(data.Role, config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	if isActive, err = aescrypt.EncryptAES(strconv.FormatBool(data.IsActive), config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	if isLocked, err = aescrypt.EncryptAES(strconv.FormatBool(data.IsLocked), config.Get().Key.JwtKey); err != nil {
		return nil, res.CustomErrorBuilder(http.StatusInternalServerError, nil, "Internal Server Error")
	}
	result = model.AccessTokenClaims{
		Id:          userId,
		IdentityId:  identityId,
		RoleId:      roleId,
		Name:        name,
		Email:       email,
		MobilePhone: mobilePhone,
		Role:        role,
		IsActive:    isActive,
		IsLocked:    isLocked,
		Exp:         time.Now().Add(constant.ACCESS_TOKEN_DURATION).Unix(),
	}
	return &result, nil
}

func ValidateAccessToken(authToken string) (*abstraction.AuthContext, *int) {

	var (
		authContext                    abstraction.AuthContext
		jwtKey                         = config.Get().Key.JwtKey
		id, identityid, roleid         int
		name, email, mobilephone, role string
		isactive, islocked             bool
		unauthorized                   = 401
		ok                             = 200
	)

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil || !token.Valid {
		return nil, &unauthorized
	}

	destructID := token.Claims.(jwt.MapClaims)["id"]
	if destructID != nil {
		if destructID == "" {
			return nil, &unauthorized
		}
		id, err = strconv.Atoi(fmt.Sprintf("%v", destructID))
		if err != nil {
			if destructID, err = aescrypt.DecryptAES(destructID.(string), jwtKey); err != nil {
				return nil, &unauthorized
			}

			id, err = strconv.Atoi(fmt.Sprintf("%v", destructID))
			if err != nil {
				return nil, &unauthorized
			}
		}
	} else {
		return nil, &unauthorized
	}
	destructIdentityID := token.Claims.(jwt.MapClaims)["identityid"]
	if destructIdentityID != nil {
		if destructIdentityID == "" {
			return nil, &unauthorized
		}
		identityid, err = strconv.Atoi(fmt.Sprintf("%v", destructIdentityID))
		if err != nil {
			if destructIdentityID, err = aescrypt.DecryptAES(destructIdentityID.(string), jwtKey); err != nil {
				return nil, &unauthorized
			}

			identityid, err = strconv.Atoi(fmt.Sprintf("%v", destructIdentityID))
			if err != nil {
				return nil, &unauthorized
			}
		}
	} else {
		return nil, &unauthorized
	}
	destructRoleId := token.Claims.(jwt.MapClaims)["roleid"]
	if destructRoleId != nil {
		if destructRoleId == "" {
			return nil, &unauthorized
		}
		roleid, err = strconv.Atoi(fmt.Sprintf("%v", destructRoleId))
		if err != nil {

			if destructRoleId, err = aescrypt.DecryptAES(destructRoleId.(string), jwtKey); err != nil {
				return nil, &unauthorized
			}

			roleid, err = strconv.Atoi(fmt.Sprintf("%v", destructRoleId))
			if err != nil {
				return nil, &unauthorized
			}
		}
	} else {
		return nil, &unauthorized
	}
	destructName := token.Claims.(jwt.MapClaims)["name"]
	if destructName != nil {
		if destructName, err = aescrypt.DecryptAES(destructName.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		name = fmt.Sprintf("%v", destructName)
	} else {
		name = ""
	}
	destructEmail := token.Claims.(jwt.MapClaims)["email"]
	if destructEmail != nil {
		if destructEmail, err = aescrypt.DecryptAES(destructEmail.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		email = fmt.Sprintf("%v", destructEmail)
	} else {
		email = ""
	}
	destructMobilePhone := token.Claims.(jwt.MapClaims)["mobilephone"]
	if destructMobilePhone != nil {
		if destructMobilePhone, err = aescrypt.DecryptAES(destructMobilePhone.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		mobilephone = fmt.Sprintf("%v", destructMobilePhone)
	} else {
		mobilephone = ""
	}
	destructRole := token.Claims.(jwt.MapClaims)["role"]
	if destructRole != nil {
		if destructRole, err = aescrypt.DecryptAES(destructRole.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		role = fmt.Sprintf("%v", destructRole)
	} else {
		role = ""
	}
	destructIsActive := token.Claims.(jwt.MapClaims)["isactive"]
	if destructIsActive != nil {
		if destructIsActive, err = aescrypt.DecryptAES(destructIsActive.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		active := false
		if fmt.Sprintf("%v", destructIsActive) == "true" {
			active = true
		}
		isactive = active
	} else {
		isactive = false
	}
	destructIsLocked := token.Claims.(jwt.MapClaims)["islocked"]
	if destructIsLocked != nil {
		if destructIsLocked, err = aescrypt.DecryptAES(destructIsLocked.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		locked := false
		if fmt.Sprintf("%v", destructIsLocked) == "true" {
			locked = true
		}
		islocked = locked
	} else {
		islocked = false
	}

	authContext = abstraction.AuthContext{
		ID:          id,
		IdentityId:  identityid,
		Name:        name,
		Email:       email,
		MobilePhone: mobilephone,
		RoleId:      roleid,
		Role:        role,
		IsActive:    isactive,
		IsLocked:    islocked,
	}

	return &authContext, &ok
}

func ValidateRefreshToken(authToken model.Token) (*abstraction.AuthContext, *int) {

	var (
		authContext                    abstraction.AuthContext
		jwtKey                         = config.Get().Key.JwtKey
		id, identityid, roleid         int
		name, email, mobilephone, role string
		isactive, islocked             bool
		err                            error
		unauthorized                   = 401
		ok                             = 200
	)

	refToken, errParseRef := jwt.Parse(authToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if errParseRef != nil || !refToken.Valid {
		return nil, &unauthorized
	}

	accToken, _ := jwt.Parse(authToken.AccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	destructID := accToken.Claims.(jwt.MapClaims)["id"]
	if destructID != nil {
		if destructID == "" {
			return nil, &unauthorized
		}
		id, err = strconv.Atoi(fmt.Sprintf("%v", destructID))
		if err != nil {
			if destructID, err = aescrypt.DecryptAES(destructID.(string), jwtKey); err != nil {
				return nil, &unauthorized
			}

			id, err = strconv.Atoi(fmt.Sprintf("%v", destructID))
			if err != nil {
				return nil, &unauthorized
			}
		}
	} else {
		return nil, &unauthorized
	}
	destructIdentityID := accToken.Claims.(jwt.MapClaims)["identityid"]
	if destructIdentityID != nil {
		if destructIdentityID == "" {
			return nil, &unauthorized
		}
		identityid, err = strconv.Atoi(fmt.Sprintf("%v", destructIdentityID))
		if err != nil {
			if destructIdentityID, err = aescrypt.DecryptAES(destructIdentityID.(string), jwtKey); err != nil {
				return nil, &unauthorized
			}

			identityid, err = strconv.Atoi(fmt.Sprintf("%v", destructIdentityID))
			if err != nil {
				return nil, &unauthorized
			}
		}
	} else {
		return nil, &unauthorized
	}
	destructRoleId := accToken.Claims.(jwt.MapClaims)["roleid"]
	if destructRoleId != nil {
		if destructRoleId == "" {
			return nil, &unauthorized
		}
		roleid, err = strconv.Atoi(fmt.Sprintf("%v", destructRoleId))
		if err != nil {

			if destructRoleId, err = aescrypt.DecryptAES(destructRoleId.(string), jwtKey); err != nil {
				return nil, &unauthorized
			}

			roleid, err = strconv.Atoi(fmt.Sprintf("%v", destructRoleId))
			if err != nil {
				return nil, &unauthorized
			}
		}
	} else {
		return nil, &unauthorized
	}
	destructName := accToken.Claims.(jwt.MapClaims)["name"]
	if destructName != nil {
		if destructName, err = aescrypt.DecryptAES(destructName.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		name = fmt.Sprintf("%v", destructName)
	} else {
		name = ""
	}
	destructEmail := accToken.Claims.(jwt.MapClaims)["email"]
	if destructEmail != nil {
		if destructEmail, err = aescrypt.DecryptAES(destructEmail.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		email = fmt.Sprintf("%v", destructEmail)
	} else {
		email = ""
	}
	destructMobilePhone := accToken.Claims.(jwt.MapClaims)["mobilephone"]
	if destructMobilePhone != nil {
		if destructMobilePhone, err = aescrypt.DecryptAES(destructMobilePhone.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		mobilephone = fmt.Sprintf("%v", destructMobilePhone)
	} else {
		mobilephone = ""
	}
	destructRole := accToken.Claims.(jwt.MapClaims)["role"]
	if destructRole != nil {
		if destructRole, err = aescrypt.DecryptAES(destructRole.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		role = fmt.Sprintf("%v", destructRole)
	} else {
		role = ""
	}
	destructIsActive := accToken.Claims.(jwt.MapClaims)["isactive"]
	if destructIsActive != nil {
		if destructIsActive, err = aescrypt.DecryptAES(destructIsActive.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		active := false
		if fmt.Sprintf("%v", destructIsActive) == "true" {
			active = true
		}
		isactive = active
	} else {
		isactive = false
	}
	destructIsLocked := accToken.Claims.(jwt.MapClaims)["islocked"]
	if destructIsLocked != nil {
		if destructIsLocked, err = aescrypt.DecryptAES(destructIsLocked.(string), jwtKey); err != nil {
			return nil, &unauthorized
		}
		locked := false
		if fmt.Sprintf("%v", destructIsLocked) == "true" {
			locked = true
		}
		islocked = locked
	} else {
		islocked = false
	}

	authContext = abstraction.AuthContext{
		ID:          id,
		IdentityId:  identityid,
		Name:        name,
		Email:       email,
		MobilePhone: mobilephone,
		RoleId:      roleid,
		Role:        role,
		IsActive:    isactive,
		IsLocked:    islocked,
	}

	return &authContext, &ok
}

func ChangeTokenForLogout(ctx abstraction.Context) error {
	tokenString := ctx.Request().Header.Get("Authorization")
	if tokenString == "" {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil)
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, errParseToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().Key.JwtKey), nil
	})

	if errParseToken != nil || !token.Valid {
		return errParseToken
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Unix()
	token.Claims = claims
	token.Valid = false

	newTokenString, errSignedString := token.SignedString([]byte(config.Get().Key.JwtKey))
	if errSignedString != nil {
		return errSignedString
	}

	ctx.Response().Header().Set("Authorization", "Bearer "+newTokenString)

	return nil
}
