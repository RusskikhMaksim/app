package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

var BranchName = "development"

func main() {
	app := fiber.New()
	log.Println("Version: ", BranchName)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/liveness", func(c *fiber.Ctx) error {
		c.SendStatus(200)

		return nil
	})

	app.Get("/readiness", func(c *fiber.Ctx) error {
		c.SendStatus(200)

		return nil
	})

	app.Post("/import/company", func(c *fiber.Ctx) error {
		dec := json.NewDecoder(bytes.NewReader(c.Body()))
		r := &Company{}

		if err := dec.Decode(r); err != nil {
			log.Println(err.Error())
		}

		log.Printf("%+v", r)

		c.Response().SetStatusCode(200)
		return nil
	})

	app.Listen(":3000")
}
