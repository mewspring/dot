package ast_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/graphism/dot"
	"github.com/graphism/dot/ast"
)

func TestParseFile(t *testing.T) {
	golden := []struct {
		in  string
		out string
	}{
		{in: "../internal/parser/testdata/empty.dot"},
		{in: "../internal/parser/testdata/graph.dot"},
		{in: "../internal/parser/testdata/digraph.dot"},
		{in: "../internal/parser/testdata/strict.dot"},
		{in: "../internal/parser/testdata/multi.dot"},
		{in: "../internal/parser/testdata/named_graph.dot"},
		{in: "../internal/parser/testdata/node_stmt.dot"},
		{in: "../internal/parser/testdata/edge_stmt.dot"},
		{in: "../internal/parser/testdata/attr_stmt.dot"},
		{in: "../internal/parser/testdata/attr.dot"},
		{
			in:  "../internal/parser/testdata/subgraph.dot",
			out: "../internal/parser/testdata/subgraph.golden",
		},
		{
			in:  "../internal/parser/testdata/semi.dot",
			out: "../internal/parser/testdata/semi.golden",
		},
		{
			in:  "../internal/parser/testdata/empty_attr.dot",
			out: "../internal/parser/testdata/empty_attr.golden",
		},
		{
			in:  "../internal/parser/testdata/attr_lists.dot",
			out: "../internal/parser/testdata/attr_lists.golden",
		},
		{
			in:  "../internal/parser/testdata/attr_sep.dot",
			out: "../internal/parser/testdata/attr_sep.golden",
		},
		{in: "../internal/parser/testdata/subgraph_vertex.dot"},
		{
			in:  "../internal/parser/testdata/port.dot",
			out: "../internal/parser/testdata/port.golden",
		},
	}
	for _, g := range golden {
		file, err := dot.ParseFile(g.in)
		if err != nil {
			t.Errorf("%q: unable to parse file; %v", g.in, err)
			continue
		}
		// If no output path is specified, the input is already golden.
		out := g.in
		if len(g.out) > 0 {
			out = g.out
		}
		buf, err := ioutil.ReadFile(out)
		if err != nil {
			t.Errorf("%q: unable to read file; %v", g.in, err)
			continue
		}
		got := file.String()
		// Remove trailing newline.
		want := string(bytes.TrimSpace(buf))
		if got != want {
			t.Errorf("%q: graph mismatch; expected %q, got %q", g.in, want, got)
		}
	}
}

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
	_ ast.Vertex = &ast.Node{}
	_ ast.Vertex = &ast.Subgraph{}
)
