package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/pdf-ocr/config"
	"github.com/yrnThiago/pdf-ocr/internal/infra/nats"
	"github.com/yrnThiago/pdf-ocr/internal/utils"
	"go.uber.org/zap"
)

type HttpServer struct {
	port string
}

func NewHttpServer(port string) *HttpServer {
	return &HttpServer{port: port}
}

func Init() {
	server := NewHttpServer(config.Env.Port)
	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {

		file, err := c.FormFile("pdf")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "something went wrong"})
		}

		fileID := utils.GenerateUuid()
		filePath := utils.GetPdfPath(fileID)
		c.SaveFile(file, filePath)

		go nats.PdfPublisher.Publish(fileID)

		return c.Status(200).JSON(fiber.Map{"message": "wait until we proceess your file"})
	})

	config.Logger.Info("http server listening on", zap.String("port", server.port))

	config.Logger.Fatal(
		"something went wrong",
		zap.Error(app.Listen(":"+server.port)),
	)
}
