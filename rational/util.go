package rational

import "math/big"

// Simplify simplifies the rational number by dividing it's numerator and
// denominator by the GCD.
func (ev *Rational) Simplify() {
	if ev.denominator < 0 {
		ev.numerator *= -1
		ev.denominator *= -1
	}

	currentNumerator := ev.numerator
	currentDenominator := ev.denominator

	if currentNumerator < 0 && currentDenominator > 0 {
		currentNumerator *= -1
	} else if currentDenominator < 0 && currentNumerator >= 0 {
		currentDenominator *= -1
	}

	n := big.NewInt(currentNumerator)
	d := big.NewInt(currentDenominator)

	gcd := new(big.Int).GCD(nil, nil, n, d).Int64()

	if gcd > 1 {
		ev.numerator /= gcd
		ev.denominator /= gcd
	}
}

// IsNatural determines whether the rational number is also natural.
func (ev Rational) IsNatural() bool {
	if ev.numerator%ev.denominator == 0 {
		return true
	}
	return false
}

// Float64 returns the float64 representation of a rational number.
func (ev Rational) Float64() float64 {
	return float64(ev.numerator) / float64(ev.denominator)
}

// Get returns a value.
func (ev Rational) Get() (numerator, denominator int64) {
	return ev.numerator, ev.denominator
}

// GetNumerator returns a numerator.
func (ev Rational) GetNumerator() int64 {
	return ev.numerator
}

// GetDenominator returns a denominator.
func (ev Rational) GetDenominator() int64 {
	return ev.denominator
}

// GetModule returns rational number's module.
func (ev Rational) GetModule() Rational {
	solveNegatives(&ev.numerator, &ev.denominator)
	if ev.LessThanNum(0) {
		ev = ev.MultiplyByNum(-1)
	}
	return ev
}

// IsNull determines whether the value is zero.
func (ev Rational) IsNull() (n bool) {
	if ev.numerator == 0 {
		n = true
	}
	return
}

// RationalsAreNull determines whether the slice of Rationals contains only zero values.
func RationalsAreNull(l []Rational) (isNull bool) {
	isNull = true
	for _, v := range l {
		if !v.IsNull() {
			isNull = false
		}
	}
	return
}

func solveNegatives(n, d *int64) {
	if *d < 0 {
		*n *= -1
		*d *= -1
	}
}
