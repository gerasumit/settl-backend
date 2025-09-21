postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:latest

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root settl

dropdb:
	docker exec -it postgres12 dropdb settl
	
migrateup:
	migrate -path db/migrations -database "postgresql://root:admin@localhost:5432/settl?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:admin@localhost:5432/settl?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup