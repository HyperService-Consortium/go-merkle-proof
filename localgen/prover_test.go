//link to Goleveldb
package prover

import (
	"testing"
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"github.com/Myriad-Dreamin/go-rlp"
	"encoding/hex"
)

const (
	EDB_PATH = "D:/Go Ethereum/data/geth/chaindata"
)


// TODO - iterator
func checkAll(db *leveldb.DB){
	iter := db.NewIterator(nil, nil)
	for iter.Next(){
		key   := iter.Key()
		value := iter.Value()
		fmt.Println([]byte(key),[]byte(value))
	}
}


type transactionsEncoding struct {
	Txs Transactions
	IdleBytes []byte
}


func TestTxTrieGen(t *testing.T) {
	var bn uint64 = 1069
	rootHashStr := "0x4a857b36196d7c7c51466b12297192ea68179b8fc40b751fadfd2e1e5cf57ede"
	transactionHash := "be6f6acaec1314776bf9117cfcd4e316e00c176cc249a63bc9667d083ab32849"

	db, err := leveldb.OpenFile(EDB_PATH, nil)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	querynode, err := db.Get(stringtobodykey(bn, rootHashStr), nil)
	if err != nil {
		t.Error(err)
	}
	var nodeDecode transactionsEncoding
	// fmt.Println(querynode)
	err = rlp.DecodeBytes(querynode, &nodeDecode)
	if hex.EncodeToString(DeriveSha(nodeDecode.Txs)) != transactionHash {
		t.Errorf("generate error")
	}
}
