postgres:
	docker run --name postgresxm -p 5432:5432 -e POSTGRES_DB=companies -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root companies

dropdb:
	docker exec -it postgres14 dropdb companies

migrateup:
	migrate -path backend/db/migration -database "postgresql://root:secret@localhost:5432/companies?sslmode=disable" -verbose up

migratedown:
	migrate -path backend/db/migration -database "postgresql://root:secret@localhost:5432/companies?sslmode=disable" -verbose down

rundocker:
	docker-compose up -d

test:
	go test -v -cover ./...

mock:
	mockgen -package mock --destination service/mocks/mock.go github.com/tedoham/xm-test/internal/service CompanyService 

.PHONY: postgres createdb dropdb migrateup migratedown test mock rundocker