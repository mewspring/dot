# NOTICE: New project home!

The DOT parser has been merged with the gonum graph project (see [graphism/dot#8](https://github.com/graphism/dot/issues/8)).

For active development, the new home of `graphism/dot` is the [gonum/graph](https://github.com/gonum/graph) repository, under the import path `formats/dot`.

```diff
-import "github.com/graphism/dot"
+import "github.com/gonum/graph/formats/dot"
```

# DOT

[![Build Status](https://travis-ci.org/graphism/dot.svg?branch=master)](https://travis-ci.org/graphism/dot)
[![Coverage Status](https://coveralls.io/repos/github/graphism/dot/badge.svg?branch=master)](https://coveralls.io/github/graphism/dot?branch=master)
[![GoDoc](https://godoc.org/github.com/graphism/dot?status.svg)](https://godoc.org/github.com/graphism/dot)

This package provides access to [Graphviz DOT] files.

[Graphviz DOT]: http://www.graphviz.org/doc/info/lang.html

## Installation

```
$ go get github.com/graphism/dot/...
```

## Credits

This project has been inspired by [Walter Schulze](https://github.com/awalterschulze)'s [gographviz](https://github.com/awalterschulze/gographviz) library, and also uses [Marius Ackerman](https://github.com/goccmack) and Walter's [Gocc](https://github.com/goccmack/gocc) compiler kit to generate lexers and parsers from a [BNF grammar](https://github.com/graphism/dot/blob/master/internal/dot.bnf) of the DOT file format.

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
