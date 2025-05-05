package types

import (
	"context"

	pdf_ocr "github.com/yrnThiago/pdf-ocr/internal/grpc/services/genproto"
)

type PdfService interface {
	ExtractFromPdf(context.Context, *pdf_ocr.Pdf) (string, error)
}
