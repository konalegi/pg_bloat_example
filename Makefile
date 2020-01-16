setup:
	docker-compose build
	docker-compose run --rm app bin/db_init.sh

run:
	docker-compose run --rm app

open_db:
	docker-compose run --rm app psql -U postgres -d pg_bloat_db -h postgres

dump_db:
	docker-compose run --rm app pg_dump -U postgres -d pg_bloat_db -h postgres --schema-only > db.sql