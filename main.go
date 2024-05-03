package main

import (
	"log"

	"github.com/deviwlndr/ws-ATS/config"

	"github.com/deviwlndr/undangan_rapat"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/deviwlndr/undangan_rapat/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(_714220054.GetAllRapatMakrab()))
}