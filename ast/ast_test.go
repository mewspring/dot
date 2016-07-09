package ast_test

import "github.com/graphism/dot/ast"

// Verify that all statements implement the Stmt interface.
var (
	_ ast.Stmt = &ast.NodeStmt{}
	_ ast.Stmt = &ast.EdgeStmt{}
	_ ast.Stmt = &ast.AttrStmt{}
	_ ast.Stmt = &ast.Attr{}
	_ ast.Stmt = &ast.Subgraph{}
)

// Verify that all vertices implement the Vertex interface.
var (
	_ ast.Vertex = &ast.NodeID{}
	_ ast.Vertex = &ast.Subgraph{}
)
