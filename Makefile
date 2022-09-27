proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/**/pb/*.proto
server:
	go run cmd/main.go

vet:
	go vet cmd/main.go
	shadow cmd/main.go
.PHONY:vet