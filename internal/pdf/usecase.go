package pdf

import (
	"bytes"
	"context"

	"github.com/ledongthuc/pdf"
	"go.uber.org/zap"

	"github.com/yrnThiago/pdf-ocr/api/pb"
	"github.com/yrnThiago/pdf-ocr/config"
	"github.com/yrnThiago/pdf-ocr/internal/grpc/client"
	"github.com/yrnThiago/pdf-ocr/internal/utils"
)

type PdfUseCase struct {
	GrpcClient *client.GrpcClient
}

func NewPdfUseCase() *PdfUseCase {
	return &PdfUseCase{
		GrpcClient: client.NewGrpcClient(),
	}
}

func (p *PdfUseCase) ExtractFromPdf(ctx context.Context, req *pdf_ocr.Pdf) (string, error) {
	path := utils.GetPdfPath(req.ID)
	return getContentFromPdf(path)
}

func getContentFromPdf(path string) (string, error) {
	content, err := readPdf(path)
	if err != nil {
		config.Logger.Fatal(
			"file not found",
			zap.String("path", path),
		)
	}

	return content, nil
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)

	defer f.Close()

	if err != nil {
		return "", err
	}
	var buf bytes.Buffer

	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)

	return buf.String(), nil
}
