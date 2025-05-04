package handler

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yrnThiago/pdf-ocr/internal/grpc/services/genproto"
	"github.com/yrnThiago/pdf-ocr/internal/grpc/services/types"
)

type PdfHandler struct {
	pdfService types.PdfService
	pdf_ocr.UnimplementedPdfServiceServer
}

func NewPdfService(grpc *grpc.Server, pdfService types.PdfService) {
	handler := &PdfHandler{
		pdfService: pdfService,
	}

	pdf_ocr.RegisterPdfServiceServer(grpc, handler)
}

func (h *PdfHandler) AddPdf(
	ctx context.Context,
	req *pdf_ocr.AddPdfRequest,
) (*pdf_ocr.AddPdfResponse, error) {
	pdf := &pdf_ocr.Pdf{
		Path: req.Path,
	}

	content, err := h.pdfService.AddPdf(ctx, pdf)
	if err != nil {
		return nil, err
	}

	res := &pdf_ocr.AddPdfResponse{
		Text: content,
	}

	return res, nil
}
