package main

import (
	"io"
	"os"
)

// MarkDownNode
type MarkDownNode interface{}

type MarkDownTree struct {
	Nodes []MarkDownNode
}

// BlockText
type Block struct {
	*Text
}

// Text
type Text struct {
	Parts []MarkDownNode
}

// Em
type Em struct {
	Text string
}

type MarkDownParser struct {
	source []byte
}

type HTMLMarkDownRenderer struct {
	node MarkDownNode
}

func NewParser() *MarkDownParser {
	return &MarkDownParser{[]byte{}}
}

func (m *MarkDownParser) Reset() {
	m.source = m.source[:0]
}

func parseBlockNodes(input []byte) []MarkDownNode {
	var nodes []MarkDownNode
	for len(input) > 0 {

	}

	return nodes
}

func parseInlineNodes(input []byte) *Text {
	var text = new(Text)
	var nodes []MarkDownNode

	for len(input) > 0 {

	}

	text.Parts = nodes
	return text
}

func (m *MarkDownParser) Parse(input []byte) MarkDownTree {
	tree := MarkDownTree{}
	tree.Nodes = parseBlockNodes(input)

	return tree
}

func NewRenderer(node MarkDownNode) *HTMLMarkDownRenderer {
	return &HTMLMarkDownRenderer{
		node: node,
	}
}

func (r *HTMLMarkDownRenderer) WriteTo(node MarkDownNode, w io.Writer) (n int64, err error) {
	switch r.node.(type) {
	case *Text:
		for _, part := range r.node.(*Text).Parts {
			r.WriteTo(part, w)
		}
	case *Em:
		w.Write([]byte("<em>"))
		w.Write([]byte(r.node.(*Em).Text))
		w.Write([]byte("</em>"))
	}

	return 0, nil
}

func main() {
	p := NewParser()
	tree := p.Parse([]byte("Hello _World_"))
	r := NewRenderer(tree)
	r.WriteTo(nil, os.Stdout)
	p.Reset()
}
