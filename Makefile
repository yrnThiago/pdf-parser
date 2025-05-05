run-dev:
	@go run cmd/main.go

gen:
	@protoc \
		--proto_path=internal/grpc/protobuf pdf.proto \
		--go_out=internal/grpc/services/genproto --go_opt=paths=source_relative \
		--go-grpc_out=internal/grpc/services/genproto --go-grpc_opt=paths=source_relative \
