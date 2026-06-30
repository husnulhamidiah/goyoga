package goyoga

import yoga "github.com/husnulhamidiah/goyoga/gen"

type Node struct {
	ref yoga.NodeRef
}

func wrapNode(ref yoga.NodeRef) *Node {
	if ref == nil {
		return nil
	}
	return &Node{ref: ref}
}

func (n *Node) yogaRef() yoga.NodeRef {
	if n == nil {
		return nil
	}
	return n.ref
}

func (n *Node) yogaConstRef() yoga.NodeConstRef {
	if n == nil {
		return nil
	}
	return yoga.NodeConstRef(n.ref)
}
