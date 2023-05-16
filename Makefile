DB_URI = "postgresql://root:secretpassword@localhost:5432/dbname?sslmode=disable"

migrate-up:
	migrate -path migrations/ -database ${DB_URI} -verbose up

migrate-down:
	migrate -path migrations/ -database ${DB_URI} -verbose down

migrate-fix:
	migrate -path migrations/ -database ${DB_URI} force VERSION
