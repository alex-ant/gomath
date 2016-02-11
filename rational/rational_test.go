package rational

import "testing"

func TestNewFromFloat(t *testing.T) {
	r1, err1 := NewFromFloat(1.25)
	if err1 != nil {
		t.Error(err1)
	}

	r2, err2 := NewFromFloat(-1.25)
	if err2 != nil {
		t.Error(err2)
	}

	r3, err3 := NewFromFloat(float64(7))
	if err3 != nil {
		t.Error(err3)
	}

	if r1 != New(5, 4) {
		t.Error("failed to convert a positive float to rational")
	}

	if r2 != New(-5, 4) {
		t.Error("failed to convert a negative float to rational")
	}

	if r3 != New(7, 1) {
		t.Error("failed to convert a positive float to rational 2")
	}
}

func TestDivide(t *testing.T) {
	r1 := New(3, 4)
	r2 := New(5, 7)
	r3 := New(-5, 7)
	r4 := New(5, -7)
	r5 := New(-5, -7)
	r6 := New(-3, 4)

	if r1.Divide(r2) != New(21, 20) {
		t.Error("failed to perform division of a positive rational by another positive rational")
	}

	if r1.Divide(r3) != New(-21, 20) {
		t.Error("failed to perform division of a positive rational by a rational with negative numerator")
	}

	if r1.Divide(r4) != New(-21, 20) {
		t.Error("failed to perform division of a positive rational by a rational with negative denominator")
	}

	if r1.Divide(r5) != New(21, 20) {
		t.Error("failed to perform division of a positive rational by a rational with negative numerator and denominator")
	}

	if r3.Divide(r1) != New(-20, 21) {
		t.Error("failed to perform division of a rational with negative numerator by a positive rational")
	}

	if r4.Divide(r1) != New(-20, 21) {
		t.Error("failed to perform division of a rational with negative denominator by a positive rational")
	}

	if r5.Divide(r1) != New(20, 21) {
		t.Error("failed to perform division of a rational with negative numerator and denominator by a positive rational")
	}

	if r6.Divide(r3) != New(21, 20) {
		t.Error("failed to perform division of a rational with negative numerator by a rational with negative numerator")
	}

	if r6.Divide(r4) != New(21, 20) {
		t.Error("failed to perform division of a rational with negative numerator by a rational with negative denominator")
	}
}

func TestDivideByNum(t *testing.T) {
	r1 := New(3, 4)
	r2 := New(-3, 4)
	r3 := New(3, -4)
	r4 := New(-3, -4)

	if r1.DivideByNum(7) != New(3, 28) {
		t.Error("failed to perform division of a positive rational by a positive integer")
	}

	if r2.DivideByNum(7) != New(-3, 28) {
		t.Error("failed to perform division of a rational with negative numerator by a positive integer")
	}

	if r3.DivideByNum(7) != New(-3, 28) {
		t.Error("failed to perform division of a rational with negative denominator by a positive integer")
	}

	if r4.DivideByNum(7) != New(3, 28) {
		t.Error("failed to perform division of a rational with negative numerator and denominator by a positive integer")
	}

	if r1.DivideByNum(-7) != New(-3, 28) {
		t.Error("failed to perform division of a positive rational by a negative integer")
	}

	if r2.DivideByNum(-7) != New(3, 28) {
		t.Error("failed to perform division of a rational with negative numerator by a negative integer")
	}

	if r3.DivideByNum(-7) != New(3, 28) {
		t.Error("failed to perform division of a rational with negative denominator by a negative integer")
	}

	if r4.DivideByNum(-7) != New(-3, 28) {
		t.Error("failed to perform division of a rational with negative numerator and denominator by a negative integer")
	}
}

