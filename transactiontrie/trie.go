
package prover

import (
	"bytes"
	"github.com/Myriad-Dreamin/go-mpt"
	"github.com/Myriad-Dreamin/go-rlp"
	"github.com/syndtr/goleveldb/leveldb"
)


type DerivableList interface {
	Len() int
	GetRlp(i int) []byte
}

type voidDB struct {

}

func (vdb *voidDB) Get([]byte) ([]byte, error) {
	return nil, nil
}

var __op, _ = leveldb.OpenFile("./testdb", nil)
var __v, _ = trie.NewNodeBasefromDB(__op)
var emptyHash = trie.Hash(stringtohash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"))


func NewVoidTrie() (*trie.Trie, error) {
	return trie.NewTrie(emptyHash, __v)
}


func NewTxTrie(list DerivableList) (*trie.Trie, error) {
	keybuf := new(bytes.Buffer)
	txTrie, err := NewVoidTrie()


	if err != nil {
		return nil, err
	}
	for i := 0; i < list.Len(); i++ {
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i))
		txTrie.Update(keybuf.Bytes(), list.GetRlp(i))
	}
	return txTrie, nil
}



type transactionsEncoding struct {
	Txs Transactions
	IdleInterface interface{}
}

func GenerateTxTrieFromLocal(filepath string, blockNumber uint64, rootHashStr string) (*trie.Trie, error) {
	db, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	querynode, err := db.Get(stringtobodykey(blockNumber, rootHashStr), nil)
	if err != nil {
		return nil, err
	}
	var nodeDecode transactionsEncoding
	err = rlp.DecodeBytes(querynode, &nodeDecode)
	if err != nil {
		return nil, err
	}
	return NewTxTrie(nodeDecode.Txs)
}

