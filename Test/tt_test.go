package learn

import (
	"fmt"
	"testing"
)

func TestForTest(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		hello()
		keyMap := make(map[string]int)
		fmt.Println(keyMap["a"])
		a, b := keyMap["a"]
		fmt.Println(a)
		fmt.Println(b)

		keyMap["a"] = 1
		a1, b1 := keyMap["a"]
		fmt.Println(a1)
		fmt.Println(b1)
	})
}
