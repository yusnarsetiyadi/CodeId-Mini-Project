package employee

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	"compass_mini_api/internal/factory"
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

func TestGetAllEmployee(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.Employee
		Db         *gorm.DB
	}
	type args struct {
		ctx             *abstraction.Context
		queryPagination *abstraction.QueryPagination
		queryFilter     *abstraction.QueryFilter
		queryOrder      *abstraction.QueryOrder
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewEmployee(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			queryPagination: &abstraction.QueryPagination{
				Limit:  "10",
				Offset: "0",
			},
			queryFilter: &abstraction.QueryFilter{
				Conditions: "%5B%7B%22column%22%3A%22name%22%2C%22value%22%3A%22yusnar%22%2C%22comparation%22%3A%22%25%22%7D%2C%7B%22column%22%3A%22supervisorid%22%2C%22value%22%3A%2211%22%2C%22comparation%22%3A%22%3D%22%7D%2C%7B%22column%22%3A%22joindate%22%2C%22value%22%3A%222024-01-24_2024-01-28%22%2C%22comparation%22%3A%22date%22%7D%5D",
			},
			queryOrder: &abstraction.QueryOrder{
				Order:     "name",
				Direction: "asc",
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
		got, err := s.GetAllEmployee(tests.args.ctx, tests.args.queryPagination, tests.args.queryOrder, tests.args.queryFilter)
		if (err != nil) != tests.wantErr {
			t.Logf("GetAllEmployee() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetAllEmployee() got = %v, want %v", got, tests.want)
		}
	})

	tests.args.queryPagination = nil
	t.Run("test failed", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.GetAllEmployee(tests.args.ctx, tests.args.queryPagination, tests.args.queryOrder, tests.args.queryFilter)
		if (err != nil) != tests.wantErr {
			t.Logf("GetAllEmployee() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetAllEmployee() got = %v, want %v", got, tests.want)
		}
	})
}

func TestGetAllEmployeeSupervisor(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.Employee
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
			Repository: repository.NewEmployee(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			queryPagination: &abstraction.QueryPagination{
				Limit:  "10",
				Offset: "0",
			},
			queryFilter: &abstraction.QueryFilter{
				Conditions: "%5B%7B%22column%22%3A%22name%22%2C%22value%22%3A%22herru%22%2C%22comparation%22%3A%22%25%22%7D%5D",
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
		got, err := s.GetAllEmployeeSupervisor(tests.args.ctx, tests.args.queryPagination, tests.args.queryFilter)
		if (err != nil) != tests.wantErr {
			t.Logf("GetAllEmployeeSupervisor() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetAllEmployeeSupervisor() got = %v, want %v", got, tests.want)
		}
	})

	tests.args.queryPagination = nil
	t.Run("test failed", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.GetAllEmployeeSupervisor(tests.args.ctx, tests.args.queryPagination, tests.args.queryFilter)
		if (err != nil) != tests.wantErr {
			t.Logf("GetAllEmployeeSupervisor() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetAllEmployeeSupervisor() got = %v, want %v", got, tests.want)
		}
	})
}

func TestCreate(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.Employee
		Db         *gorm.DB
	}
	type args struct {
		ctx     *abstraction.Context
		payload *dto.CreateEmployeeRequest
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewEmployee(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			payload: &dto.CreateEmployeeRequest{
				Name:         "Yusnar test",
				Email:        "yusnar@code.id",
				PhoneNumber:  "+6281234567813",
				CompanyId:    43,
				Company:      "CODE.ID",
				DivisionId:   16,
				Division:     "Maintenance",
				SupervisorId: 11,
				Supervisor:   "Herru Purnomo Santoso",
				JoinDate:     "2024-01-24",
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
		got, err := s.Create(tests.args.ctx, tests.args.payload)
		if (err != nil) != tests.wantErr {
			t.Logf("Create() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Create() got = %v, want %v", got, tests.want)
		}
	})

	tests.args.payload.JoinDate = "2021"
	t.Run("test error ParseDate", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Create(tests.args.ctx, tests.args.payload)
		if (err != nil) != tests.wantErr {
			t.Logf("Create() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Create() got = %v, want %v", got, tests.want)
		}
	})

	tests.args.payload.JoinDate = "2024-01-24"
	tests.args.payload.SupervisorId = 9999
	t.Run("test error FindByIdEmployee", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Create(tests.args.ctx, tests.args.payload)
		if (err != nil) != tests.wantErr {
			t.Logf("Create() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Create() got = %v, want %v", got, tests.want)
		}
	})
}

