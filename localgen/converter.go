package prover

// unint64(8-bytes) to bytes
func uint64tobytes(number uint64) []byte {
	enc := make([]byte, 8)
	for idx := uint64(0); idx < 8; idx++ {
		enc[7 - idx] = byte((number >> (idx << 3) ) & 0xff)
	}
	//fmt.Println(enc);
	return enc
}

//input a string and return common.Hash
func stringtohash(hashstr string) Hash {
	var hres Hash
	ofs := uint64(0)

	if hashstr[1] == 'x' {
		ofs = 1
	}

	for idx	:= uint64(0); idx < 32; idx++ {
		hres[idx] |= byte(hexmaps[hashstr[ (idx + ofs) << 1     ]] << 4)
		hres[idx] |= byte(hexmaps[hashstr[((idx + ofs) << 1) | 1]])
	}
	return hres
}

//input a string and return a byte slice
func stringtobytes(bytes string) []byte {
	glen := len(bytes)
	if  glen <= 1 || ((glen & 1) == 1) {
		return nil
	}

	glen >>= 1
	ofs := 0

	if bytes[1] == 'x' {
		ofs = 1
	}

	glen -= ofs

	bres := make([]byte, glen, glen)
	for idx := 0; idx < glen; idx++ {
		bres[idx] |= byte(hexmaps[bytes[(idx + ofs) << 1 ]] << 4)
		bres[idx] |= byte(hexmaps[bytes[(idx + ofs) << 1 | 1]])
	}
	return bres
}

//input a string and return a nibble slice
func stringtonibbles(nibbles string) []byte {
	glen := len(nibbles)

	ofs := 0
	if nibbles[1] == 'x' {
		ofs = 2
	}
	glen -= ofs

	bres := make([]byte, glen, glen)
	for idx := 0; idx < glen; idx++ {
		bres[idx] = byte(hexmaps[nibbles[idx + ofs]])
	}
	return bres
}

// input a block number and the corresponding block hash, return a headerkey
func stringtoheaderkey(number uint64, hashstr string) []byte {
	return append(append(headerPrefix, uint64tobytes(number)...), stringtohash(hashstr).bytes()...);
}
// blockReceiptsKey = blockReceiptsPrefix + num (uint64 big endian) + hash
func stringtoreceiptkey(number uint64, hashstr string) []byte {
	return append(append(blockReceiptsPrefix, uint64tobytes(number)...), stringtohash(hashstr).bytes()...)
}
// blockReceiptsKey = blockReceiptsPrefix + num (uint64 big endian) + hash
func stringtobodykey(number uint64, hashstr string) []byte {
	return append(append(blockBodyPrefix, uint64tobytes(number)...), stringtohash(hashstr).bytes()...)
}
func stringtoindexkey(hashstr string) []byte {
	return append(txLookupPrefix, stringtohash(hashstr).bytes()...)
}

func init() {
	for idx := '0'; idx <= '9'; idx++ {
		hexmaps[idx] = uint64(idx - '0')
	}
	for idx := 'a'; idx <= 'f'; idx++ {
		hexmaps[idx] = uint64(idx - 'a' + 10)
	}
	for idx := 'A'; idx <= 'F'; idx++ {
		hexmaps[idx] = uint64(idx - 'A' + 10)
	}
}
