package learn

import (
	"testing"
)

func TestForTest(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		hello()
	})
}
