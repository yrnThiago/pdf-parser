package server

import (
	"log"
	"net"

	"github.com/yrnThiago/pdf-ocr/config"
	handler "github.com/yrnThiago/pdf-ocr/internal/handler/pdf"
	"github.com/yrnThiago/pdf-ocr/internal/usecase"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	host string
	port string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{host: config.Env.GrpcHost, port: config.Env.GrpcPort}
}

func (s *gRPCServer) Run() error {
	listen, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pdfService := usecase.NewPdfUseCase()
	handler.NewPdfService(grpcServer, pdfService)

	log.Printf("listening port on %s", s.port)

	return grpcServer.Serve(listen)
}
