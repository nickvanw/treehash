package treehash

import (
	"crypto/sha256"
	"encoding/hex"
)

type MultiTreeHash struct {
	nodes [][sha256.Size]byte
}

// Add adds another hash to the tree
func (m *MultiTreeHash) Add(hsh string) {
	var b [sha256.Size]byte
	hex.Decode(b[:], []byte(hsh))
	m.nodes = append(m.nodes, b)
}

// Hash returns the current hash based on the tree
func (m *MultiTreeHash) Hash() string {
	hsh := compute(m.nodes)
	return hex.EncodeToString(hsh)
}
