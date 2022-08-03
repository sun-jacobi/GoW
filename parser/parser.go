package parser

import (
	"bytes"
	"errors"
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
// Basic methods for the parser

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
