// Package astx implements utility functions for generating abstract syntax
// trees of Graphviz DOT graphs.
package astx

import (
	"github.com/graphism/dot/gocc/ast"
	"github.com/graphism/dot/gocc/token"
	"github.com/mewkiz/pkg/errutil"
)

// === [ Graphs ] ==============================================================

// NewGraph returns a new graph based on the given graph strictness, direction,
// ID and statements.
func NewGraph(strict, directed, id, stmts interface{}) (*ast.Graph, error) {
	panic("astx.NewGraph: not yet implemented")
}

// === [ Statements ] ==========================================================

// NewStmtList returns a new statement list based on the given statement.
func NewStmtList(stmt interface{}) ([]ast.Stmt, error) {
	panic("astx.NewStmtList: not yet implemented")
}

// AppendStmt appends stmt to the given statement list.
func AppendStmt(list, stmt interface{}) ([]ast.Stmt, error) {
	panic("astx.AppendStmt: not yet implemented")
}

// --- [ Node statement ] ------------------------------------------------------

// NewNodeStmt returns a new node statement based on the given node ID and
// attributes.
func NewNodeStmt(nodeID, attrs interface{}) (*ast.NodeStmt, error) {
	panic("astx.NewNodeStmt: not yet implemented")
}

// --- [ Edge statement ] ------------------------------------------------------

// NewEdgeStmt returns a new edge statement based on the given source vertex,
// outgoing edge and attributes.
func NewEdgeStmt(from, to, attrs interface{}) (*ast.EdgeStmt, error) {
	panic("astx.NewEdgeStmt: not yet implemented")
}

// NewEdge returns a new edge based on the given edge direction, vertex and
// outgoing edge.
func NewEdge(directed, vertex, to interface{}) (*ast.EdgeStmt, error) {
	panic("astx.NewEdge: not yet implemented")
}

// --- [ Attribute statement ] -------------------------------------------------

// NewAttrStmt returns a new attribute statement based on the given graph
// component kind and attributes.
func NewAttrStmt(kind, attrs interface{}) (*ast.AttrStmt, error) {
	k, ok := kind.(ast.Kind)
	if !ok {
		return nil, errutil.Newf("invalid graph component kind type; expected ast.Kind, got %T", kind)
	}
	a, ok := attrs.([]*ast.Attr)
	if !ok {
		return nil, errutil.Newf("invalid attributes type; expected []*ast.Attr, got %T", attrs)
	}
	return &ast.AttrStmt{Kind: k, Attrs: a}, nil
}

// NewAttrList returns a new attribute list based on the given attribute.
func NewAttrList(attr interface{}) ([]*ast.Attr, error) {
	a, ok := attr.(*ast.Attr)
	if !ok {
		return nil, errutil.Newf("invalid attribute type; expected *ast.Attr, got %T", attr)
	}
	return []*ast.Attr{a}, nil
}

// AppendAttr appends attr to the given attribute list.
func AppendAttr(list, attr interface{}) ([]*ast.Attr, error) {
	l, ok := list.([]*ast.Attr)
	if !ok {
		return nil, errutil.Newf("invalid attribute list type; expected []*ast.Attr, got %T", list)
	}
	a, ok := attr.(*ast.Attr)
	if !ok {
		return nil, errutil.Newf("invalid attribute type; expected *ast.Attr, got %T", attr)
	}
	return append(l, a), nil
}

// AppendAttrList appends attrs to the given attribute list.
func AppendAttrList(list, attrs interface{}) ([]*ast.Attr, error) {
	l, ok := list.([]*ast.Attr)
	if list != nil && !ok {
		return nil, errutil.Newf("invalid attribute list type; expected []*ast.Attr, got %T", list)
	}
	a, ok := attrs.([]*ast.Attr)
	if attrs != nil && !ok {
		return nil, errutil.Newf("invalid attributes type; expected []*ast.Attr, got %T", attrs)
	}
	return append(l, a...), nil
}

// --- [ Attribute ] -----------------------------------------------------------

// NewAttr returns a new attribute based on the given key-value pair.
func NewAttr(key, val interface{}) (*ast.Attr, error) {
	k, ok := key.(string)
	if !ok {
		return nil, errutil.Newf("invalid key type; expected string, got %T", key)
	}
	v, ok := val.(string)
	if !ok {
		return nil, errutil.Newf("invalid value type; expected string, got %T", val)
	}
	return &ast.Attr{Key: k, Val: v}, nil
}

// --- [ Subgraph ] ------------------------------------------------------------

// NewSubgraph returns a new subgraph based on the given subgraph ID and
// statements.
func NewSubgraph(id, stmts interface{}) (*ast.Subgraph, error) {
	panic("astx.NewSubgraph: not yet implemented")
}

// === [ Vertices ] ============================================================

// --- [ Node identifier ] -----------------------------------------------------

// NewNodeID returns a new node ID based on the given node id and port.
func NewNodeID(id, port interface{}) (*ast.NodeID, error) {
	panic("astx.NewNodeID: not yet implemented")
}

// NewPort returns a new port based on the given id and compass point.
func NewPort(id, compassPoint interface{}) (*ast.Port, error) {
	// NOTE: If compassPoint is nil, id may be either an identifier or a compass
	// point.
	//
	// The following strings are valid compass points:
	//
	//    "n", "ne", "e", "se", "s", "sw", "w", "nw", "c" and "_"
	panic("astx.n: not yet implemented")
}

// === [ Identifiers ] =========================================================

// NewID returns a new identifier based on the given ID token.
func NewID(id interface{}) (string, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return "", errutil.Newf("invalid identifier type; expected *token.Token, got %T", id)
	}
	return string(i.Lit), nil
}
