package astx_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/graphism/dot"
)

func TestParseFile(t *testing.T) {
	golden := []struct {
		in  string
		out string
	}{
		{in: "../parser/testdata/empty.dot"},
		{in: "../parser/testdata/graph.dot"},
		{in: "../parser/testdata/digraph.dot"},
		{in: "../parser/testdata/strict.dot"},
		{in: "../parser/testdata/multi.dot"},
		{in: "../parser/testdata/named_graph.dot"},
		{in: "../parser/testdata/node_stmt.dot"},
		{in: "../parser/testdata/edge_stmt.dot"},
		{in: "../parser/testdata/attr_stmt.dot"},
		{in: "../parser/testdata/attr.dot"},
		{
			in:  "../parser/testdata/subgraph.dot",
			out: "../parser/testdata/subgraph.golden",
		},
		{
			in:  "../parser/testdata/semi.dot",
			out: "../parser/testdata/semi.golden",
		},
		{
			in:  "../parser/testdata/empty_attr.dot",
			out: "../parser/testdata/empty_attr.golden",
		},
		{
			in:  "../parser/testdata/attr_lists.dot",
			out: "../parser/testdata/attr_lists.golden",
		},
		{
			in:  "../parser/testdata/attr_sep.dot",
			out: "../parser/testdata/attr_sep.golden",
		},
		{in: "../parser/testdata/subgraph_vertex.dot"},
		{
			in:  "../parser/testdata/port.dot",
			out: "../parser/testdata/port.golden",
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
