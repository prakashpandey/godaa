package transaction

// Detail ...
type Detail struct {
	Address string
	Amount  float64
	Type    Type
}

// Transaction ...
type Transaction struct {
	From      string
	To        string
	In        []Detail
	Out       []Detail
	Timestamp int64
}

// New ...
func New() *Transaction {
	return &Transaction{}
}

// Hash ...
func (t *Transaction) Hash() string {
	return ""
}
