package rational

import "testing"

func TestNewFromFloat(t *testing.T) {
	r, err := NewFromFloat(1.25)
	if err != nil {
		t.Error(err)
	}

	if r != New(5, 4) {
		t.Error("failed to convert float to rational")
	}
}

func TestDivide(t *testing.T) {
	r1 := New(3, 4)
	r2 := New(5, 7)

	if r1.Divide(r2) != New(21, 20) {
		t.Error("failed to perform division by rational")
	}
}

func TestDivideByNum(t *testing.T) {
	r := New(3, 4)

	if r.DivideByNum(7) != New(3, 28) {
		t.Error("failed to perform division by integer")
	}
}

func TestMultiply(t *testing.T) {
	r1 := New(3, 4)
	r2 := New(5, 7)

	if r1.Multiply(r2) != New(15, 28) {
		t.Error("failed to perform multiplication by rational")
	}
}

func TestMultiplyByNum(t *testing.T) {
	r := New(3, 4)

	if r.MultiplyByNum(4) != New(3, 1) {
		t.Error("failed to perform multiplication by integer")
	}
}

func TestAdd(t *testing.T) {
	r1 := New(1, 2)
	r2 := New(5, 4)

	if r1.Add(r2) != New(7, 4) {
		t.Error("failed to perform addition of rational")
	}
}

func TestAddNum(t *testing.T) {
	r := New(2, 4)

	if r.AddNum(2) != New(5, 2) {
		t.Error("failed to perform addition of integer")
	}
}

func TestSubtract(t *testing.T) {
	r1 := New(1, 2)
	r2 := New(5, 4)

	if r1.Subtract(r2) != New(-3, 4) {
		t.Error("failed to perform subtraction of rational")
	}
}

func TestSubtractNum(t *testing.T) {
	r := New(2, 4)

	if r.SubtractNum(2) != New(-3, 2) {
		t.Error("failed to perform subtraction of integer")
	}
}
