package main

import (
	"grow/parser"
	"os"
)

type Engine struct {
	parser parser.Parser
}

func Setup() *Engine {
	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("Failed to open the file")
	}
	return &Engine{
		parser: *parser.NewParser(string(src)),
	}
}

func (*Engine) Run() {

}
