package rational

// GreaterThan returns true if a rational is greater than the passed one.
func (ev Rational) GreaterThan(e Rational) bool {
	solveNegatives(&ev.numerator, &ev.denominator)
	solveNegatives(&e.numerator, &e.denominator)
	if e.GetDenominator() == 0 {
		return ev.Float64() > 0
	}
	return ev.numerator*e.denominator > e.numerator*ev.denominator
}

// GreaterThanNum returns true if a rational is greater than the passed integer.
func (ev Rational) GreaterThanNum(i int64) bool {
	return ev.GreaterThan(New(i, 1))
}

// LessThan returns true if a rational is less than the passed one.
func (ev Rational) LessThan(e Rational) bool {
	solveNegatives(&ev.numerator, &ev.denominator)
	solveNegatives(&e.numerator, &e.denominator)
	if e.GetDenominator() == 0 {
		return ev.Float64() < 0
	}
	return ev.numerator*e.denominator < e.numerator*ev.denominator
}

// LessThanNum returns true if a rational is less than the passed integer.
func (ev Rational) LessThanNum(i int64) bool {
	return ev.LessThan(New(i, 1))
}
