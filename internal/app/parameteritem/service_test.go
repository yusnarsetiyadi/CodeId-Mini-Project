package parameteritem

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/repository"
	"compass_mini_api/pkg/test"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestGetAllDivision(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.ParameterItem
		Db         *gorm.DB
	}
	type args struct {
		ctx             *abstraction.Context
		queryPagination *abstraction.QueryPagination
		queryFilter     *abstraction.QueryFilter
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewParameterItem(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			queryPagination: &abstraction.QueryPagination{
				Limit:  "10",
				Offset: "0",
			},
			queryFilter: &abstraction.QueryFilter{
				Conditions: "%5B%7B%22column%22%3A%22name%22%2C%22value%22%3A%22Maintenance%22%2C%22comparation%22%3A%22%3D%22%7D%5D",
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
		got, err := s.GetAllDivision(tests.args.ctx, tests.args.queryPagination, tests.args.queryFilter)
		if (err != nil) != tests.wantErr {
			t.Logf("GetAllDivision() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetAllDivision() got = %v, want %v", got, tests.want)
		}
	})
}
