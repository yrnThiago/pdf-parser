package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/yrnThiago/pdf-ocr/config"
	"github.com/yrnThiago/pdf-ocr/internal/services/genproto"
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
		panic(err)
	}

	return &GrpcClient{
		PdfServiceClient: pdf_ocr.NewPdfServiceClient(conn),
	}
}

func (c *GrpcClient) AddPdf(filePath string) (string, error) {
	pdfResponse, err := c.PdfServiceClient.AddPdf(context.Background(), &pdf_ocr.AddPdfRequest{
		Path: filePath,
	})
	if err != nil {
		panic(err)
	}

	return pdfResponse.Text, nil
}

func getGrpcUrl() string {
	return fmt.Sprintf("%s:%s", config.Env.GrpcHost, config.Env.GrpcPort)
}
