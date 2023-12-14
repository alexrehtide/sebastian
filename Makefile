.DEFAULT_GOAL := dev

dev:
	air -c .air.toml

client:
	go run cmd/client/client.go

gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/v1/sebastian.proto
