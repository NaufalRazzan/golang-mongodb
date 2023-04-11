package main

import (
	"fmt"
	"golang-mongodb/configs"
	"golang-mongodb/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	fmt.Println("Application started and running smoothly")
	
	// run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	app.Listen(":6000")
}
