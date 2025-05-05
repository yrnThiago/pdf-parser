package main

import (
	"github.com/yrnThiago/pdf-ocr/config"
	grpcServer "github.com/yrnThiago/pdf-ocr/internal/grpc/server"
	httpServer "github.com/yrnThiago/pdf-ocr/internal/http/server"
	"github.com/yrnThiago/pdf-ocr/internal/infra/nats"
)

func main() {
	config.Init()
	config.LoggerInit()

	nats.Init()
	nats.PublisherInit()
	nats.ConsumerInit()

	go httpServer.Init()
	grpcServer.Init()
}
