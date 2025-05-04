package main

import (
	"github.com/yrnThiago/pdf-ocr/config"
	grpcServer "github.com/yrnThiago/pdf-ocr/internal/grpc/server"
	httpServer "github.com/yrnThiago/pdf-ocr/internal/http/server"
)

func main() {
	config.Init()

	httpServer := httpServer.NewHttpServer(":3000")
	go httpServer.Run()

	grpcServer := grpcServer.NewGRPCServer(":50051")
	grpcServer.Run()
}
