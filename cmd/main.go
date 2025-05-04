package main

import (
	"github.com/yrnThiago/pdf-ocr/internal/server"
)

func main() {
	httpServer := server.NewHttpServer(":3000")
	go httpServer.Run()

	grpcServer := server.NewGRPCServer(":50051")
	grpcServer.Run()
}
