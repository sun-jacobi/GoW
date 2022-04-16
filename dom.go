package main

type Element struct {
	tag_name   string
	attributes AttrMap
}

type Value struct {
	element Element
	text    string
}

type AttrMap map[string]string

func text(data string) *Node {
	p := new(Node)
	p.value = Value{text: data}
	return p
}

func elem(tag_name string, attributes AttrMap, children []*Node) *Node {
	p := new(Node)
	p.children = children
	p.value = Value{
		element: Element{
			tag_name:   tag_name,
			attributes: attributes,
		},
	}
	return p
}

type Node struct {
	children []*Node
	value    Value
}
