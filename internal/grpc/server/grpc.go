package server

import (
	"net"

	"github.com/yrnThiago/pdf-ocr/config"
	handler "github.com/yrnThiago/pdf-ocr/internal/grpc/handler/pdf"
	"github.com/yrnThiago/pdf-ocr/internal/grpc/services"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	pdfService := services.NewPdfUseCase()
	handler.NewPdfService(grpcServer, pdfService)

	config.Logger.Info("grpc server listening on", zap.String("port", server.port))

	config.Logger.Fatal(
		"something went wrong",
		zap.Error(grpcServer.Serve(listen)),
	)
}
