package main

import (
	"github.com/gofiber/fiber/v2"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

func main() {
	app := fiber.New()

	input := []byte("function hello() { console.log('hello') }; function goodbye(){}")

	parser := sitter.NewParser()
	typescript := typescript.GetLanguage()
	// javascript := javascript.GetLanguage()
	// python := python.GetLanguage()
	// golang := golang.GetLanguage()
	// rust := rust.GetLanguage()

	parser.SetLanguage(typescript)

	tree := parser.Parse(nil, input)
	root := tree.RootNode()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(root.String())
	})

	app.Listen(":3000")
}
