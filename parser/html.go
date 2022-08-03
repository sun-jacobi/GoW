package parser

import (
	"errors"
	"grow/dom"
)

type HtmlParser struct {
	Parser
}

func (p *HtmlParser) Parse_tag() (string, error) {
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
func (p *HtmlParser) Parse_node() (dom.Node, error) {
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
func (p *HtmlParser) Parse_nodes() ([]dom.Node, error) {
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
func (p *HtmlParser) Parse_text() (*dom.Text, error) {
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
func (p *HtmlParser) Parse_element() (*dom.Elem, error) {
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
func (p *HtmlParser) Parse_attr_val() (string, error) {
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
func (p *HtmlParser) Parse_attr() (string, string, error) {
	name, err := p.Parse_tag()
	val, err := p.Parse_attr_val()
	return name, val, err
}

// Parse attributes
func (p *HtmlParser) Parse_attrs() (*map[string]string, error) {
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

func (p *HtmlParser) Parse() dom.Node {
	nodes, _ := p.Parse_nodes()
	return dom.NewElem("html", make(map[string]string), nodes)
}
