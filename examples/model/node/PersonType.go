package node

// Code generated by goradd. DO NOT EDIT.

import (
	"github.com/goradd/goradd/pkg/orm/query"
)

type personTypeNode struct {
	query.NodeI
}


func (n *personTypeNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	return nodes
}

func (n *personTypeNode) PrimaryKeyNode_() (*query.ColumnNode) {
	return n.ID()
}

func (n *personTypeNode) EmbeddedNode_() query.NodeI {
	return n.NodeI
}

func (n *personTypeNode) ID() *query.ColumnNode {

	cn := query.NewColumnNode (
		"goradd",
		"person_type",
		"id",
		"ID",
		query.ColTypeUnsigned,
	)
	query.SetParentNode(cn, n)
	return cn
}
func (n *personTypeNode) Name() *query.ColumnNode {

	cn := query.NewColumnNode (
		"goradd",
		"person_type",
		"name",
		"Name",
		query.ColTypeString,
	)
	query.SetParentNode(cn, n)
	return cn
}