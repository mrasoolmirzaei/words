testing:
	cd backend && go test ./... -v

run-backend:
	cd backend/cmd/words-api && go run . server

run-front:
	cd frontend/words && npm install && npm start

migration:
	./backend/migrations/apply_migrations_up.sh

migrate-down:
	./backend/migrations/apply_migrations_down.sh
	