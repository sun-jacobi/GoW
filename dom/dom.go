package dom

// Interfaces
// There are 2 main classes of nodes in Dom model: Element and Text nodes.
// ----------------------------------------------------------------------------
// All node implement the Node interface, which have some children nodes
type Node interface {
	Node()
}

// All elements implement the Elem interface
type ElementNode interface {
	Node
	ElementNode()
}

// All texts implement the Text interface
type TextNode interface {
	Node
	TextNode()
}

// ----------------------------------------------------------------------------
// Structures

type Elem struct {
	tag_name   string
	attributes map[string]string
	children   []Node
}

type Text struct {
	literal string
}

// ----------------------------------------------------------------------------
// Methods

func NewText(data string) *Text {
	return &Text{
		literal: data,
	}
}

func NewElem(tag_name string, attributes map[string]string, children []Node) *Elem {
	p := new(Elem)
	p.tag_name = tag_name
	p.attributes = attributes
	p.children = children
	return p
}

// ----------------------------------------------------------------------------
func (*Text) Node()        {}
func (*Text) TextNode()    {}
func (*Elem) ElementNode() {}
func (*Elem) Node()        {}
