package parser_test

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
		{in: "../../testdata/golden/empty.dot"},
		{in: "../../testdata/golden/graph.dot"},
		{in: "../../testdata/golden/digraph.dot"},
		{in: "../../testdata/golden/strict.dot"},
		{in: "../../testdata/golden/multi.dot"},
		{in: "../../testdata/golden/named_graph.dot"},
		{in: "../../testdata/golden/node_stmt.dot"},
		{in: "../../testdata/golden/edge_stmt.dot"},
		{in: "../../testdata/golden/attr_stmt.dot"},
		{in: "../../testdata/golden/attr.dot"},
		{
			in:  "../../testdata/golden/subgraph.dot",
			out: "../../testdata/golden/subgraph.golden",
		},
		{
			in:  "../../testdata/golden/semi.dot",
			out: "../../testdata/golden/semi.golden",
		},
		{
			in:  "../../testdata/golden/empty_attr.dot",
			out: "../../testdata/golden/empty_attr.golden",
		},
		{
			in:  "../../testdata/golden/attr_lists.dot",
			out: "../../testdata/golden/attr_lists.golden",
		},
		{
			in:  "../../testdata/golden/attr_sep.dot",
			out: "../../testdata/golden/attr_sep.golden",
		},
		{in: "../../testdata/golden/subgraph_vertex.dot"},
		{
			in:  "../../testdata/golden/port.dot",
			out: "../../testdata/golden/port.golden",
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
