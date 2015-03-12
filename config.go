package treehash

type confFunc func(t *TreeHash)

// BlockSize is used to set the block size from the default of 1MB
func BlockSize(size int) confFunc {
	return func(t *TreeHash) {
		t.blockSize = size
	}
}
