package db

// node []byte format
// node[0:2] = type
// node[2:4] = number of keys
// node[4:(numKeys * 8)] = pointers
// node[:(numKeys * 2)] = offset
// after offset is the KVs of the data

//KV format
// n[0:2] key length
// n[2:4] value length
// n[4:(key length)] key
// n[klen: (value length)] value

type Node []byte

func init() {

}
