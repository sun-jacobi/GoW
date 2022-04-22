package Dom

// Interfaces
// There are 2 main classes of nodes in Dom model: Element and Text nodes.
// ----------------------------------------------------------------------------
// All node implement the Node interface, which have some children nodes
type Node interface{}

// All elements implement the Elem interface
type ElementNode interface {
	Node
	NewElem() *ElementNode
}

// All texts implement the Text interface
type TextNode interface {
	Node
	NewText() *TextNode
}

// ----------------------------------------------------------------------------
// Structures
type AttrMap map[string]string

type Elem struct {
	tag_name   string
	attributes AttrMap
	children   []Node
}

type Text struct {
	literal string
}

// ----------------------------------------------------------------------------
// Methods

func (t *Text) NewText(data string) *Text {
	return &Text{
		literal: data,
	}
}

func (t *Elem) NewElem(tag_name string, attributes AttrMap, children []Node) *Elem {
	p := new(Elem)
	p.tag_name = tag_name
	p.attributes = attributes
	p.children = children
	return p
}
