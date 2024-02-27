postgres:
	docker run --name some-postgres -e POSTGRES_USER=hit -e POSTGRES_PASSWORD=hitsara -p 5432:5432 -d postgres

createdb: 
	docker exec -it some-postgres createdb -U hit -O hit simple_bank

dropdb:
	docker exec -it some-postgres dropdb -U hit --password simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://hit:hitsara@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://hit:hitsara@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate


.PHONY: postgres createdb dropdb migrateup migratedown
