package main

import (
	"log"

	"github.com/rayfanaqbil/Zenverse-BP/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/rayfanaqbil/Zenverse-BP/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/rayfanaqbil/Zenverse-BP/docs"
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/rayfanaqbil
// @contact.email 714220044.@std.ulbi.ac.id

// @host zenversegames-ba223a40f69e.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
