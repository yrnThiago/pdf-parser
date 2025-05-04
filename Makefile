gen:
	@protoc \
		--proto_path=internal/protobuf pdf.proto \
		--go_out=internal/services/genproto --go_opt=paths=source_relative \
		--go-grpc_out=internal/services/genproto --go-grpc_opt=paths=source_relative \
