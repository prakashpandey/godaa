package tree

// Tree ...
type Tree struct {
	Hash string
	// Value  *transaction.Transaction
	Value  string
	Parent *Tree
	Childs []*Tree
}

// SetHash ...
func (t *Tree) SetHash() string {
	return ""
}
