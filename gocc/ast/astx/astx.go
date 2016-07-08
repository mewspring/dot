// Package astx implements utility functions for generating abstract syntax
// trees.
package astx

import "github.com/graphism/dot/gocc/ast"

// NewGraph returns a new graph based on the given graph strictness, direction,
// identifier and statements.
func NewGraph(strict, directed, id, stmts interface{}) (*ast.Graph, error) {
	panic("not yet implemented")
}

// NewID returns a new identifier based on the given ID token.
func NewID(id interface{}) (string, error) {
	panic("not yet implemented")
}
