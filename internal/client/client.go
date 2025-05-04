package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/yrnThiago/pdf-ocr/internal/services/genproto"
)

type GrpcClient struct {
	PdfServiceClient pdf_ocr.PdfServiceClient
}

func NewGrpcClient() *GrpcClient {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	return &GrpcClient{
		PdfServiceClient: pdf_ocr.NewPdfServiceClient(conn),
	}
}

func (c *GrpcClient) AddPdf(filePath string) {
	_, err := c.PdfServiceClient.AddPdf(context.Background(), &pdf_ocr.AddPdfRequest{
		Path: filePath,
	})
	if err != nil {
		panic(err)
	}
}
