package auth

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/model"
	"compass_mini_api/internal/repository"
	"compass_mini_api/pkg/test"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewService(t *testing.T) {
	// init
	test.Init(t)
	// prepare args
	type args struct {
		f *factory.Factory
	}
	tests := struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		args: args{
			f: factory.NewFactory(),
		},
		want:    nil,
		wantErr: false,
	}
	got := NewService(tests.args.f)
	if !reflect.DeepEqual(got, tests.want) {
		t.Logf("NewService() got = %v, want %v", got, tests.want)
	}
}

func TestLogin(t *testing.T) {
	// init
	ctx, conn := test.Init(t)
	// prepare args
	type fields struct {
		Repository repository.User
		Db         *gorm.DB
	}
	type args struct {
		ctx         *abstraction.Context
		payload     *dto.AuthLoginRequest
		queryEntity *abstraction.QueryEntity
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewUser(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			payload: &dto.AuthLoginRequest{
				MobilePhone: "+6281234567890",
				Password:    "Test12345@",
				Geolocation: dto.Geolocation{
					IP:          "111.94.121.97",
					Asn:         "AS23700 Linknet-Fastnet ASN",
					Netmask:     16,
					Hostname:    "fm-dyn-111-94-121-97.fast.net.id.",
					City:        "Jakarta",
					PostCode:    "15710",
					Country:     "Indonesia",
					CountryCode: "ID",
					Latitude:    -6.1743998527526855,
					Longitude:   106.82939910888672,
				},
			},
			queryEntity: &abstraction.QueryEntity{
				Entity: "Android",
			},
		},
		want:    nil,
		wantErr: false,
	}
	t.Run("test success", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Login(tests.args.ctx, tests.args.payload, tests.args.queryEntity)
		if (err != nil) != tests.wantErr {
			t.Logf("Login() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Login() got = %v, want %v", got, tests.want)
		}
	})

	tests.args.payload.MobilePhone = "+62812345678999"
	t.Run("test error FindByPhoneQuery", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Login(tests.args.ctx, tests.args.payload, tests.args.queryEntity)
		if (err != nil) != tests.wantErr {
			t.Logf("Login() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Login() got = %v, want %v", got, tests.want)
		}
	})

	tests.args.payload.MobilePhone = "+6281234567890"
	tests.args.payload.Password = "wrongpassword"
	t.Run("test error CompareHashAndPassword", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Login(tests.args.ctx, tests.args.payload, tests.args.queryEntity)
		if (err != nil) != tests.wantErr {
			t.Logf("Login() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Login() got = %v, want %v", got, tests.want)
		}
	})

	tests.args.payload.MobilePhone = "+6281234567890"
	tests.args.payload.Password = "Test12345@"
	tests.args.queryEntity.Entity = "AndroidWrong"
	t.Run("test error len(*dataAuth)", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Login(tests.args.ctx, tests.args.payload, tests.args.queryEntity)
		if (err != nil) != tests.wantErr {
			t.Logf("Login() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Login() got = %v, want %v", got, tests.want)
		}
	})
}

