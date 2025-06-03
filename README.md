migrate create -ext sql -dir db/migrations create_users_table

migrate -path db/migrations/ -database 'postgres://alam@localhost:5432/go-absen?sslmode=disable' up
migrate -path db/migrations/ -database 'postgres://alam@localhost:5432/go-absen?sslmode=disable' down
