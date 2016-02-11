package misc

import "testing"

func TestMultiGCD(t *testing.T) {
	line := []int64{16, 12, 4, 28}

	if MultiGCD(line) != 4 {
		t.Error("failed tocalculate the GCD")
	}
}
