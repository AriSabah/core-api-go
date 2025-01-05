package server

import (
	"github.com/gofiber/fiber/v2"

	"fs-api-go/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "fs-api-go",
			AppName:      "fs-api-go",
		}),

		db: database.New(),
	}

	return server
}
