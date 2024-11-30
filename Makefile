postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

stopPostgres:
	docker stop postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root banking 

dropdb: 
	docker exec -it postgres dropdb banking 

migrateup:
	migrate -path Database/migration -database "postgresql://root:secret@localhost:5432/banking?sslmode=disable" -verbose up

migratedown:
	migrate -path Database/migration -database "postgresql://root:secret@localhost:5432/banking?sslmode=disable" -verbose down

makesqlc:
	docker run --rm -v "C:\Users\Md. Rashid Aziz\Nawab\Learnings\Go Learning\Udemy\Banking:/src" -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres stopPostgres createdb dropdb migrateup migratedown makesqlc test