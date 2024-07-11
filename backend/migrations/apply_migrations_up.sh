docker cp ./backend/migrations/up/1.sql lendo-postgres:/up_1.sql
docker exec -u postgres lendo-postgres psql postgres postgres -f /up_1.sql