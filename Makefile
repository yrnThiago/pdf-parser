run-dev:
	@go run cmd/main.go

gen:
	@protoc \
		--proto_path=api pdf.proto \
		--go_out=api/pb  --go_opt=paths=source_relative \
		--go-grpc_out=api/pb --go-grpc_opt=paths=source_relative \
