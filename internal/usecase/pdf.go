package usecase

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ledongthuc/pdf"

	pdf_ocr "github.com/yrnThiago/pdf-ocr/internal/services/genproto"
)

type PdfUseCase struct{}

func NewPdfUseCase() *PdfUseCase {
	return &PdfUseCase{}
}

func (p *PdfUseCase) AddPdf(ctx context.Context, req *pdf_ocr.Pdf) error {
	getContentFromPdf(req.Path)

	return nil
}

func getContentFromPdf(path string) error {
	pdf.DebugOn = true
	content, err := readPdf(path) // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return nil
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
