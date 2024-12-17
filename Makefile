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

migrateup1:
	migrate -path Database/migration -database "postgresql://root:secret@localhost:5432/banking?sslmode=disable" -verbose up 1

migratedown:
	migrate -path Database/migration -database "postgresql://root:secret@localhost:5432/banking?sslmode=disable" -verbose down

migratedown1:
	migrate -path Database/migration -database "postgresql://root:secret@localhost:5432/banking?sslmode=disable" -verbose down 1

makesqlc:
	docker run --rm -v "C:\Users\Md. Rashid Aziz\Nawab\Banking:/src" -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	proto-buf-exe/bin/protoc.exe --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	./evans --host localhost --port 9090 -r repl

# remove all containers
# rmAllCont:
# 	docker rm -v -f $(docker ps -qa)

# remove all images 
# rmAllImg: 
# 	docker rmi $(docker images -a -q)

# mock: 
# 	docker run --rm `
#   	-v "${PWD}:/app" `
#   	-w /app `
#   	golang:1.23 bash -c "
#   	go install github.com/golang/mock/mockgen@v1.6.0 && \
#   	/go/bin/mockgen -destination=Database/mock/store.go github.com/rashid642/banking/Database/sqlc Store
#   	"

.PHONY: postgres stopPostgres createdb dropdb migrateup migratedown makesqlc test proto