func TestGetEmployeeById(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.Employee
		Db         *gorm.DB
	}
	type args struct {
		ctx     *abstraction.Context
		payload *dto.GetEmployeeByIdRequestParam
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewEmployee(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			payload: &dto.GetEmployeeByIdRequestParam{
				Id: 1,
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
		got, err := s.GetEmployeeById(tests.args.ctx, tests.args.payload)
		if (err != nil) != tests.wantErr {
			t.Logf("GetEmployeeById() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetEmployeeById() got = %v, want %v", got, tests.want)
		}
	})
	tests.args.payload.Id = 0
	t.Run("test error FindByIdEmployee", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.GetEmployeeById(tests.args.ctx, tests.args.payload)
		if (err != nil) != tests.wantErr {
			t.Logf("GetEmployeeById() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("GetEmployeeById() got = %v, want %v", got, tests.want)
		}
	})
}

func TestUpdate(t *testing.T) {
	ctx, conn := test.Init(t)
	type fields struct {
		Repository repository.Employee
		Db         *gorm.DB
	}
	type args struct {
		ctx     *abstraction.Context
		payload *dto.UpdateEmployeeRequest
		paramId *dto.UpdateEmployeeRequestParam
	}
	tests := struct {
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		fields: fields{
			Repository: repository.NewEmployee(conn),
			Db:         conn,
		},
		args: args{
			ctx: ctx,
			payload: &dto.UpdateEmployeeRequest{
				Name:         nil,
				Email:        nil,
				PhoneNumber:  nil,
				CompanyId:    nil,
				Company:      nil,
				DivisionId:   nil,
				Division:     nil,
				SupervisorId: nil,
				Supervisor:   nil,
				JoinDate:     nil,
			},
			paramId: &dto.UpdateEmployeeRequestParam{
				Id: 1,
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
		got, err := s.Update(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("Update() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Update() got = %v, want %v", got, tests.want)
		}
	})

	joindate := "2024-01-28"
	tests.args.payload.JoinDate = &joindate
	t.Run("success JoinDateNotNil", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Update(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("Update() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Update() got = %v, want %v", got, tests.want)
		}
	})
	joindateFalse := "2024"
	tests.args.payload.JoinDate = &joindateFalse
	t.Run("failed ParseJoinDate", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Update(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("Update() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Update() got = %v, want %v", got, tests.want)
		}
	})
	resigndate := "2024-01-28"
	tests.args.payload.JoinDate = &joindate
	tests.args.payload.ResignDate = &resigndate
	t.Run("success ResignDateNotNil", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Update(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("Update() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Update() got = %v, want %v", got, tests.want)
		}
	})
	resigndateFalse := "2024"
	tests.args.payload.ResignDate = &resigndateFalse
	t.Run("failed ParseResignDate", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Update(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("Update() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Update() got = %v, want %v", got, tests.want)
		}
	})
	tests.args.payload.ResignDate = &resigndate
	idSpv := 11
	spv := "Herru Purnomo Santoso"
	tests.args.payload.SupervisorId = &idSpv
	tests.args.payload.Supervisor = &spv
	t.Run("failed SpvNotNil", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Update(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("Update() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Update() got = %v, want %v", got, tests.want)
		}
	})
	idSpvGet := 0
	tests.args.payload.SupervisorId = &idSpvGet
	t.Run("failed FindByIdEmployee", func(t *testing.T) {
		s := service{
			Repository: tests.fields.Repository,
			Db:         tests.fields.Db,
		}
		got, err := s.Update(tests.args.ctx, tests.args.payload, tests.args.paramId)
		if (err != nil) != tests.wantErr {
			t.Logf("Update() error = %v, wantErr %v", err, tests.wantErr)
		}
		if !reflect.DeepEqual(got, tests.want) {
			t.Logf("Update() got = %v, want %v", got, tests.want)
		}
	})
}
