package db

import "encoding/binary"

// Branch format
// node[0:2] Number of Kids
// node[2:] Pointers to kids

// Leaf format
// node[0:2] Number Of Keys
// node[2: 2 *NKeys] Offset
// node[:] Key / Values

type Leaf []byte

type Branch []byte

func NewLeaf() Leaf {
	return Leaf{}
}

func NewBranch() Branch {
	return Branch{}
}

func (L Leaf) GetNKeys() uint16 {
	return binary.LittleEndian.Uint16(L[0:2])
}

func (B Branch) GetNChildren() uint16 {
	return binary.LittleEndian.Uint16(B[0:2])
}

func (L Leaf) AddNewKey() {
	binary.LittleEndian.PutUint16(L, (L.GetNKeys() + 1))
}

func (L Leaf) GetKey() []byte {
	return
}
