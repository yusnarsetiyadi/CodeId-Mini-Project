package factory

import (
	"compass_mini_api/internal/repository"
	"compass_mini_api/pkg/database"

	"gorm.io/gorm"
)

type Factory struct {
	Db *gorm.DB

	// repository
	repository_initiated
}

type repository_initiated struct {
	UserRepository          repository.User
	FeatureRepository       repository.Feature
	ParameterItemRepository repository.ParameterItem
	CompanyRepository       repository.Company
	EmployeeRepository      repository.Employee
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("POSTGRES")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	// auth
	f.FeatureRepository = repository.NewFeature(f.Db)

	// user
	f.UserRepository = repository.NewUser(f.Db)

	// employee
	f.EmployeeRepository = repository.NewEmployee(f.Db)

	// parameteritem
	f.ParameterItemRepository = repository.NewParameterItem(f.Db)

	// company
	f.CompanyRepository = repository.NewCompany(f.Db)
}
