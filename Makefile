user-proto:
	protoc --go_out=. --go_opt=module=github.com/morf1lo/notification-system --go-grpc_out=. --go-grpc_opt=module=github.com/morf1lo/notification-system ./internal/user/proto/*.proto

user:
	go run cmd/user/main.go

feed:
	go run cmd/feed/main.go

worker:
	go run cmd/worker/main.go
