package main

import (
	"github.com/yrnThiago/pdf-ocr/config"
	"github.com/yrnThiago/pdf-ocr/internal/server"
)

func main() {
	config.Init()

	httpServer := server.NewHttpServer(":3000")
	go httpServer.Run()

	grpcServer := server.NewGRPCServer(":50051")
	grpcServer.Run()
}
