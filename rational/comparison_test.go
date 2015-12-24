package rational

import "testing"

func TestGreaterThan(t *testing.T) {
	r1 := New(3, 4)
	r2 := New(5, 7)

	if !r1.GreaterThan(r2) || r1.LessThan(r2) {
		t.Error("failed to compare if rational is greater than another rational")
	}
}

func TestGreaterThanNum(t *testing.T) {
	r1 := New(3, 4)

	if !r1.GreaterThanNum(-1) || r1.LessThanNum(-1) {
		t.Error("failed to compare if rational is greater than integer")
	}
}

func TestLessThan(t *testing.T) {
	r1 := New(5, 7)
	r2 := New(3, 4)

	if !r1.LessThan(r2) || r1.GreaterThan(r2) {
		t.Error("failed to compare if rational is less than another rational")
	}
}

func TestLessThanNum(t *testing.T) {
	r1 := New(15, 4)

	if !r1.GreaterThanNum(2) || r1.LessThanNum(2) {
		t.Error("failed to compare if rational is greater than integer")
	}
}
