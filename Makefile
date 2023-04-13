
GO_MODULE := github.com/HCMUT-UWC-2-0/backend
POSTGRES_USER := root
POSTGRES_PASSWORD := secret
POSTGRES_DB := uwcdb
POSTGRES_PORT := 5432
POSTGRES_HOST := localhost
POSTGRES_SOURCE := postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable
# Run this command once when init project
bootstrap:
	docker-compose -f docker-compose.dev.yml up -d --build --remove-orphans;
# For development
downstrap:
	docker-compose -f docker-compose.dev.yml down --remove-orphans --volumes

clean_dev:
	docker-compose -f docker-compose.dev.yml down --volumes --remove-orphans --rmi local

update:
	make clean_dev
	make bootstrap

# For db
	
createdb:
	docker exec -it postgres createdb --username=${POSTGRES_USER} --owner={POSTGRES_USER} ${POSTGRES_DB}
dropdb:
	docker exec -it postgres dropdb ${POSTGRES_DB}

migrateup:
	migrate -path db/migration -database "${POSTGRES_SOURCE}" -verbose up

migrateup1:
	migrate -path db/migration -database "${POSTGRES_SOURCE}" -verbose up 1

migratedown:
	migrate -path db/migration -database "${POSTGRES_SOURCE}" -verbose down

# Followings commands 're used for making seed data

copy_sql:
	cat db/initdb/*.sql >> db/initdb/init.sql
remove_sql:
	rm db/initdb/init.sql
init_sql:
	docker exec -it postgres psql -U root ${POSTGRES_DB} -a -f ${POSTGRES_DB}/initdb/init.sql
seed:
	go run main.go seed-backofficers

# This command is used for Postgres interaction

psql:
	docker exec -it postgres psql -U root -d ${POSTGRES_DB} 

test:
	go test -v -cover ./...

# Run server in development
server:
	go run main.go 

sqlc:
	sqlc generate

# mock database
mock:
	mockgen -package mockdb -destination db/mock/store.go ${GO_MODULE}/db/sqlc Store

deploy: 
	docker-compose -f docker-compose.yml up -d --build --remove-orphans



.PHONY: bootstrap downstrap clean_dev update createdb dropdb migrateup migratedown copy_sql remove_sql init_sql seed psql test server sqlc mock
