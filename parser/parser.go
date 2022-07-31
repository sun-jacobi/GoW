package parser

import (
	"bytes"
	"errors"
	"grow/dom"
)

// ----------------------------------------------------------------------------
// Parser
type Parser struct {
	input string
	// current position in input (point to the char next to the lastchar
	pos      int
	lastchar byte // this char
}

func NewParser(input string) *Parser {
	return &Parser{
		input:    input,
		pos:      0,
		lastchar: 0,
	}
}

// ----------------------------------------------------------------------------
// Public methods for the html parser

// Read the next char without consuming it
func (p Parser) Next() (byte, error) {
	if !p.eof() {
		return byte(0), errors.New("Error: No next char when consuming.")
	} else {
		return p.input[p.pos], nil
	}
}

// Return the next char, and advance self.pos
func (p *Parser) Consume() (byte, error) {
	if !p.eof() {
		return byte(0), errors.New("Error: No next char when consuming.")
	}
	thischar := p.input[p.pos]
	p.pos += 1
	p.lastchar = thischar
	return thischar, nil
}

// Consume char until test fail
func (p *Parser) Consume_while(test func(byte) bool) (string, error) {
	var result bytes.Buffer
	var ch byte
	var err error
	for ch, err = p.Consume(); !test(ch) || err != nil; {
		result.WriteByte(ch)
		if !p.eof() {
			break
		}
	}
	if err != nil {
		return result.String(), nil
	}
	return result.String(), nil
}

// Comsume the whitespace
func (p *Parser) Consume_whitespace() error {
	_, err := p.Consume_while(
		func(b byte) bool {
			return (b == 32)
		})
	return err
}

// Parse a tag or attribute
func (p *Parser) Parse_tag() (string, error) {
	tag, err := p.Consume_while(
		func(b byte) bool {
			return (b >= byte('a') && b <= byte('z')) ||
				(b >= byte('A') && b <= byte('Z')) || (b >= byte('0') && b <= byte('9'))
		})
	if err != nil {
		return tag, err
	} else {
		return tag, nil
	}
}

// Parse Node
func (p *Parser) Parse_node() (dom.Node, error) {
	ch, err := p.Consume()
	if err != nil {
		return nil, err
	}
	switch ch {
	case byte('<'):
		return p.Parse_element()
	default:
		return p.Parse_text()
	}
}

// Parse Nodes
func (p *Parser) Parse_nodes() ([]dom.Node, error) {
	nodes := make([]dom.Node, 0)
	for {
		err := p.Consume_whitespace()
		if err != nil {
			return nil, err
		}
		nextChar, _ := p.Next()
		if !p.eof() && (nextChar == '/' && p.lastchar == '<') {
			break
		}
		node, _ := p.Parse_node()
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// Parse Text
func (p *Parser) Parse_text() (*dom.Text, error) {
	text, err := p.Consume_while(
		func(b byte) bool {
			return b != byte('<')
		})
	if err != nil {
		return nil, err
	}
	return dom.NewText(text), nil
}

// Parse element
func (p *Parser) Parse_element() (*dom.Elem, error) {
	if p.lastchar != byte('<') {
		return nil, errors.New("Expected \" < \" ")
	}
	tag, err := p.Parse_tag()
	if err != nil {
		return nil, err
	}
	attrs, err := p.Parse_attrs()
	children, err := p.Parse_nodes()

	return dom.NewElem(tag, *attrs, children), nil
}

// Parse attribute_value
func (p *Parser) Parse_attr_val() (string, error) {
	quote, err := p.Consume()
	if err != nil {
		return "", err
	}
	if quote != byte('\'') {
		return "", errors.New("Expected \" ")
	}
	value, err := p.Consume_while(
		func(b byte) bool { return b != quote })
	return value, err
}

// Parse attribute
func (p *Parser) Parse_attr() (string, string, error) {
	name, err := p.Parse_tag()
	val, err := p.Parse_attr_val()
	return name, val, err
}

// Parse attributes
func (p *Parser) Parse_attrs() (*map[string]string, error) {
	attrs := make(map[string]string)
	for {
		err := p.Consume_whitespace()
		if err != nil {
			return nil, err
		}
		nextChar, err := p.Next()
		if nextChar == byte('>') {
			break
		}
		name, value, err := p.Parse_attr()
		attrs[name] = value

	}
	return &attrs, nil
}

func (p *Parser) Parse() dom.Node {
	nodes, _ := p.Parse_nodes()
	return dom.NewElem("html", make(map[string]string), nodes)
}

// ----------------------------------------------------------------------------
// Utility Functions (private)

// Get the char at a certain position
func (p *Parser) charAt(index int) byte {
	return p.input[index]
}

// Get the length of whole input
func (p *Parser) len() int {
	return len(p.input)
}

func (p *Parser) eof() bool {
	return p.pos <= p.len()-1
}
