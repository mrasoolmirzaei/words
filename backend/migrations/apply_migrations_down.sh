docker cp ./backend/migrations/down/1.sql lendo-postgres:/down_1.sql
docker exec -u postgres lendo-postgres psql postgres postgres -f /down_1.sql