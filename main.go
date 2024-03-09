package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
	"strings"
)

func main() {
	engine := html.New("./template", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layout",
	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/index2")
	})
	app.Get("/:page", func(ctx *fiber.Ctx) error {
		page := ctx.Params("page", "index1")
		return ctx.Render(strings.Trim(page, ".html"), fiber.Map{})
	})
	log.Fatal(app.Listen(":3000"))
}
