package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUuid() string {
	return uuid.New().String()
}

func GetPdfPath(fileId string) string {
	return fmt.Sprintf("internal/uploads/%s.pdf", fileId)
}
func IsEmpty(param string) bool {
	return param == ""
}
