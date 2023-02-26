package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

type syntaxResponse struct {
	StartPosition sitter.Point `json:"startPosition"`
	EndPosition   sitter.Point `json:"endPosition"`
	StartByte     uint32       `json:"startByte"`
	EndByte       uint32       `json:"endByte"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))

	input := []byte("let x = 1; console.log(x);")

	parser := sitter.NewParser()
	typescript := typescript.GetLanguage()
	// javascript := javascript.GetLanguage()
	// python := python.GetLanguage()
	// golang := golang.GetLanguage()
	// rust := rust.GetLanguage()

	parser.SetLanguage(typescript)

	tree := parser.Parse(nil, input)
	firstChild := tree.RootNode().Child(1).Child(0)

	body := syntaxResponse{
		StartPosition: firstChild.StartPoint(),
		EndPosition:   firstChild.EndPoint(),
		StartByte:     firstChild.StartByte(),
		EndByte:       firstChild.EndByte(),
	}

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(body)
		return c.JSON(body)
	})

	app.Listen(":3000")
}
