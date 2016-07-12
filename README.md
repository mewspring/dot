# DOT

[![Build Status](https://travis-ci.org/graphism/dot.svg?branch=master)](https://travis-ci.org/graphism/dot)
[![Coverage Status](https://coveralls.io/repos/github/graphism/dot/badge.svg?branch=master)](https://coveralls.io/github/graphism/dot?branch=master)
[![GoDoc](https://godoc.org/github.com/graphism/dot?status.svg)](https://godoc.org/github.com/graphism/dot)

This package provides access to [Graphviz DOT] files.

[Graphviz DOT]: http://www.graphviz.org/doc/info/lang.html

## Installation

Install [Gocc] `go get github.com/goccmack/gocc`.

[Gocc]: https://github.com/goccmack/gocc

```
$ go get -d github.com/graphism/dot
$ cd ${GOPATH}/src/github.com/graphism/dot/internal
$ make gen
$ go get github.com/graphism/dot/...
```

## Credits

This project has been inspired by [Walter Schulze](https://github.com/awalterschulze)'s [gographviz](https://github.com/awalterschulze/gographviz) library, and also uses [Marius Ackerman](https://github.com/goccmack) and Walter's [Gocc](https://github.com/goccmack/gocc) compiler kit to generate lexers and parsers from a [BNF grammar](https://github.com/graphism/dot/blob/master/gocc/dot.bnf) of the DOT file format.

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
