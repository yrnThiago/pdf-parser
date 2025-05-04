package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yrnThiago/pdf-ocr/internal/client"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (h *HttpServer) Run() error {
	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {

		file, err := c.FormFile("pdf")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "something went wrong"})
		}

		fileID := uuid.New().String()
		filePath := "internal/uploads/" + fileID + ".pdf"
		c.SaveFile(file, filePath)

		client := client.NewGrpcClient()
		content, err := client.AddPdf(filePath)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "something went wrong"})
		}

		return c.Status(200).JSON(fiber.Map{"message": content})
	})

	log.Printf("Starting server on %s", h.addr)
	return app.Listen(h.addr)
}