func TestSplash(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.User
		Db         *gorm.DB
	}
	type args struct {
		ctx     *abstraction.Context
		payload *dto.AuthSplashRequest
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewUser(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			payload: &dto.AuthSplashRequest{
				Token: model.Token{
					AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImE0MzNlNjY0Yzc5MDQ3YzE3NTYwN2UwOTQ0YTdkNzk4NGMzMDEzMDVmMGQ4Nzc3ZTI5NTAyZGVhNmMiLCJpZGVudGl0eWlkIjoiMTkzNWIxNzZlMGMxY2ViYWQwMmIxY2IxNmQwODJlODEzZDQyMjVhMTAwZTA3OGY1OWRkMTA5ZDJlMCIsIm5hbWUiOiI3Y2FlOWZjZGU1MTZhYzJiOGZkNTYzYmQ4YjlkMjM1MDY1OGYyNDZmZjczZDU0YzRjMzAwZTI1M2Q3ZjIzN2I5Y2QzYiIsImVtYWlsIjoiNzc2ZTFiNjM0ZTliYmNhYWZkM2FmNGM1NjZjYTM4NDRjNjRiZDAzNWRiOWFhNDBiNTA3NmU2MGEyN2E0ZjhlNzJhZjgwOWNhODBkZmI4OGMxNTQyIiwibW9iaWxlcGhvbmUiOiI3Yjc0YTAzMjBjNWFkODg1NmEwMTZiMTJhODFkMWFlYzdjZDY5Y2FhNjU3ODM3ZDRiY2U0NDNjMjg5NmE1NDkyOGNhYjNhNGFiMjA5NzNkYTliNjgiLCJyb2xlaWQiOiIzNGY4YjQ1ZGJiMzYwNDY3YmIwOWE3YmIyYzIxMzBlZDE3ZGVlMjRhMzZkNTcwZjEwNzY4ODhmMDgxIiwicm9sZSI6Ijc0NjRlMTdmMGFkNDk0NjcyNDIzOTZiNmJiZWU2NzNlMjU3OTU4MzFjNDI2ZjU3OGRlOGFjMzk5ZmNlZTUwNzRiZDJmYTc1NGFmNzZhNzc3OTciLCJpc2FjdGl2ZSI6IjIyZjAzOWZmZDYxN2MyMjRhYWYxZTk0N2VhZTA5YWJiZTIyZTdkZWMxZDEyOGZlNzYwNjhmOWMyMGU0NTE5Y2MiLCJpc2xvY2tlZCI6ImE4ZTc2ZWQ0MWExY2FjMTAwZGNhMGE4OWZkNmY4ZTE5M2UyODcwYjE2ODk4OGU2ODY2NGNhNDI5MGQ3NWZkYWQwMSIsImV4cCI6MTcwODQwMDMxOX0.ZqTS51qxdx1AxsRWDnZ4Wns4Sp4AUBnCLcVrLjxAHjA",
					RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzODc1MTl9.-p5hO1xIXN3RGV7rI5qAwZRbcHLIVWR3mm41uCVeKbo",
				},
			},
		},
		want:    nil,
		wantErr: false,
	}
	t.Run("test success", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Splash(tests.args.ctx, tests.args.payload)
		if (err != nil) != tests.wantErr {
			t.Logf("Splash() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Splash() got = %v, want %v", got, tests.want)
		}
	})
	tests.args.payload.AccessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjdjNGQzMzkyMjFhYmZjNzM5ZDU0MzI0OTJhMTY1ODc2ODVlNmY2ZGNhZDUxMGQyNzY0MmExNzA3OTIiLCJpZGVudGl0eWlkIjoiOGMzNmVkMmVmZjJjMjVjOWI3NmQ3MjRlMWVmYWYxNjE0Yzk1N2U1YTgxZGRlMDRmMWQ5MTJmMTE4NCIsIm5hbWUiOiI4NjM3MWYwNmFmNjkyOGNlNzhkMWQ1ODMwNTQ3OTFiNDk1NjI4NzI5MjQ0MjQ0YjE2ZWNlNzY0ZDZmNzlhNjA5MGI4ZCIsImVtYWlsIjoiYjdmMDljYmI2ZDUzYmY5NGVhNzBkZGI5OTI1YzY0MjllOTkxNWIxY2U0ZDZiZjIxNTE5NjY5OTdhNWFiZTI2ZWY4OGJlOGE1NGU3M2IxMzhlODFkIiwibW9iaWxlcGhvbmUiOiI4MWRiOTA2Zjg2ZWI1NTkzYWI4OTU2NWMwNzFmYmMyNjVkZWE1Zjg2NzkyZThiZTEyNDA0OThiYTk4OWZhNDRmMWQ4ZWE2Yjg0OTFhODg5Y2FhMDIiLCJyb2xlaWQiOiJiMjgxZDFmMTZkMjM3MmIyNTdlNmFhNWIzMmZiMzJkOGRjNDAwMzg5NmE4MWYxYzAwODYwNWQ2ZmMzIiwicm9sZSI6IjhjY2M1MGUzZDI5OGU5Njk3MTYyODg2ZWYyMTU1NTM5ZGEwZDk5MzU1ZTMxNTdlODk5YmRkN2ZmZDVkNGZjM2I0MWZkZDU2YTc0MzQyOThlYTQiLCJpc2FjdGl2ZSI6IjgzZGM5YzE0Y2JiOWNkYmY3ZTUzYWVkOWIxZmQxOTFjMDlkMzZmZGFjZjQ0YTc3YzdmOWQwZDI3ZDVlOGFlZTkiLCJpc2xvY2tlZCI6ImU3ZmVhOGE4MjQ5MmFlZTU4MjVkOWExZDc0ZDNkNTE4NWZiZjRkYTgwMjA5NTE0YzJkODU1Y2VlYTdhMDg0ZmExMiIsImV4cCI6MTcwNDQxMzc4N30.lB4TX8DDht_FRolvp79vq5O8XXSqQWVNtmPWsLGvj1I"
	t.Run("failed ValidateAccessToken", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Splash(tests.args.ctx, tests.args.payload)
		if (err != nil) != tests.wantErr {
			t.Logf("Splash() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Splash() got = %v, want %v", got, tests.want)
		}
	})
	tests.args.payload.AccessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjdjNGQzMzkyMjFhYmZjNzM5ZDU0MzI0OTJhMTY1ODc2ODVlNmY2ZGNhZDUxMGQyNzY0MmExNzA3OTIiLCJpZGVudGl0eWlkIjoiOGMzNmVkMmVmZjJjMjVjOWI3NmQ3MjRlMWVmYWYxNjE0Yzk1N2U1YTgxZGRlMDRmMWQ5MTJmMTE4NCIsIm5hbWUiOiI4NjM3MWYwNmFmNjkyOGNlNzhkMWQ1ODMwNTQ3OTFiNDk1NjI4NzI5MjQ0MjQ0YjE2ZWNlNzY0ZDZmNzlhNjA5MGI4ZCIsImVtYWlsIjoiYjdmMDljYmI2ZDUzYmY5NGVhNzBkZGI5OTI1YzY0MjllOTkxNWIxY2U0ZDZiZjIxNTE5NjY5OTdhNWFiZTI2ZWY4OGJlOGE1NGU3M2IxMzhlODFkIiwibW9iaWxlcGhvbmUiOiI4MWRiOTA2Zjg2ZWI1NTkzYWI4OTU2NWMwNzFmYmMyNjVkZWE1Zjg2NzkyZThiZTEyNDA0OThiYTk4OWZhNDRmMWQ4ZWE2Yjg0OTFhODg5Y2FhMDIiLCJyb2xlaWQiOiJiMjgxZDFmMTZkMjM3MmIyNTdlNmFhNWIzMmZiMzJkOGRjNDAwMzg5NmE4MWYxYzAwODYwNWQ2ZmMzIiwicm9sZSI6IjhjY2M1MGUzZDI5OGU5Njk3MTYyODg2ZWYyMTU1NTM5ZGEwZDk5MzU1ZTMxNTdlODk5YmRkN2ZmZDVkNGZjM2I0MWZkZDU2YTc0MzQyOThlYTQiLCJpc2FjdGl2ZSI6IjgzZGM5YzE0Y2JiOWNkYmY3ZTUzYWVkOWIxZmQxOTFjMDlkMzZmZGFjZjQ0YTc3YzdmOWQwZDI3ZDVlOGFlZTkiLCJpc2xvY2tlZCI6ImU3ZmVhOGE4MjQ5MmFlZTU4MjVkOWExZDc0ZDNkNTE4NWZiZjRkYTgwMjA5NTE0YzJkODU1Y2VlYTdhMDg0ZmExMiIsImV4cCI6MTcwNDQxMzc4N30.lB4TX8DDht_FRolvp79vq5O8XXSqQWVNtmPWsLGvj1I"
	tests.args.payload.RefreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ0MTM3ODd9.KAHL6AP_WAdIN-7LzpKbYGCvppqHp676sfS3zAexyvc"
	t.Run("failed ValidateRefreshToken", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Splash(tests.args.ctx, tests.args.payload)
		if (err != nil) != tests.wantErr {
			t.Logf("Splash() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Splash() got = %v, want %v", got, tests.want)
		}
	})
}

