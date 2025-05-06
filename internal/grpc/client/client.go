package client

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/yrnThiago/pdf-ocr/api/pb"
	"github.com/yrnThiago/pdf-ocr/config"
)

type GrpcClient struct {
	PdfServiceClient pdf_ocr.PdfServiceClient
}

func NewGrpcClient() *GrpcClient {
	conn, err := grpc.NewClient(
		getGrpcUrl(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		config.Logger.Fatal(
			"failed to create grpc client",
			zap.Error(err),
		)
	}

	return &GrpcClient{
		PdfServiceClient: pdf_ocr.NewPdfServiceClient(conn),
	}
}

func (c *GrpcClient) ExtractFromPdf(id string) (*pdf_ocr.PdfResponse, error) {
	pdfResponse, err := c.PdfServiceClient.ExtractFromPdf(context.Background(), &pdf_ocr.PdfRequest{
		ID: id,
	})
	if err != nil {
		config.Logger.Fatal(
			"something went wrong",
			zap.Error(err),
		)
	}

	return pdfResponse, nil
}

func getGrpcUrl() string {
	return fmt.Sprintf("%s:%s", config.Env.GrpcHost, config.Env.GrpcPort)
}
