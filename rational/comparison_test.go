package rational

import "testing"

func TestGreaterThan(t *testing.T) {
	r1 := New(3, 4)
	r2 := New(5, 7)
	r3 := New(-5, 7)
	r4 := New(5, -7)
	r5 := New(-3, -4)
	r6 := New(-3, -4)
	r7 := New(-5, -7)

	if !r1.GreaterThan(r2) || r1.LessThan(r2) {
		t.Error("failed to compare if a positive rational is greater than another positive rational")
	}

	if !r1.GreaterThan(r3) || r1.LessThan(r3) {
		t.Error("failed to compare if a positive rational is greater than a rational with negative numerator")
	}

	if !r1.GreaterThan(r4) || r1.LessThan(r4) {
		t.Error("failed to compare if a positive rational is greater than a rational with negative denominator")
	}

	if !r5.GreaterThan(r2) || r5.LessThan(r2) {
		t.Error("failed to compare if a positive rational is greater than a rational with negative numerator and denominator")
	}

	if !r6.GreaterThan(r7) || r6.LessThan(r7) {
		t.Error("failed to compare if a rational with negative numerator and denominator is greater than a rational with negative denominator and numerator")
	}
}

func TestGreaterThanNum(t *testing.T) {
	r1 := New(5, 4)
	r2 := New(-5, 4)
	r3 := New(5, -4)
	r4 := New(-5, -4)

	if !r1.GreaterThanNum(-1) || r1.LessThanNum(-1) {
		t.Error("failed to compare if a positive rational is greater than a negative integer")
	}

	if !r1.GreaterThanNum(1) || r1.LessThanNum(1) {
		t.Error("failed to compare if a positive rational is greater than a positive integer")
	}

	if !r2.GreaterThanNum(-2) || r2.LessThanNum(-2) {
		t.Error("failed to compare if a rational with negative numerator is greater than a negative integer")
	}

	if !r3.GreaterThanNum(-2) || r3.LessThanNum(-2) {
		t.Error("failed to compare if a rational with negative denominator is greater than a negative integer")
	}

	if !r4.GreaterThanNum(-2) || r4.LessThanNum(-2) {
		t.Error("failed to compare if a rational with negative numerator and denominator is greater than a negative integer")
	}

	if !r4.GreaterThanNum(1) || r4.LessThanNum(1) {
		t.Error("failed to compare if a rational with negative numerator and denominator is greater than a positive integer")
	}

	if !r1.GreaterThanNum(0) || r1.LessThanNum(0) {
		t.Error("failed to compare if a positive rational is greater than zero")
	}
}

func TestLessThan(t *testing.T) {
	r1 := New(5, 7)
	r2 := New(3, 4)
	r3 := New(-5, 7)
	r4 := New(5, -7)
	r5 := New(-3, -4)
	r6 := New(-5, -7)
	r7 := New(-3, -4)

	if !r1.LessThan(r2) || r1.GreaterThan(r2) {
		t.Error("failed to compare if a positive rational is less than another positive rational")
	}

	if !r3.LessThan(r2) || r3.GreaterThan(r2) {
		t.Error("failed to compare if a positive rational is less than a rational with negative numerator")
	}

	if !r4.LessThan(r2) || r4.GreaterThan(r2) {
		t.Error("failed to compare if a positive rational is less than a rational with negative denominator")
	}

	if !r1.LessThan(r5) || r1.GreaterThan(r5) {
		t.Error("failed to compare if a positive rational is less than a rational with negative numerator and denominator")
	}

	if !r6.LessThan(r7) || r6.GreaterThan(r7) {
		t.Error("failed to compare if a rational with negative numerator and denominator is less than another rational with negative numerator and denominator")
	}
}

func TestLessThanNum(t *testing.T) {
	r1 := New(7, 4)
	r2 := New(-5, 4)
	r3 := New(5, -4)
	r4 := New(-7, -4)

	if !r1.LessThanNum(2) || r1.GreaterThanNum(2) {
		t.Error("failed to compare if a positive rational is less than a positive integer")
	}

	if !r2.LessThanNum(-1) || r2.GreaterThanNum(-1) {
		t.Error("failed to compare if a rational with negative numerator is less than a negative integer")
	}

	if !r3.LessThanNum(-1) || r3.GreaterThanNum(-1) {
		t.Error("failed to compare if a rational with negative denominator is less than a negative integer")
	}

	if !r4.LessThanNum(2) || r4.GreaterThanNum(2) {
		t.Error("failed to compare if a rational with negative numerator and denominator is less than a positive integer")
	}

	if !r2.LessThanNum(0) || r2.GreaterThanNum(0) {
		t.Error("failed to compare if a rational with negative numerator is less than zero")
	}
}
