package main

import (
	"github.com/yrnThiago/pdf-ocr/config"
	httpServer "github.com/yrnThiago/pdf-ocr/internal/http/server"
	"github.com/yrnThiago/pdf-ocr/internal/infra/nats"
)

func main() {
	config.Init()
	config.LoggerInit()

	nats.Init()
	nats.PublisherInit()
	nats.ConsumerInit()

	httpServer.Init()
}
