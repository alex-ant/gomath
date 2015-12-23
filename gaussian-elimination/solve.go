package gaussian

import (
	"errors"

	"github.com/alex-ant/gomath/rational"
)

// SolveGaussian solves the system of linear equations via The Gaussian Elimination.
func SolveGaussian(eqM [][]rational.Rational) (res []rational.Rational, err error) {
	if len(eqM) > len(eqM[0]) {
		err = errors.New("the number of equations can not be greater than the number of variables")
		return
	}

	for i := 0; i < len(eqM)-1; i++ {
		eqM = sortMatrix(eqM, i)

		var varC rational.Rational
		for k := i; k < len(eqM); k++ {
			if k == i {
				varC = eqM[k][i]
			} else {
				multipliedLine := make([]rational.Rational, len(eqM[i]))
				for z, zv := range eqM[i] {
					multipliedLine[z] = zv.Multiply(eqM[k][i].Divide(varC)).MultiplyByNum(-1)
				}
				newLine := make([]rational.Rational, len(eqM[k]))
				for z, zv := range eqM[k] {
					newLine[z] = zv.Add(multipliedLine[z])
				}
				eqM[k] = newLine
			}
		}
	}

	// Removing empty lines.
	var resultEqM [][]rational.Rational
	for i := len(eqM) - 1; i >= 0; i-- {
		if !rational.RationalsAreNull(eqM[i]) {
			resultEqM = append(resultEqM, eqM[i])
		}
	}

	firstNonZeroIndex := func(sl []rational.Rational) (index int) {
		for i, v := range sl {
			if v.GetNumerator() != 0 {
				index = i
				return
			}
		}
		return
	}

	// Back substitution.
	for z := 0; z < len(resultEqM)-1; z++ {
		var processIndex int
		var firstLine []rational.Rational
		for i := z; i < len(resultEqM); i++ {
			v := resultEqM[i]
			if i == z {
				processIndex = firstNonZeroIndex(v)
				firstLine = v
			} else {
				mult := v[processIndex].Divide(firstLine[processIndex]).MultiplyByNum(-1)
				for j, jv := range v {
					resultEqM[i][j] = firstLine[j].Multiply(mult).Add(jv)
				}
			}
		}
	}

	// Calculating variables.
	res = make([]rational.Rational, len(eqM[0])-1)

	if firstNonZeroIndex(resultEqM[0]) == len(resultEqM[0])-2 {
		for i, v := range resultEqM {
			res[len(res)-1-i] = v[len(v)-1].Divide(v[len(resultEqM)-1-i])
		}
	}

	/*fmt.Println("================ Aaaa")
	for _, v := range resultEqM {
		fmt.Println(v)
	}*/

	return
}

func sortMatrix(m [][]rational.Rational, initRow int) (m2 [][]rational.Rational) {
	indexed := make(map[int]bool)

	for i := 0; i < initRow; i++ {
		m2 = append(m2, m[i])
		indexed[i] = true
	}

	greaterThanMax := func(rr1, rr2 []rational.Rational) (greater bool) {
		for i := 0; i < len(rr1); i++ {
			if rr1[i].GetModule().GreaterThan(rr2[i].GetModule()) {
				greater = true
				return
			} else if rr1[i].GetModule().LessThan(rr2[i].GetModule()) {
				return
			}
		}
		return
	}

	type maxStruct struct {
		index   int
		element []rational.Rational
	}

	for i := initRow; i < len(m); i++ {
		max := maxStruct{-1, make([]rational.Rational, len(m[i]))}
		var firstNotIndexed int
		for k, kv := range m {
			if !indexed[k] {
				firstNotIndexed = k
				if greaterThanMax(kv, max.element) {
					max.index = k
					max.element = kv
				}
			}
		}
		if max.index != -1 {
			m2 = append(m2, max.element)
			indexed[max.index] = true
		} else {
			m2 = append(m2, m[firstNotIndexed])
			indexed[firstNotIndexed] = true
		}
	}

	return
}
