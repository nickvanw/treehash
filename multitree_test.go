package treehash

import "testing"

func TestMultiTree(t *testing.T) {
	tt := []struct {
		nodes []string
		hash  string
	}{
		{
			nodes: []string{
				"6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918",
				"86e50149658661312a9e0b35558d84f6c6d3da797f552a9657fe0558ca40cdef",
				"7688b6ef52555962d008fff894223582c484517cea7da49ee67800adc7fc8866",
			},
			hash: "df427582902af3190c6f72877cc248d5b680b31c443341da68e283454047c4be",
		},
	}

	for _, v := range tt {
		var th MultiTreeHash
		for _, h := range v.nodes {
			th.Add(h)
		}
		if th.Hash() != v.hash {
			t.Fatalf("want: %q, got: %q", v.hash, th.Hash())
		}
	}

}
