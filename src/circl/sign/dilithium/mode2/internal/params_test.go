package internal

import (
	"testing"

	"circl/sign/dilithium/internal/common"
)

// Tests specific to the current mode

func TestVectorDeriveUniformLeqEta(t *testing.T) {
	var p common.Poly
	var seed [32]byte
	p2 := common.Poly{
		2, 8380416, 1, 2, 8380411, 8380413, 1, 1, 8380414, 8380415,
		8380415, 0, 6, 3, 6, 1, 8380413, 8380411, 8380416, 2,
		8380412, 3, 8380413, 8380412, 8380411, 6, 2, 8380411, 2,
		8380413, 6, 8380414, 0, 8380415, 6, 8380413, 1, 3, 6,
		8380415, 3, 4, 4, 1, 0, 6, 2, 3, 8380415, 8380414, 1, 1,
		6, 4, 8380416, 8380415, 6, 8380411, 4, 4, 3, 0, 3, 1,
		8380413, 8380415, 2, 3, 8380413, 8380413, 6, 8380414,
		8380411, 5, 8380414, 8380412, 8380416, 8380415, 8380411,
		5, 8380414, 8380412, 6, 8380414, 2, 0, 8380416, 8380413,
		3, 8380411, 4, 5, 8380416, 0, 8380412, 5, 8380416, 8380414,
		2, 1, 2, 5, 8380416, 8380416, 5, 8380413, 8380416, 5,
		8380411, 8380411, 4, 6, 1, 8380412, 8380416, 8380416, 3,
		8380416, 4, 8380415, 0, 4, 6, 6, 1, 8380416, 8380411, 5,
		4, 3, 8380416, 8380416, 6, 8380413, 8380415, 3, 8380416,
		8380412, 8380413, 2, 8380412, 8380413, 8380414, 6, 8380413,
		4, 5, 5, 8380414, 3, 8380416, 8380416, 8380415, 8380414,
		8380414, 8380414, 3, 8380414, 8380411, 5, 0, 3, 0, 1,
		8380416, 8380415, 3, 8380413, 8380415, 5, 4, 0, 2, 8380413,
		3, 8380413, 8380413, 2, 1, 1, 8380412, 1, 8380411, 8380416,
		8380415, 8380414, 8380413, 8380416, 2, 8380412, 2, 6,
		8380413, 8380411, 0, 3, 1, 8380414, 1, 3, 8380412, 3, 3,
		8380416, 2, 4, 8380412, 8380413, 8380414, 3, 0, 6, 8380412,
		8380415, 8380416, 5, 5, 0, 6, 8380415, 8380413, 8380414,
		5, 6, 3, 6, 3, 5, 5, 8380415, 6, 4, 8380416, 3, 2, 6,
		8380414, 2, 8380415, 6, 8380412, 8380413, 8380415, 1, 2,
		0, 8380411, 6, 8380413, 6, 8380414, 8380416, 8380415,
		8380414, 1, 8380412,
	}
	for i := 0; i < 32; i++ {
		seed[i] = byte(i)
	}
	PolyDeriveUniformLeqEta(&p, &seed, 30000)
	p.Normalize()
	if p != p2 {
		t.Fatalf("%v != %v", p, p2)
	}
}
