dev:
	go run cmd/main.go

migrate-up:
	migrate -path ./migrations -database 'postgres://root:secretpassword@localhost:5432/dbname?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://root:secretpassword@localhost:5432/dbname?sslmode=disable' down
