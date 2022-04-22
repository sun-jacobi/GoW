package parser

import (
	"bytes"
	"errors"
)

// ----------------------------------------------------------------------------
// Parser for the HTML
type Parser struct {
	input string
	pos   int // current position in input (points to current char)
}

func New_Parser(input string) *Parser {
	p := new(Parser)
	p.input = input
	p.pos = 0
	return p
}

// ----------------------------------------------------------------------------
// Public methods for the html parser

// Read the current char without consumming it
func (p Parser) Next() (byte, error) {
	if !p.eof() {
		return byte(0), errors.New("Parser::Next")
	} else {
		return p.input[p.pos+1], nil
	}
}

// Return the current char, and advance self.pos
func (p *Parser) Consume() byte {
	c := p.input[p.pos]
	p.pos += 1
	return c
}

// Consume char until test fail
func (p *Parser) Consume_while(test func(byte) bool) (string, error) {
	var result bytes.Buffer
	var cur_char byte
	next_char, e := p.Next()
	if e != nil {
		return result.String(), errors.New("Parser::Consume_while")
	}
	for !test(next_char) {
		cur_char = p.Consume()
		result.WriteByte(cur_char)
		next_char, e = p.Next()
		if e != nil {
			return result.String(), errors.New("Parser::Consume_while")
		}
	}
	return result.String(), nil
}

// Comsume the whitespace
func (p *Parser) Consume_whitespace() {
	p.Consume_while(
		func(b byte) bool {
			return (b == 32)
		})
}

// Parse a tag_name or attribute
func (p *Parser) Parse_tag() (string, error) {
	s, e := p.Consume_while(
		func(b byte) bool {
			return (b > 97 && b < 122) ||
				(b > 65 && b < 90) || (b > 48 && b < 57)
		})
	if e != nil {
		return s, errors.New("Parser::Parse_tag")
	} else {
		return s, nil
	}
}

// Parse text
func (p *Parser) Parse_text() (Text, error)

// Parse element

// Parse a node
/*
func (p *Parser) Parse_node() (string, error) {
	if p.Next() == '<' {
		p.
	}
}
*/

// ----------------------------------------------------------------------------
// Utility Functions (private)

// Get the char at a certain position
func (p Parser) charAt(index int) byte {
	return p.input[index]
}

// Get the length of whole input
func (p Parser) len() int {
	return len(p.input)
}

func (p Parser) eof() bool {
	return p.pos < p.len()-1
}
