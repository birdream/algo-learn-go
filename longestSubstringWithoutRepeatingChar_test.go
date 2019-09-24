package learn

import (
	"testing"
)

func TestLSWRC(t *testing.T) {
	t.Run("should correct", testLSWRC1)
	t.Run("should correct", testLSWRC2)
	t.Run("should correct", testLSWRC3)
	t.Run("should correct", testLSWRC4)
	t.Run("should correct", testLSWRC5)
}

func testLSWRC1(t *testing.T) {
	target := "aaaabkjhsdgfmmwwhh"
	res := methodLSWRC(target)
	if res != "abkjhsdgfm" {
		t.Error("Not correct")
	}
}

func testLSWRC2(t *testing.T) {
	target := "aaabbbcccddd"
	res := methodLSWRC(target)
	if res != "ab" {
		t.Errorf("Not correct %s", res)
	}
}

func testLSWRC3(t *testing.T) {
	target := "aaabbbcccd2dd"
	res := methodLSWRC(target)
	if res != "cd2" {
		t.Errorf("Not correct %s", res)
	}
}

func testLSWRC4(t *testing.T) {
	target := `aaabbbcc|\cd2dd`
	res := methodLSWRC(target)
	if res != `|\cd2` {
		t.Errorf("Not correct %s", res)
	}
}

func testLSWRC5(t *testing.T) {
	target := `ccccccccccc`
	res := methodLSWRC(target)
	if res != `c` {
		t.Errorf("Not correct %s", res)
	}
}
