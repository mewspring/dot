// Package astx implements utility functions for generating abstract syntax
// trees of Graphviz DOT graphs.
package astx

import "github.com/graphism/dot/gocc/ast"

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
// component and attributes.
func NewAttrStmt(component, attrs interface{}) (*ast.AttrStmt, error) {
	panic("astx.NewAttrStmt: not yet implemented")
}

// TODO: Add AttrList.

// --- [ Attribute ] -----------------------------------------------------------

// NewAttr returns a new attribute based on the given key-value pair.
func NewAttr(key, value interface{}) (*ast.Attr, error) {
	panic("astx.NewAttr: not yet implemented")
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
	panic("astx.NewID: not yet implemented")
}
