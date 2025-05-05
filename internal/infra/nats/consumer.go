package nats

import (
	"context"
	"fmt"
	"strings"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/pdf-ocr/config"
	"github.com/yrnThiago/pdf-ocr/internal/grpc/client"
)

const (
	pdfSubject = "pdf"
	pdfFilter  = pdfSubject + ".>"
)

type Consumer struct {
	ConsumerJS jetstream.Consumer
	GrpcClient *client.GrpcClient
}

func ConsumerInit() {
	ordersConsumer := NewConsumer("pdf_processor", "pdf_processor", pdfFilter)
	ordersConsumer.HandlingNewPdfs()

	config.Logger.Info(
		"consumers successfully initialized",
	)
}

func NewConsumer(name, durable, filterSubject string) *Consumer {
	consumerConfig := NewConsumerConfig(name, durable, filterSubject)

	return &Consumer{
		ConsumerJS: NewConsumerJS(consumerConfig),
		GrpcClient: client.NewGrpcClient(),
	}
}

func NewConsumerConfig(name, durable, filterSubject string) jetstream.ConsumerConfig {
	return jetstream.ConsumerConfig{
		Name:          name,
		Durable:       durable,
		FilterSubject: filterSubject,
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
	}
}

func NewConsumerJS(configJS jetstream.ConsumerConfig) jetstream.Consumer {
	ctx := context.Background()
	stream, err := JetStream.Stream(ctx, pdfSubject)
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}

	consumerJS, err := stream.CreateOrUpdateConsumer(ctx, configJS)
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}

	return consumerJS
}

func (c *Consumer) HandlingNewPdfs() {
	_, err := c.ConsumerJS.Consume(func(msg jetstream.Msg) {
		pdfID := getPdfIdFromMsg(msg)
		msg.Ack()

		config.Logger.Info(
			"new pdf received",
			zap.String("pdf id", pdfID),
		)

		_, err := c.GrpcClient.ExtractFromPdf(pdfID)
		if err != nil {
			panic(err)
		}

		config.Logger.Info("pdf successfully processed", zap.String("pdf id", pdfID))
	})
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}
}

func getPdfIdFromMsg(msg jetstream.Msg) string {
	return strings.Replace(string(msg.Subject()), fmt.Sprintf("%s.", pdfSubject), "", 1)
}
