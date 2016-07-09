package dot

import (
	"github.com/gonum/graph"
	"github.com/graphism/dot/ast"
)

// NewGraph returns a new dot AST graph based on the nodes and edges of the
// given gonum graph. The dot AST graph may later be manually annotated with
// additional attributes.
func NewGraph(src graph.Directed) *ast.Graph {
	dst := &ast.Graph{}
	nodes := src.Nodes()
	for _, n := range nodes {
		_ = n
	}
	panic("not yet implemented")
	return dst
}
