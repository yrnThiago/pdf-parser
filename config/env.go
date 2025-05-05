package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	GrpcHost string
	GrpcPort string

	Port string

	NatsUrl string
}

var Env EnvVariables

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(".env missing")
	}

	Env = EnvVariables{
		GrpcHost: os.Getenv("GRPC_HOST"),
		GrpcPort: os.Getenv("GRPC_PORT"),
		Port:     os.Getenv("API_PORT"),
		NatsUrl:  os.Getenv("NATS_URL"),
	}
}
