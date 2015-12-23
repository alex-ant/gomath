package rational

import (
	"math/big"

	"github.com/alex-ant/gomath/misc"
)

// Simplify simplifies the rational number by dividing it's numerator and
// denominator by the GCD.
func (ev *Rational) Simplify() {
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

// SimplifyLine miltiplies all the rationals in the line by the
// Greatest Common Delimiter if one is greater than 1.
func SimplifyLine(line []Rational) []Rational {
	var denominators []int64
	for _, v := range line {
		denominators = append(denominators, v.GetDenominator())
	}

	gcd := misc.MultiGCD(denominators)

	if gcd > 1 {
		for i, v := range line {
			line[i] = v.MultiplyByNum(gcd)
		}
	}

	return line
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
	if *n < 0 && *d < 0 {
		*n *= -1
		*d *= -1
	}
}
