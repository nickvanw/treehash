package treehash

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

const (
	megabyte = 1 << 20
)

// TreeHash implments ReadCloser, data written to is is internally hashed
// into a hash tree as documented here:
// http://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html
type TreeHash struct {
	nodes [][sha256.Size]byte
	hash  []byte // the final hash

	data []byte

	sha hash.Hash

	//config
	blockSize int
}

// NewTreeHash creates and initializes a new TreeHash using the default values
// The tree sum is only available after calling close, to ensure the last block
// of data is written to the tree properly
func NewTreeHash(cfgs ...confFunc) *TreeHash {
	t := &TreeHash{
		nodes:     [][sha256.Size]byte{},
		data:      []byte{},
		sha:       sha256.New(),
		blockSize: megabyte,
	}
	for _, v := range cfgs {
		v(t)
	}
	return t
}

// Write hashes the data in p and adds it to the internal sha256 tree
func (t *TreeHash) Write(p []byte) (n int, err error) {
	n = len(p)
	t.data = append(t.data, p...)

	for len(t.data) >= t.blockSize {
		t.nodes = append(t.nodes, sha256.Sum256(t.data[:t.blockSize]))
		t.sha.Write(t.data[:t.blockSize])
		t.data = t.data[t.blockSize:]
	}

	return n, nil
}

// Close closes the writer and computes the final tree hash.
// After the first call, any writes will create an invalid hash tree
func (t *TreeHash) Close() error {
	// The last piece of data
	if len(t.data) > 0 {
		t.nodes = append(t.nodes, sha256.Sum256(t.data))
		t.sha.Write(t.data)
		t.data = []byte{} // zero the data out
	}
	t.hash = compute(t.nodes)

	return nil
}

// TreeHash returns the string of the treehash. This will only return
// a hash after Close has been called
func (t *TreeHash) TreeHash() string {
	return hex.EncodeToString(t.hash)
}

// Hash returns the sha256 sum of all of the data written to the reader.
// This will only be valid after Close has been called
func (t *TreeHash) Hash() string {
	hsh := t.sha.Sum(nil)
	return hex.EncodeToString(hsh)
}

func compute(d [][sha256.Size]byte) []byte {
	// create a copy of the data
	prev := make([][sha256.Size]byte, len(d))
	copy(prev, d)

	// while we still have multiple hashes present in the "tree"
	for len(prev) > 1 {
		// create the next level up in the "tree"
		next := make([][sha256.Size]byte, 0)
		for i := 0; i < len(prev); i = i + 2 {
			// if there are at least two remaining
			if len(prev)-i > 1 {
				sum := sha256.Sum256(append(prev[i][:], prev[i+1][:]...))
				next = append(next, sum)
			} else {
				// attach the odd chunk on
				next = append(next, prev[i])
			}
		}
		prev = next
	}
	return prev[0][:]
}
