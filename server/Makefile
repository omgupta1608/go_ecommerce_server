include .env

db-migrate-up:
	migrate -database ${DATABASE_URL}?sslmode=disable -path db/internal/migrations up