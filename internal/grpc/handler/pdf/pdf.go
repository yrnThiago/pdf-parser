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

func (h *PdfHandler) ExtractFromPdf(
	ctx context.Context,
	req *pdf_ocr.PdfRequest,
) (*pdf_ocr.PdfResponse, error) {
	pdf := &pdf_ocr.Pdf{
		ID: req.ID,
	}

	content, err := h.pdfService.ExtractFromPdf(ctx, pdf)
	if err != nil {
		return nil, err
	}

	res := &pdf_ocr.PdfResponse{
		User: &pdf_ocr.User{
			ID:   "123",
			Name: "Nome",
			// ...
		},
		Text: content,
	}

	return res, nil
}
