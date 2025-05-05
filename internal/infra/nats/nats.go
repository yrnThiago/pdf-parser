package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/pdf-ocr/config"
	"github.com/yrnThiago/pdf-ocr/internal/utils"
)

var (
	NatsConn  *nats.Conn
	JetStream jetstream.JetStream
)

func Init() {
	var err error
	natsURL := getNatsURL()
	NatsConn, err = nats.Connect(natsURL)
	if err != nil {
		config.Logger.Fatal(
			"unable to connect with nats server",
		)
	}

	JetStream, err = jetstream.New(NatsConn)
	if err != nil {
		config.Logger.Fatal(
			"jetstream",
			zap.Error(err),
		)
	}

	config.Logger.Info(
		"Nats successfully initialized",
	)
}

func getNatsURL() string {
	if utils.IsEmpty(config.Env.NatsUrl) {
		return nats.DefaultURL
	}

	return config.Env.NatsUrl
}
