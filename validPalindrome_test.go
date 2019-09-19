package learn

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	t.Run("should work", testValidPalindrome)
}

func testValidPalindrome(t *testing.T) {
	if !validPalindrome("abcba") {
		t.Error("Not correct")
	}

	if !validPalindrome("9111, 2 221119") {
		t.Error("Not correct")
	}

	if !validPalindrome("Fa9111 2z2z21;119 // }Af") {
		t.Error("Not correct")
	}

	if validPalindrome("Abcbaa") {
		t.Error("Not correct")
	}
}
