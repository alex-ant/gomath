package rational

import "testing"

func TestSimplify(t *testing.T) {
	r1 := New(6, 8)
	r1.Simplify()

	r2 := New(-6, 8)
	r2.Simplify()

	r3 := New(6, -8)
	r3.Simplify()

	r4 := New(-6, -8)
	r4.Simplify()

	if r1 != New(3, 4) {
		t.Error("failed to simplify positive rational")
	}

	if r2 != New(-3, 4) {
		t.Error("failed to simplify rational with negative numerator")
	}

	if r3 != New(-3, 4) {
		t.Error("failed to simplify rational with negative denominator")
	}

	if r4 != New(3, 4) {
		t.Error("failed to simplify rational with negative numerator and denominator")
	}
}

func TestSimplifyLine(t *testing.T) {
	r := []Rational{
		New(1, 4),
		New(2, 8),
		New(3, 16),
		New(4, 4),
		New(-4, 4),
		New(4, -4),
		New(-4, -4),
	}

	r2 := SimplifyLine(r)

	r2Test := []Rational{
		New(1, 1),
		New(1, 1),
		New(3, 4),
		New(4, 1),
		New(-4, 1),
		New(-4, 1),
		New(4, 1),
	}

	success := true
	for i, v := range r2 {
		if v != r2Test[i] {
			success = false
		}
	}

	if !success {
		t.Error("failed to simplify line of rationals")
	}
}

func TestFloat64(t *testing.T) {
	r1 := New(7, 8).Float64()
	r2 := New(-7, 8).Float64()
	r3 := New(7, -8).Float64()
	r4 := New(-7, -8).Float64()

	if r1 != 0.875 {
		t.Error("failed to convert positive rational to float")
	}

	if r2 != -0.875 {
		t.Error("failed to convert rational with negative numerator to float")
	}

	if r3 != -0.875 {
		t.Error("failed to convert rational with negative denominator to float")
	}

	if r4 != 0.875 {
		t.Error("failed to convert rational with negative numerator and denominator to float")
	}
}

func TestGet(t *testing.T) {
	r1 := New(7, 8)

	n, d := r1.Get()

	if n != 7 || d != 8 {
		t.Error("failed to get rational parts")
	}
}

func TestGetNumerator(t *testing.T) {
	r1 := New(7, 8)

	n := r1.GetNumerator()

	if n != 7 {
		t.Error("failed to get rational numerator")
	}
}

func TestGetDenominator(t *testing.T) {
	r1 := New(7, 8)

	d := r1.GetDenominator()

	if d != 8 {
		t.Error("failed to get rational denominator")
	}
}

func TestGetModule(t *testing.T) {
	r1 := New(7, 8)
	r2 := New(-7, 8)
	r3 := New(7, -8)
	r4 := New(-7, -8)

	m1 := r1.GetModule()
	m2 := r2.GetModule()
	m3 := r3.GetModule()
	m4 := r4.GetModule()

	m0 := New(7, 8)

	if m1 != m0 {
		t.Error("failed to get module of a positive rational")
	}

	if m2 != m0 {
		t.Error("failed to get module of a rational with negative numerator")
	}

	if m3 != m0 {
		t.Error("failed to get module of a rational with negative denominator")
	}

	if m4 != m0 {
		t.Error("failed to get module of a rational with negative numerator and denominator")
	}
}

func TestIsNull(t *testing.T) {
	r1 := New(0, 1)
	r2 := New(1, 1)

	if !r1.IsNull() || r2.IsNull() {
		t.Error("failed to determine whether a rational is null")
	}
}

func TestRationalsAreNull(t *testing.T) {
	r1 := []Rational{
		New(0, 1),
		New(0, 1),
		New(0, 1),
		New(0, 1),
		New(0, 1),
	}

	r2 := []Rational{
		New(1, 4),
		New(2, 8),
		New(3, 16),
		New(4, 4),
		New(-4, 4),
		New(4, -4),
		New(-4, -4),
	}

	if !RationalsAreNull(r1) || RationalsAreNull(r2) {
		t.Error("failed to determine whether a rationals' slice is null")
	}
}
