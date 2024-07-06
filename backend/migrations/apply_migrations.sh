docker cp 1.sql lendo-postgres:/1.sql
docker exec -u postgres lendo-postgres psql postgres postgres -f /1.sql