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
		{in: "../gocc/parser/testdata/empty.dot"},
		{in: "../gocc/parser/testdata/graph.dot"},
		{in: "../gocc/parser/testdata/digraph.dot"},
		{in: "../gocc/parser/testdata/strict.dot"},
		{in: "../gocc/parser/testdata/multi.dot"},
		{in: "../gocc/parser/testdata/named_graph.dot"},
		{in: "../gocc/parser/testdata/node_stmt.dot"},
		{in: "../gocc/parser/testdata/edge_stmt.dot"},
		{in: "../gocc/parser/testdata/attr_stmt.dot"},
		{in: "../gocc/parser/testdata/attr.dot"},
		{
			in:  "../gocc/parser/testdata/subgraph.dot",
			out: "../gocc/parser/testdata/subgraph.golden",
		},
		{
			in:  "../gocc/parser/testdata/semi.dot",
			out: "../gocc/parser/testdata/semi.golden",
		},
		{
			in:  "../gocc/parser/testdata/empty_attr.dot",
			out: "../gocc/parser/testdata/empty_attr.golden",
		},
		{
			in:  "../gocc/parser/testdata/attr_lists.dot",
			out: "../gocc/parser/testdata/attr_lists.golden",
		},
		{
			in:  "../gocc/parser/testdata/attr_sep.dot",
			out: "../gocc/parser/testdata/attr_sep.golden",
		},
		{in: "../gocc/parser/testdata/subgraph_vertex.dot"},
		{
			in:  "../gocc/parser/testdata/port.dot",
			out: "../gocc/parser/testdata/port.golden",
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
	_ ast.Vertex = &ast.NodeID{}
	_ ast.Vertex = &ast.Subgraph{}
)
