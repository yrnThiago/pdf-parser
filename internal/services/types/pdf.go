package types

import (
	"context"

	pdf_ocr "github.com/yrnThiago/pdf-ocr/internal/services/genproto"
)

type PdfService interface {
	AddPdf(context.Context, *pdf_ocr.Pdf) (string, error)
}
