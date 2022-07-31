package main

import (
	"grow/dom"
	"grow/parser"
	"grow/render"
	"os"
)

type Engine struct {
	parser parser.Parser
	render render.Render
}

func Setup() *Engine {
	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("Failed to open the file")
	}
	return &Engine{
		parser: *parser.NewParser(string(src)),
		render: render.Render{},
	}
}

func (engine *Engine) Run() {
	dom := engine.parse()
	engine.rendering(dom)
}

// Private Methods

func (engine *Engine) rendering(dom dom.Node) {
	engine.render.Rendering(dom)
}

func (engine *Engine) parse() dom.Node {
	return engine.parser.Parse()
}
