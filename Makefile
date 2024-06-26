#COLORS
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

up: clean build-dev init-param-store init-db-seed sqlc-gen

clean:
	docker compose -f docker-compose.yml down -v --remove-orphans
	#docker compose run --rm yarn

build-dev:
	#docker compose -f docker-compose.yml build --no-cache aceface-db
	docker compose -f docker-compose.yml build aceface-db
	docker compose -f docker-compose.yml build localstack

init-param-store:
	docker compose up -d --wait localstack
	aws ssm put-parameter --name "LAMBDA_ENV" --value "local" --type "String" --profile localstack --overwrite > /dev/null

init-db-seed:
	docker compose up -d --wait aceface-db
	docker compose run --rm --build migration
	docker compose exec aceface-db psql -U aceface -d postgres -a -f ./seed_data.sql

down:
	docker compose down

sqlc-gen:
	docker compose run --rm sqlc-gen

#sqlc-diff:
#	docker compose run --rm sqlc-diff

lambda-check:
	rm -Rf .aws-sam/build/GetMemberFunction/bootstrap
	rm -Rf .aws-sam/build/GetMembersFunction/bootstrap
	rm -Rf .aws-sam/cache/*
	cd src/lambda/getMember && GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ../../../.aws-sam/build/GetMemberFunction/bootstrap .
	cd src/lambda/getMembers && GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ../../../.aws-sam/build/GetMembersFunction/bootstrap .

sam-build:
	rm -Rf .aws-sam/cache/*
	sam build

sam-deploy: sam-build
	aws ssm put-parameter --name "LAMBDA_ENV" --value "aws" --type "String" --overwrite > /dev/null
	rm -Rf .aws-sam/build/template.yaml
	sam deploy --guided

sam-invoke:
	sam local invoke GetMembersFunction

sam-run-api-server:
	sam local start-api --debug