package framework

import (
	"testing"
)

func Test_filterChildNodes(t *testing.T) {
	root := &node{
		isLeaf:   false,
		segment:  "",
		handlers: ControllerHandlerChain{func(*XContext) error { return nil }},
		children: []*node{
			{
				isLeaf:   true,
				segment:  "FOO",
				handlers: ControllerHandlerChain{func(*XContext) error { return nil }},
				children: nil,
			},
			{
				isLeaf:   false,
				segment:  ":id",
				handlers: nil,
				children: nil,
			},
		},
	}

	{
		nodes := root.filterChildNodes("FOO")
		if len(nodes) != 2 {
			t.Error("foo error")
		}
	}

	{
		nodes := root.filterChildNodes(":foo")
		if len(nodes) != 2 {
			t.Error(":foo error")
		}
	}

}

func Test_matchNode(t *testing.T) {
	root := &node{
		isLeaf:   false,
		segment:  "",
		handlers: ControllerHandlerChain{func(*XContext) error { return nil }},
		children: []*node{
			{
				isLeaf:   true,
				segment:  "FOO",
				handlers: nil,
				children: []*node{
					&node{
						isLeaf:   true,
						segment:  "BAR",
						handlers: ControllerHandlerChain{func(*XContext) error { panic("not implemented") }},
						children: []*node{},
					},
				},
			},
			{
				isLeaf:   true,
				segment:  ":id",
				handlers: nil,
				children: nil,
			},
		},
	}

	{
		node := root.matchNode("foo/bar")
		if node == nil {
			t.Error("match normal node error")
		}
	}

	{
		node := root.matchNode("test")
		if node == nil {
			t.Error("match test")
		}
	}

}
