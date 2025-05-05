package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/pdf-ocr/config"
)

var PdfPublisher *Publisher

type Publisher struct {
	Config jetstream.StreamConfig
}

func PublisherInit() {
	PdfPublisher = NewPublisher("pdf", "Msgs for pdf", "pdf.>")

	config.Logger.Info(
		"publishers successfully initialized",
	)
}

func NewPublisher(name, description, subject string) *Publisher {
	return &Publisher{
		Config: NewPublisherConfig(name, description, subject),
	}
}

func NewPublisherConfig(name, description, subject string) jetstream.StreamConfig {
	return jetstream.StreamConfig{
		Name:        name,
		Description: description,
		Subjects: []string{
			subject,
		},
		MaxBytes: 1024 * 1024 * 1024,
	}
}

func (p *Publisher) Publish(msg string) {
	ctx := context.Background()
	_, err := JetStream.Publish(ctx, fmt.Sprintf("pdf.%s", msg), []byte("new pdf"))
	if err != nil {
		config.Logger.Warn(
			"msg cant be published",
			zap.Error(err),
		)
	}

	config.Logger.Info(
		"publishing new pdf",
		zap.String("pdf id", msg),
	)
}
