package prover

const (
	// Hash byte-Array's length
	HashLength = 32
	// Hash hesgtring's length
    HASHSTRINGLENGTH = 64
)

// char maps to bit integer
var hexmaps [128]uint64

// common.Hash
type Hash [HashLength]byte

func (h Hash) bytes() []byte { return h[ : ] }
