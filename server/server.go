package main

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

type syntaxNode struct {
	Type          string       `json:"type"`
	StartPosition sitter.Point `json:"startPosition"`
	EndPosition   sitter.Point `json:"endPosition"`
	StartByte     uint32       `json:"startByte"`
	EndByte       uint32       `json:"endByte"`
}

type syntaxList struct {
	SyntaxNodes []syntaxNode `json:"syntaxNodes"`
}

// TODO: Добавить автоматический спуск к ребенку когда сиблингов не осталось
func traverseTree(node *sitter.Node, sliceToFill *[]syntaxNode) {
	// NOTE: Возможно поменять на проверку StartPosition == 0 && EndPosition == 0
	if reflect.ValueOf(node).IsZero() {
		fmt.Println("Zero. Returning...")
		return
	}
	*sliceToFill = append(*sliceToFill, syntaxNode{
		Type:          node.Type(),
		StartPosition: node.StartPoint(),
		EndPosition:   node.EndPoint(),
		StartByte:     node.StartByte(),
		EndByte:       node.EndByte(),
	})
	traverseTree(node.NextNamedSibling(), sliceToFill)
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))

	input := []byte("let x = 1; console.log(x); console.log(1); console.log(2);")

	parser := sitter.NewParser()
	typescript := typescript.GetLanguage()
	// javascript := javascript.GetLanguage()
	// python := python.GetLanguage()
	// golang := golang.GetLanguage()
	// rust := rust.GetLanguage()

	parser.SetLanguage(typescript)

	tree := parser.Parse(nil, input)
	rootNode := tree.RootNode()

	allNamedChildren := make([]syntaxNode, 0)
	traverseTree(rootNode.Child(0), &allNamedChildren)
	fmt.Println(allNamedChildren)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(syntaxList{
			SyntaxNodes: allNamedChildren,
		})
	})

	app.Listen(":3000")
}
