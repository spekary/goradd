package node

// Code generated by goradd. DO NOT EDIT.

import (
	"encoding/gob"

	"github.com/goradd/goradd/pkg/orm/query"
)

type reverseNode struct {
	query.ReferenceNodeI
}

func Reverse() *reverseNode {
	n := reverseNode{
		query.NewTableNode("goraddUnit", "reverse", "Reverse"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

func (n *reverseNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	return nodes
}
func (n *reverseNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}
func (n *reverseNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}
func (n *reverseNode) Copy_() query.NodeI {
	return &reverseNode{query.CopyNode(n.ReferenceNodeI)}
}

// ID represents the id column in the database.
func (n *reverseNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goraddUnit",
		"reverse",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *reverseNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goraddUnit",
		"reverse",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ForwardCascades represents the many-to-one relationship formed by the reverse reference from the
// forward_cascades column in the reverse table.
func (n *reverseNode) ForwardCascades() *forwardCascadeNode {

	cn := &forwardCascadeNode{
		query.NewReverseReferenceNode(
			"goraddUnit",
			"reverse",
			"id",
			"forward_cascades",
			"ForwardCascades",
			"forward_cascade",
			"reverse_id",
			true,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

// ForwardCascadeUnique represents the one-to-one relationship formed by the reverse reference from the
// forward_cascade_unique column in the reverse table.
func (n *reverseNode) ForwardCascadeUnique() *forwardCascadeUniqueNode {

	cn := &forwardCascadeUniqueNode{
		query.NewReverseReferenceNode(
			"goraddUnit",
			"reverse",
			"id",
			"forward_cascade_unique",
			"ForwardCascadeUnique",
			"forward_cascade_unique",
			"reverse_id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

// ForwardNulls represents the many-to-one relationship formed by the reverse reference from the
// forward_nulls column in the reverse table.
func (n *reverseNode) ForwardNulls() *forwardNullNode {

	cn := &forwardNullNode{
		query.NewReverseReferenceNode(
			"goraddUnit",
			"reverse",
			"id",
			"forward_nulls",
			"ForwardNulls",
			"forward_null",
			"reverse_id",
			true,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

// ForwardNullUnique represents the one-to-one relationship formed by the reverse reference from the
// forward_null_unique column in the reverse table.
func (n *reverseNode) ForwardNullUnique() *forwardNullUniqueNode {

	cn := &forwardNullUniqueNode{
		query.NewReverseReferenceNode(
			"goraddUnit",
			"reverse",
			"id",
			"forward_null_unique",
			"ForwardNullUnique",
			"forward_null_unique",
			"reverse_id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

// ForwardRestricts represents the many-to-one relationship formed by the reverse reference from the
// forward_restricts column in the reverse table.
func (n *reverseNode) ForwardRestricts() *forwardRestrictNode {

	cn := &forwardRestrictNode{
		query.NewReverseReferenceNode(
			"goraddUnit",
			"reverse",
			"id",
			"forward_restricts",
			"ForwardRestricts",
			"forward_restrict",
			"reverse_id",
			true,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

// ForwardRestrictUnique represents the one-to-one relationship formed by the reverse reference from the
// forward_restrict_unique column in the reverse table.
func (n *reverseNode) ForwardRestrictUnique() *forwardRestrictUniqueNode {

	cn := &forwardRestrictUniqueNode{
		query.NewReverseReferenceNode(
			"goraddUnit",
			"reverse",
			"id",
			"forward_restrict_unique",
			"ForwardRestrictUnique",
			"forward_restrict_unique",
			"reverse_id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

func init() {
	gob.RegisterName("reverseNode2", &reverseNode{})
}
