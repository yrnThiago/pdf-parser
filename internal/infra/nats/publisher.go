package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/pdf-ocr/config"
)

var PdfPublisher *Publisher

type Publisher struct {
	NatsConn *nats.Conn
	Js       jetstream.JetStream
	Ctx      context.Context
	Config   jetstream.StreamConfig
}

func NewPublisher(name, description, subject string) *Publisher {
	return &Publisher{
		NatsConn: NC,
		Js:       JS,
		Ctx:      context.Background(),
		Config: jetstream.StreamConfig{
			Name:        name,
			Description: description,
			Subjects: []string{
				subject,
			},
			MaxBytes: 1024 * 1024 * 1024,
		},
	}
}

func (p *Publisher) Publish(msg string) {
	_, err := p.Js.Publish(p.Ctx, fmt.Sprintf("pdf.%s", msg), []byte("new pdf"))
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

func (p *Publisher) CreateStream() {
	_, err := p.Js.CreateOrUpdateStream(p.Ctx, p.Config)
	if err != nil {
		config.Logger.Fatal(
			"publisher cant be initialized",
			zap.Error(err),
		)
	}
}

func PublisherInit() {
	PdfPublisher = NewPublisher("pdf", "Msgs for pdf", "pdf.>")
	PdfPublisher.CreateStream()

	config.Logger.Info(
		"publishers successfully initialized",
	)
}
