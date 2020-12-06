package ast

import "github.com/lw396285v/goscript/common"

type ASTNode interface {
	SetParent(ASTNode)
	GetParent() ASTNode
	GetNodeType() string
	GetDataType() *common.DataType
	IsVariadic() bool
	Compile(c *Compiler)
	GetInstructions() []common.Instruction
	SetLhs()
}

type Node struct {
	Instructions []common.Instruction
	Parent       ASTNode
	NodeType     string
	DataType     *common.DataType
	Lhs          bool
	Variadic     bool
}

func (n *Node) GetInstructions() []common.Instruction {
	return n.Instructions
}

func (n *Node) SetParent(node ASTNode) {
	n.Parent = node
}

func (n *Node) GetParent() ASTNode {
	return n.Parent
}

func (n *Node) IsVariadic() bool {
	return n.Variadic
}

func (n *Node) GetNodeType() string {
	return n.NodeType
}

func (n *Node) GetDataType() *common.DataType {
	return n.DataType
}

func (n *Node) SetLhs() {
	n.Lhs = true
}
