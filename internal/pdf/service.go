package pdf

import (
	"context"

	"github.com/yrnThiago/pdf-ocr/api/pb"
)

type PdfService interface {
	ExtractFromPdf(context.Context, *pdf_ocr.Pdf) (string, error)
}
