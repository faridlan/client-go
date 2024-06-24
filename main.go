package main

import (
	"github.com/faridlan/client-go/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// var token = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiYzg1YzkyYjYxZTI0NGM5Yjg1YWQ2YjkyYTBhZTM3NDQiLCJ1c2VybmFtZSI6InVzZXIwMDEifSwiZXhwIjoxNzE2MjE1NzgzfQ.S2eAmll1KoObtJYXPLzZCAxzXDbxNXB3Uim5dP8d7QtZE417rLkW0ZaJV98_jhAtCyt8eBB55KUKz3wsH-xxNw"

func main() {

	engine := html.New("./public/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	router.SetupRoute(app)

	app.Listen("localhost:8080")

}
