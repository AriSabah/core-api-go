package server

import (
	"github.com/gofiber/fiber/v2"

	"core-api-go/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "core-api-go",
			AppName:      "core-api-go",
		}),

		db: database.New(),
	}

	return server
}
