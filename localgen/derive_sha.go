
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


func DeriveSha(list DerivableList) []byte {
	keybuf := new(bytes.Buffer)
	trie, _ := trie.NewTrie(trie.Hash(stringtohash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")), __v)
	for i := 0; i < list.Len(); i++ {
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i))
		trie.Update(keybuf.Bytes(), list.GetRlp(i))
	}
	nh, _ := trie.Commit(nil)
	return nh.Bytes()
}