func TestMultiply(t *testing.T) {
	r1 := New(3, 4)
	r2 := New(5, 7)
	r3 := New(-5, 7)
	r4 := New(5, -7)
	r5 := New(-5, -7)
	r6 := New(-3, 4)

	if r1.Multiply(r2) != New(15, 28) {
		t.Error("failed to perform multiplication of a positive rational by another positive rational")
	}

	if r1.Multiply(r3) != New(-15, 28) {
		t.Error("failed to perform multiplication of a positive rational by a rational with negative numerator")
	}

	if r1.Multiply(r4) != New(-15, 28) {
		t.Error("failed to perform multiplication of a positive rational by a rational with negative denominator")
	}

	if r1.Multiply(r5) != New(15, 28) {
		t.Error("failed to perform multiplication of a positive rational by a rational with negative numerator and denominator")
	}

	if r3.Multiply(r1) != New(-15, 28) {
		t.Error("failed to perform multiplication of a rational with negative numerator by a positive rational")
	}

	if r4.Multiply(r1) != New(-15, 28) {
		t.Error("failed to perform multiplication of a rational with negative denominator by a positive rational")
	}

	if r5.Multiply(r1) != New(15, 28) {
		t.Error("failed to perform multiplication of a rational with negative numerator and denominator by a positive rational")
	}

	if r6.Multiply(r3) != New(15, 28) {
		t.Error("failed to perform multiplication of a rational with negative numerator by a rational with negative numerator")
	}

	if r6.Multiply(r4) != New(15, 28) {
		t.Error("failed to perform multiplication of a rational with negative numerator by a rational with negative denominator")
	}
}

func TestMultiplyByNum(t *testing.T) {
	r1 := New(3, 4)
	r2 := New(-3, 4)
	r3 := New(3, -4)
	r4 := New(-3, -4)

	if r1.MultiplyByNum(4) != New(3, 1) {
		t.Error("failed to perform multiplication of a positive rational by a positive integer")
	}

	if r2.MultiplyByNum(4) != New(-3, 1) {
		t.Error("failed to perform multiplication of a rational with negative numerator by a positive integer")
	}

	if r3.MultiplyByNum(4) != New(-3, 1) {
		t.Error("failed to perform multiplication of a rational with negative denominator by a positive integer")
	}

	if r4.MultiplyByNum(4) != New(3, 1) {
		t.Error("failed to perform multiplication of a rational with negative numerator and denominator by a positive integer")
	}

	if r1.MultiplyByNum(-4) != New(-3, 1) {
		t.Error("failed to perform multiplication of a positive rational by a negative integer")
	}

	if r2.MultiplyByNum(-4) != New(3, 1) {
		t.Error("failed to perform multiplication of a rational with negative numerator by a negative integer")
	}

	if r3.MultiplyByNum(-4) != New(3, 1) {
		t.Error("failed to perform multiplication of a rational with negative denominator by a negative integer")
	}

	if r4.MultiplyByNum(-4) != New(-3, 1) {
		t.Error("failed to perform multiplication of a rational with negative numerator and denominator by a negative integer")
	}
}

func TestAdd(t *testing.T) {
	r1 := New(1, 2)
	r2 := New(5, 4)
	r3 := New(-5, 4)
	r4 := New(5, -4)
	r5 := New(-5, -4)

	if r1.Add(r2) != New(7, 4) {
		t.Error("failed to add a positive rational to another positive rational")
	}

	if r1.Add(r3) != New(-3, 4) {
		t.Error("failed to add a rational with negative numerator to positive rational")
	}

	if r1.Add(r4) != New(-3, 4) {
		t.Error("failed to add a rational with negative denominator to positive rational")
	}

	if r1.Add(r5) != New(7, 4) {
		t.Error("failed to add a rational with negative numerator and denominator to positive rational")
	}

	if r3.Add(r1) != New(-3, 4) {
		t.Error("failed to add a positive rational to rational with negative numerator")
	}

	if r4.Add(r1) != New(-3, 4) {
		t.Error("failed to add a positive rational to rational with negative denominator")
	}

	if r5.Add(r1) != New(7, 4) {
		t.Error("failed to add a positive rational to rational with negative numerator and denominator")
	}
}

