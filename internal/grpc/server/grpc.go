package server

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/yrnThiago/pdf-ocr/config"
	"github.com/yrnThiago/pdf-ocr/internal/grpc/handler"
	"github.com/yrnThiago/pdf-ocr/internal/pdf"
)

type gRPCServer struct {
	host string
	port string
}

func NewGRPCServer() *gRPCServer {
	return &gRPCServer{host: config.Env.GrpcHost, port: config.Env.GrpcPort}
}

func Init() {
	server := NewGRPCServer()
	listen, err := net.Listen("tcp", ":"+server.port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pdfService := pdf.NewPdfUseCase()
	handler.NewPdfService(grpcServer, pdfService)

	config.Logger.Info("grpc server listening on", zap.String("port", server.port))

	config.Logger.Fatal(
		"something went wrong",
		zap.Error(grpcServer.Serve(listen)),
	)
}
