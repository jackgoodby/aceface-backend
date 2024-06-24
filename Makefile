#COLORS
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

#up: clean build-dev start-and-seed sqlc-gen
up: clean build-dev start-and-seed
	#docker compose -f docker-compose.yml up -d aceface-db

clean:
	docker compose -f docker-compose.yml down -v --remove-orphans
	#docker compose run --rm yarn

build-dev:
	docker compose -f docker-compose.yml build aceface-db

start-and-seed:
	docker compose up -d --wait aceface-db
	docker compose run --rm --build migration
	docker compose exec aceface-db psql -U aceface -d postgres -a -f ./seed_data.sql

down:
	docker compose down

build:
	docker compose build --no-cache --parallel aceface-db

sqlc-gen:
	docker compose run --rm sqlc-gen

#sqlc-diff:
#	docker compose run --rm sqlc-diff