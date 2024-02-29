postgres:
	docker run --name some-postgres -e POSTGRES_USER=hit -e POSTGRES_PASSWORD=hitsara -p 5432:5432 -d postgres

createdb: 
	docker exec -it some-postgres createdb -U hit -O hit simple_bank

dropdb:
	docker exec -it some-postgres dropdb -U hit --password simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://hit:hitsara@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://hit:hitsara@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination=db/mock/store.go github.com/Hitesh-Sisara/bank-app-go/db/sqlc Store



.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
