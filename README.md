# treehash
--
    import "github.com/nickvanw/treehash"

Package treehash implements the sha256 tree hash for Amazon Glacier

## Usage

#### func  BlockSize

```go
func BlockSize(size int) confFunc
```
BlockSize is used to set the block size from the default of 1MB

#### type MultiTreeHash

```go
type MultiTreeHash struct {
}
```

MultiTreeHash allows the cumulative sum of multiple TreeHashes to one master
hash, for multipart uploads with Amazon Glacier

#### func (*MultiTreeHash) Add

```go
func (m *MultiTreeHash) Add(hsh string)
```
Add adds another hash to the tree

#### func (*MultiTreeHash) Hash

```go
func (m *MultiTreeHash) Hash() string
```
Hash returns the current hash based on the tree

#### type TreeHash

```go
type TreeHash struct {
}
```

TreeHash implments ReadCloser, data written to is is internally hashed into a
hash tree as documented here:
http://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html

#### func  NewTreeHash

```go
func NewTreeHash(cfgs ...confFunc) *TreeHash
```
NewTreeHash creates and initializes a new TreeHash using the default values The
tree sum is only available after calling close, to ensure the last block of data
is written to the tree properly

#### func (*TreeHash) Close

```go
func (t *TreeHash) Close() error
```
Close closes the writer and computes the final tree hash After the first call,
any writes will create an invalid hash tree

#### func (*TreeHash) Hash

```go
func (t *TreeHash) Hash() string
```
Hash returns the sha256 sum of all of the data written to the reader This will
only be valid after Close has been called

#### func (*TreeHash) TreeHash

```go
func (t *TreeHash) TreeHash() string
```
TreeHash returns the string of the treehash. This will only return a hash after
Close has been called

#### func (*TreeHash) Write

```go
func (t *TreeHash) Write(p []byte) (n int, err error)
```
Write hashes the data in p and adds it to the internal sha256 tree
