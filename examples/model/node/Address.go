package node

// Code generated by goradd. DO NOT EDIT.

import (
	"github.com/goradd/goradd/pkg/orm/query"
)

type addressNode struct {
	query.NodeI
}

func Address() *addressNode {
	n := addressNode{
		query.NewTableNode("goradd", "address", "Address"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

func (n *addressNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.PersonID())
	nodes = append(nodes, n.Street())
	nodes = append(nodes, n.City())
	return nodes
}
func (n *addressNode) PrimaryKeyNode_() *query.ColumnNode {
	return n.ID()
}
func (n *addressNode) EmbeddedNode_() query.NodeI {
	return n.NodeI
}

func (n *addressNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"address",
		"id",
		"ID",
		query.ColTypeString,
	)
	query.SetParentNode(cn, n)
	return cn
}

func (n *addressNode) PersonID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"address",
		"person_id",
		"PersonID",
		query.ColTypeString,
	)
	query.SetParentNode(cn, n)
	return cn
}

func (n *addressNode) Person() *personNode {
	cn := &personNode{
		query.NewReferenceNode(
			"goradd",
			"address",
			"person_id",
			"PersonID",
			"Person",
			"person",
			"id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

func (n *addressNode) Street() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"address",
		"street",
		"Street",
		query.ColTypeString,
	)
	query.SetParentNode(cn, n)
	return cn
}

func (n *addressNode) City() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"address",
		"city",
		"City",
		query.ColTypeString,
	)
	query.SetParentNode(cn, n)
	return cn
}