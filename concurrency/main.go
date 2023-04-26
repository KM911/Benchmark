package main

import (
	"app/controller"
	"app/dao"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	apiRouter := app.Group("/api")
	apiRouter.Get("/likes/:id", controller.Like)
	apiRouter.Get("/query/:id", controller.Query)
}
func main() {
	app := fiber.New()
	InitRouter(app)
	dao.InitMySql()
	app.Listen(":3000")
}
