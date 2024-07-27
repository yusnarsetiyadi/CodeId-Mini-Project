package feature

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/repository"
	"compass_mini_api/pkg/test"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestGetFeatureList(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.Feature
		Db         *gorm.DB
	}
	type args struct {
		ctx         *abstraction.Context
		queryEntity *abstraction.QueryEntity
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewFeature(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			queryEntity: &abstraction.QueryEntity{
				Entity: "iOS",
			},
		},
		want:    nil,
		wantErr: false,
	}
	t.Run("success", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.GetFeatureList(tests.args.ctx, tests.args.queryEntity)
		if (err != nil) != tests.wantErr {
			t.Logf("GetFeatureList() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetFeatureList() got = %v, want %v", got, tests.want)
		}
	})
}

func TestGetFeatureSub(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.Feature
		Db         *gorm.DB
	}
	type args struct {
		ctx         *abstraction.Context
		payload     *dto.GetFeatureSubRequestParam
		queryEntity *abstraction.QueryEntity
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewFeature(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			payload: &dto.GetFeatureSubRequestParam{
				Id: 1,
			},
			queryEntity: &abstraction.QueryEntity{
				Entity: "iOS",
			},
		},
		want:    nil,
		wantErr: false,
	}
	t.Run("success", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.GetFeatureSub(tests.args.ctx, tests.args.payload, tests.args.queryEntity)
		if (err != nil) != tests.wantErr {
			t.Logf("GetFeatureSub() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetFeatureSub() got = %v, want %v", got, tests.want)
		}
	})
}
