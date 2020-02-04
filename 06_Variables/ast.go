package main

import "fmt"

type AstNode struct {
	left  *AstNode
	right *AstNode
	op    int
	v     AstNodeValue
}

type AstNodeValue struct {
	intval int
	id     int
}

func makeAstNode(op int, left *AstNode, right *AstNode, intval int) *AstNode {
	n := AstNode{
		v: AstNodeValue{
			intval: intval,
		},
		left:  left,
		right: right,
		op:    op,
	}
	return &n
}

func makeLeaf(op int, intval int) *AstNode {
	return makeAstNode(op, nil, nil, intval)
}

func mkastunary(op int, left *AstNode, intval int) *AstNode {
	return makeAstNode(op, left, nil, intval)
}

func printAstDepth(n *AstNode) {
	fmt.Printf(" Height %d \n", getDepthStupid(n))

}
func getDepthStupid(n *AstNode) int {
	if n == nil {
		return 0
	}
	fmt.Printf("%v ", n)

	depth := 1
	depthLeft := getDepthStupid(n.left)
	depthRight := getDepthStupid(n.right)

	if depthLeft > depthRight {
		depth += depthLeft
	} else {
		depth += depthRight
	}

	return depth
}
