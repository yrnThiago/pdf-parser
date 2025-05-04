package usecase

import (
	"bytes"
	"context"

	"github.com/ledongthuc/pdf"

	"github.com/yrnThiago/pdf-ocr/internal/client"
	pdf_ocr "github.com/yrnThiago/pdf-ocr/internal/services/genproto"
)

type PdfUseCase struct {
	GrpcClient *client.GrpcClient
}

func NewPdfUseCase() *PdfUseCase {
	return &PdfUseCase{
		GrpcClient: client.NewGrpcClient(),
	}
}

func (p *PdfUseCase) AddPdf(ctx context.Context, req *pdf_ocr.Pdf) (string, error) {
	return getContentFromPdf(req.Path)
}

func getContentFromPdf(path string) (string, error) {
	pdf.DebugOn = true
	content, err := readPdf(path) // Read local pdf file
	if err != nil {
		panic(err)
	}
	return content, nil
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
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
