package main

import (
	"github.com/nobishino/undirectedchan"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(undirectedchan.Analyzer) }

