run-dev:
	@go run cmd/main.go

run-server:
	@py internal/grpc/server/server.py

gen-client:
	@protoc \
		--proto_path=api pdf.proto \
		--go_out=api/pb  --go_opt=paths=source_relative \
		--go-grpc_out=api/pb --go-grpc_opt=paths=source_relative \

gen-server:
	py -m grpc_tools.protoc \
		--proto_path=api \
		--python_out=api/pb \
		--grpc_python_out=api/pb \
		api/pdf.proto
