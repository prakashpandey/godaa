package main

import (
	"github.com/prakashpandey/godaa/markel-tree/transaction"
	"github.com/prakashpandey/godaa/markel-tree/tree"
)

var transactions []*transaction.Transaction

var markelTree *tree.Tree

func main() {
	generateTransactions()
	computeMarkelTree()
}
