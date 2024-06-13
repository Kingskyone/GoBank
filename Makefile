DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name postgres1 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres1 createdb --user=root --owner=root simple_bank

dropdb:
	docker exec -it postgres1 dropdb simple_bank

migratecreate:
	migrate create -ext sql -dir db/migration -seq 名称

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "$postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	./sqlc.exe generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	# 删除原本生成的文件
	# rm -f pb/*.go
	# rm -f doc/swagger/*.swagger.json
	.\protoc.exe --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=go_bank proto/*.proto
	# 使用statik将静态文件进行打包 生成go文件
	statik -src=doc/swagger -dest=doc

evans:
	.\evans.exe --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migratecreate migrateup migrateup1 migratedown migratedown1 sqlc test server proto evans