func TestLogout(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.User
		Db         *gorm.DB
	}
	type args struct {
		ctx *abstraction.Context
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewUser(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
		},
		want:    nil,
		wantErr: false,
	}
	t.Run("test success", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Logout(tests.args.ctx)
		if (err != nil) != tests.wantErr {
			t.Logf("Logout() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Logout() got = %v, want %v", got, tests.want)
		}
	})
	tests.args.ctx = &abstraction.Context{
		Auth: &abstraction.AuthContext{
			ID: 0,
		},
	}
	t.Run("test error Update", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Logout(tests.args.ctx)
		if (err != nil) != tests.wantErr {
			t.Logf("Logout() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Logout() got = %v, want %v", got, tests.want)
		}
	})
}

func TestGetDataToken(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.User
		Db         *gorm.DB
	}
	type args struct {
		ctx *abstraction.Context
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewUser(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
		},
		want:    nil,
		wantErr: false,
	}
	tests.args.ctx = &abstraction.Context{
		Auth: &abstraction.AuthContext{
			ID:          1,
			IdentityId:  1,
			RoleId:      1,
			Name:        "yusnar",
			MobilePhone: "0838sisanyakapankapan",
			Email:       "yusnar@code.id",
			Role:        "maintenance",
			IsActive:    true,
			IsLocked:    false,
		},
	}
	t.Run("test success", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.GetDataToken(tests.args.ctx)
		if (err != nil) != tests.wantErr {
			t.Logf("GetDataToken() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetDataToken() got = %v, want %v", got, tests.want)
		}
	})
}
