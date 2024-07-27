test_auth_service:
	go test ./internal/app/auth -coverprofile cover.out && go tool cover -html cover.out 

test_company_service:
	go test ./internal/app/company -coverprofile cover.out && go tool cover -html cover.out 

test_employee_service:
	go test ./internal/app/employee -coverprofile cover.out && go tool cover -html cover.out 

test_feature_service:
	go test ./internal/app/feature -coverprofile cover.out && go tool cover -html cover.out 

test_parameteritem_service:
	go test ./internal/app/parameteritem -coverprofile cover.out && go tool cover -html cover.out 

test_user_service:
	go test ./internal/app/user -coverprofile cover.out && go tool cover -html cover.out 

run_go:
	go run main.go