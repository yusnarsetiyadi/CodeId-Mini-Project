package user

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/repository"
	"compass_mini_api/pkg/test"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestChangePassword(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.User
		Db         *gorm.DB
	}
	type args struct {
		ctx     *abstraction.Context
		payload *dto.UserChangePasswordRequest
		paramId *dto.UserChangePasswordRequestParam
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
			payload: &dto.UserChangePasswordRequest{
				OldPassword: "Test12345@",
				NewPassword: "Test1234@",
			},
			paramId: &dto.UserChangePasswordRequestParam{
				Id: 3,
			},
		},
		want:    nil,
		wantErr: false,
	}
	t.Run("failed ParamId", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.ChangePassword(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("ChangePassword() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("ChangePassword() got = %v, want %v", got, tests.want)
		}
	})
	tests.args.ctx.Auth.ID = 0
	t.Run("failed FindById", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.ChangePassword(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("ChangePassword() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("ChangePassword() got = %v, want %v", got, tests.want)
		}
	})
}
