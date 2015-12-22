package rational

import (
	"math"
	"strconv"
	"strings"
)

// Rational stores a rational value.
type Rational struct {
	numerator   int64
	denominator int64
}

// New returns new rational number representation.
func New(n, d int64) (ev Rational) {
	return Rational{
		numerator:   n,
		denominator: d,
	}
}

// NewFromFloat returns new rational number representation retrieved from float64.
func NewFromFloat(f float64) (ev Rational, err error) {
	d, _ := math.Modf(f)

	fStr := strings.Split(strconv.FormatFloat(f, 'f', -1, 64), ".")[1]

	var numerator int64
	numerator, err = strconv.ParseInt(fStr, 10, 64)
	if err != nil {
		return
	}
	denominator := int64(math.Pow(10, float64(len(fStr))))

	ev = New(denominator*int64(d)+numerator, denominator)
	ev.Simplify()

	return
}

// Divide divides a rational value by the provided one.
func (ev Rational) Divide(e Rational) (nv Rational) {
	newNumerator := ev.numerator * e.denominator
	newDenominator := ev.denominator * e.numerator
	if newNumerator == newDenominator {
		nv = New(1, 1)
	} else {
		solveNegatives(&newNumerator, &newDenominator)
		nv = New(newNumerator, newDenominator)
		nv.Simplify()
	}

	return
}

// DivideByNum divides a rational value by the provided integer.
func (ev Rational) DivideByNum(i int64) Rational {
	return ev.Divide(New(i, 1))
}

// Multiply multiplies a rational value by provided one.
func (ev Rational) Multiply(e Rational) (nv Rational) {
	newNumerator := ev.numerator * e.numerator
	if newNumerator != 0 {
		newDenominator := ev.denominator * e.denominator
		if newNumerator == newDenominator {
			nv = New(1, 1)
		} else {
			solveNegatives(&newNumerator, &newDenominator)
			nv = New(newNumerator, newDenominator)
			nv.Simplify()
		}
	}
	return
}

// MultiplyByNum multiplies a rational value by the provided integer.
func (ev Rational) MultiplyByNum(i int64) Rational {
	return ev.Multiply(New(i, 1))
}

// Add adds the provided rational value to an existing one.
func (ev Rational) Add(e Rational) (nv Rational) {
	newNumerator := ev.numerator*e.denominator + e.numerator*ev.denominator
	if newNumerator != 0 {
		newDenominator := ev.denominator * e.denominator
		if newNumerator == newDenominator {
			nv = New(1, 1)
		} else {
			solveNegatives(&newNumerator, &newDenominator)
			nv = New(newNumerator, newDenominator)
			nv.Simplify()
		}
	}
	return
}

// AddNum adds the provided integer to an existing rational number.
func (ev Rational) AddNum(i int64) Rational {
	return ev.Add(New(i, 1))
}

// Subtract subtracts the provided rational value from an existing one.
func (ev Rational) Subtract(e Rational) (nv Rational) {
	newNumerator := ev.numerator*e.denominator - e.numerator*ev.denominator
	if newNumerator != 0 {
		newDenominator := ev.denominator * e.denominator
		if newNumerator == newDenominator {
			nv = New(1, 1)
		} else {
			solveNegatives(&newNumerator, &newDenominator)
			nv = New(newNumerator, newDenominator)
			nv.Simplify()
		}
	}
	return
}

// SubtractNum subtracts the provided integer from an existing rational number.
func (ev Rational) SubtractNum(i int64) Rational {
	return ev.Subtract(New(i, 1))
}

// GetRationalMatrix converts int64 matrix into Rational matrix.
func GetRationalMatrix(m [][]int64) (m2 [][]Rational) {
	m2 = make([][]Rational, len(m))
	for i, iv := range m {
		mr := make([]Rational, len(m[i]))
		for j, jv := range iv {
			mr[j] = New(jv, 1)
		}
		m2[i] = mr
	}
	return
}