func TestAddNum(t *testing.T) {
	r1 := New(2, 4)
	r2 := New(-2, 4)
	r3 := New(2, -4)
	r4 := New(-2, -4)

	if r1.AddNum(2) != New(5, 2) {
		t.Error("failed to add a positive rational to positive integer")
	}

	if r2.AddNum(2) != New(3, 2) {
		t.Error("failed to add a rational with negative numerator to positive integer")
	}

	if r3.AddNum(2) != New(3, 2) {
		t.Error("failed to add a rational with negative denominator to positive integer")
	}

	if r4.AddNum(2) != New(5, 2) {
		t.Error("failed to add a rational with negative numerator and denominator to positive integer")
	}

	if r1.AddNum(-2) != New(-3, 2) {
		t.Error("failed to add a positive rational to negative integer")
	}

	if r2.AddNum(-2) != New(-5, 2) {
		t.Error("failed to add a rational with negative numerator to negative integer")
	}

	if r3.AddNum(-2) != New(-5, 2) {
		t.Error("failed to add a rational with negative denominator to negative integer")
	}

	if r4.AddNum(-2) != New(-3, 2) {
		t.Error("failed to add a rational with negative numerator and denominator to negative integer")
	}

	r0 := New(0, 1)
	if r0.AddNum(156) != New(156, 1) {
		t.Error("failed to add a zero rational to positive integer")
	}
}

func TestSubtract(t *testing.T) {
	r1 := New(1, 2)
	r2 := New(5, 4)
	r3 := New(-5, 4)
	r4 := New(5, -4)
	r5 := New(-5, -4)

	if r1.Subtract(r2) != New(-3, 4) {
		t.Error("failed to subtract a positive rational from another positive rational")
	}

	if r1.Subtract(r3) != New(7, 4) {
		t.Error("failed to subtract a rational with negative numerator from positive rational")
	}

	if r1.Subtract(r4) != New(7, 4) {
		t.Error("failed to subtract a rational with negative denominator from positive rational")
	}

	if r1.Subtract(r5) != New(-3, 4) {
		t.Error("failed to subtract a rational with negative numerator and denominator from positive rational")
	}

	if r3.Subtract(r1) != New(-7, 4) {
		t.Error("failed to subtract a positive rational from rational with negative numerator")
	}

	if r4.Subtract(r1) != New(-7, 4) {
		t.Error("failed to subtract a positive rational from rational with negative denominator")
	}

	if r5.Subtract(r1) != New(3, 4) {
		t.Error("failed to subtract a positive rational from rational with negative numerator and denominator")
	}
}

func TestSubtractNum(t *testing.T) {
	r1 := New(2, 4)
	r2 := New(-2, 4)
	r3 := New(2, -4)
	r4 := New(-2, -4)

	if r1.SubtractNum(2) != New(-3, 2) {
		t.Error("failed to subtract a positive integer from positive rational")
	}

	if r2.SubtractNum(2) != New(-5, 2) {
		t.Error("failed to subtract a positive integer from rational with negative numerator")
	}

	if r3.SubtractNum(2) != New(-5, 2) {
		t.Error("failed to subtract a positive integer from rational with negative denominator")
	}

	if r4.SubtractNum(2) != New(-3, 2) {
		t.Error("failed to subtract a positive integer from rational with negative numerator and denominator")
	}

	if r1.SubtractNum(-2) != New(5, 2) {
		t.Error("failed to subtract a negative integer from positive rational")
	}

	if r2.SubtractNum(-2) != New(3, 2) {
		t.Error("failed to subtract a negative integer from rational with negative numerator")
	}

	if r3.SubtractNum(-2) != New(3, 2) {
		t.Error("failed to subtract a negative integer from rational with negative denominator")
	}

	if r4.SubtractNum(-2) != New(5, 2) {
		t.Error("failed to subtract a negative integer from rational with negative numerator and denominator")
	}
}
