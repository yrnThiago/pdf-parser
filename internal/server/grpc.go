package server

import (
	"log"
	"net"

	handler "github.com/yrnThiago/pdf-ocr/internal/handler/pdf"
	"github.com/yrnThiago/pdf-ocr/internal/usecase"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pdfService := usecase.NewPdfUseCase()
	handler.NewPdfService(grpcServer, pdfService)

	log.Printf("listening port on %s", ":50051")

	return grpcServer.Serve(listen)
}
