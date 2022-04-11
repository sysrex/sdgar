package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/sysrex/sdgar/docs"
	"honnef.co/go/tools/config"
)

//@Description Geting for welcome endpoint
//@Tags Welcome
//@Accept json
//@Product json
//@Success 200 {string} string
//@router /api [get]
func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to awsome Go fiber rest api")
}

//Function making for setting internal routes
func SetupRoutes(app *fiber.App) {
	apiVersion := "/v1"
	apiURL := "/api"
	//Welcome endpoint
	//not required AuthorizationRequired
	app.Get(apiVersion+apiURL, welcome)

	//Swagger endpoint making for API document
	app.Get("/v1/api/swagger/*", swagger.Handler)

}

func main() {
	//Connect to the daatbase
	database.ConnectDatabase()

	appPort := config.Config("APP_PORT")
	//Fiber new instance
	app := fiber.New()

	//Used logger for keeppink request logging ID
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-01-2006 15:04:05",
		TimeZone:   "Europe/London",
	}))

	SetupRoutes(app)
	log.Fatal(app.Listen(":" + appPort))
}
