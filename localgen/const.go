package prover

var (
	//Got LastHeader
	headHeaderKey = []byte("LastHeader")
	//Got LastBlock
	headBlockKey = []byte("LastBlock")
	//Got a BlockHeader
	headerPrefix = []byte("h")
	//Got index of transactions from db
	txLookupPrefix = []byte("l") 
	//Got receipts
	blockReceiptsPrefix = []byte("r")
	//Got BlockBody
	blockBodyPrefix = []byte("b")
)
