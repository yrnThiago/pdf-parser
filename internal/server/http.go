package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
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

		filePath := "internal/uploads/" + file.Filename
		c.SaveFile(file, filePath)

		client := client.NewGrpcClient()
		client.AddPdf(filePath)

		return c.Status(200).JSON(fiber.Map{"message": "success"})
	})

	log.Printf("Starting server on %s", h.addr)

	return app.Listen(h.addr)
}
