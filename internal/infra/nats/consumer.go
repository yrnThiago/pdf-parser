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

const pdfSubject = "pdf"

var pdfFilter = fmt.Sprintf("%s.>", pdfSubject)

type Consumer struct {
	Js          jetstream.JetStream
	Ctx         context.Context
	Config      jetstream.ConsumerConfig
	ConsumerCtx jetstream.Consumer
	GrpcClient  *client.GrpcClient
}

func NewConsumer(name, durable, filterSubject string) *Consumer {
	return &Consumer{
		Js:  JS,
		Ctx: context.Background(),
		Config: jetstream.ConsumerConfig{
			Name:          name,
			Durable:       durable,
			FilterSubject: filterSubject,
			AckPolicy:     jetstream.AckExplicitPolicy,
			DeliverPolicy: jetstream.DeliverAllPolicy,
		},
		GrpcClient: client.NewGrpcClient(),
	}
}

func ConsumerInit() {
	ordersConsumer := NewConsumer("pdf_processor", "pdf_processor", pdfFilter)
	ordersConsumer.CreateStream()
	ordersConsumer.HandlingNewPdfs()

	config.Logger.Info(
		"consumers successfully initialized",
	)
}

func (c *Consumer) HandlingNewPdfs() {
	_, err := c.ConsumerCtx.Consume(func(msg jetstream.Msg) {
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

func (c *Consumer) CreateStream() {
	stream, err := c.Js.Stream(c.Ctx, pdfSubject)
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}

	c.ConsumerCtx, err = stream.CreateOrUpdateConsumer(c.Ctx, c.Config)
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}
}

func getPdfIdFromMsg(msg jetstream.Msg) string {
	return strings.Replace(string(msg.Subject()), fmt.Sprintf("%s.", pdfSubject), "", 1)
}
