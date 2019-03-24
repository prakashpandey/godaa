package main

import (
	"github.com/prakashpandey/godaa/markel-tree/transaction"
	"github.com/prakashpandey/godaa/markel-tree/tree"
)

func computeMarkelTree() {
	makeNodesEven()
	for i := 0; i < len(transactions)-2; i = i + 2 {
		xNode := generateNode(nil, nil, transactions[i])
		yNode := generateNode(nil, nil, transactions[i+1])
		parent := tree.Tree{
			Value:  xNode.Hash + yNode.Hash,
			Childs: []*tree.Tree{xNode, yNode},
			Parent: nil,
		}
		parent.SetHash()
	}
}

func generateNode(parent *tree.Tree, childs []*tree.Tree, transaction *transaction.Transaction) *tree.Tree {
	node := &tree.Tree{
		Hash:   "",
		Value:  transaction.Hash(),
		Parent: parent,
		Childs: childs,
	}
	return node
}

func makeNodesEven() {
	if len(transactions)%2 != 0 {
		addTransaction(transactions[len(transactions)-1])
	}
}
