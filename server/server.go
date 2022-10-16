package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitializeApp() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)
	router.Get("/urls", getAllUrlies)
	router.Get("/url/:id", getUrl)
	router.Post("/url", createRewrite)
	router.Patch("/url", updateRewrite)
	router.Delete("/url/:id", deleteUrl)

	router.Listen(":5000")

}
