//link to Goleveldb
package prover

import (
	"testing"
	"fmt"
	"encoding/hex"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/Myriad-Dreamin/go-mpt"
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

func TestTxTrieGen(t *testing.T) {
	var bn uint64 = 1069
	rootHashStr := "0x4a857b36196d7c7c51466b12297192ea68179b8fc40b751fadfd2e1e5cf57ede"
	transactionHash := "be6f6acaec1314776bf9117cfcd4e316e00c176cc249a63bc9667d083ab32849"
	filepath := EDB_PATH

	txTrie, err := GenerateTxTrieFromLocal(filepath, bn, rootHashStr)
	if err != nil {
		t.Error(err)
		return
	}
	var bt trie.Hash
	bt, err = txTrie.Commit(nil)
	if err != nil {
		t.Error(err)
	}
	if hex.EncodeToString(bt.Bytes()) != transactionHash {
		t.Errorf("generate error")
	}
}
