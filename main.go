package main

import (
	"fiber_jwt/dbase"
	"fiber_jwt/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//mysql db connect
	dbase.Connect()

	app := fiber.New()
	router.Setup(app)

	app.Listen(":8080")
}
