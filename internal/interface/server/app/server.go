package app

import (
	"app/internal/registry"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Server struct {
	router *fiber.App
}

func NewAppServer() *Server {
	return &Server{
		router: fiber.New(),
	}
}

func (s *Server) GetRouter() *fiber.App {
	return s.router
}

func (s *Server) InitRoutes(cnt *registry.Container) {
	s.router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	s.router.Get("/liveness", func(c *fiber.Ctx) error {
		c.SendStatus(200)

		return nil
	})

	s.router.Get("/readiness", func(c *fiber.Ctx) error {
		c.SendStatus(200)

		return nil
	})

	s.router.Post("/import/company/:companyName", func(c *fiber.Ctx) error {
		err := cnt.Usecases.Import.ImportCompany(
			context.Background(),
			c.Params("companyName"),
			c.Body(),
		)

		if err != nil {
			c.Response().SetStatusCode(500)
			log.Error(err)
			return nil
		}

		c.Response().SetStatusCode(200)
		return nil
	})
}

func (s *Server) Run(port string) error {
	err := s.router.Listen(port)

	return err
}
