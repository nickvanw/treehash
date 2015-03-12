package treehash

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tt := []struct {
		blockSize int
		data      []byte
		nodes     []string
		left      []byte
		treeSum   string
		totalSum  string
	}{
		{
			blockSize: 2,
			data:      []byte(`12345`),
			nodes: []string{
				"6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918",
				"86e50149658661312a9e0b35558d84f6c6d3da797f552a9657fe0558ca40cdef",
			},
			left:     []byte("5"),
			treeSum:  "f8b07479ea9cca853be2c6c7a5ee93bab2f0efe955326e92d5ee83f1167fe06c",
			totalSum: "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5",
		},
	}

	for _, v := range tt {
		tr := NewTreeHash(BlockSize(v.blockSize))
		tr.Write(v.data)
		if len(tr.nodes) != len(v.nodes) {
			t.Fatalf("wanted %v nodes, got: %v", len(v.nodes), len(tr.nodes))
		}
		for i, n := range tr.nodes {
			if hex.EncodeToString(n[:]) != v.nodes[i] {
				t.Logf("incorrect node number: %v", i)
				t.Logf("got: %q", hex.EncodeToString(n[:]))
				t.Logf("want: %q", v.nodes[i])
				t.Fatalf("sha256 node incorrect")
			}
		}
		if !reflect.DeepEqual(v.left, tr.data) {
			t.Logf("want: %#v", v.left)
			t.Logf("got: %#v", tr.data)
			t.Fatalf("data left in buffer was incorrect")
		}
		tr.Close()
		if tr.TreeHash() != v.treeSum {
			t.Fatalf("want tree hash: %q, got: %q", v.treeSum, tr.TreeHash())
		}
		if tr.Hash() != v.totalSum {
			t.Fatalf("want hash: %q, got: %q", v.totalSum, tr.Hash())
		}
	}
}
