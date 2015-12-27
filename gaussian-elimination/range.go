package gaussian

import (
	"math"

	"github.com/alex-ant/gomath/rational"
)

// GetAllOptionsInRange retuns a list of all the possible variables' values within
// the passed range.
func GetAllOptionsInRange(res [][]rational.Rational, min, max int64, natural bool) (res2 [][]rational.Rational) {
	unknowns := len(res[0]) - 1

	comb := make([]int64, unknowns)
	for i := range comb {
		comb[i] = min
	}

	for i := 1; i <= int(math.Pow(float64(int(max-min+1)), float64(unknowns))); i++ {
		var tmpRes []rational.Rational
		fitsLimit := true
		for _, jv := range res {
			if jv[0] == rational.New(0, 0) {
				break
			}
			var eqValue rational.Rational
			for k, kv := range jv {
				if k == 0 {
					eqValue = kv
				} else {
					eqValue = eqValue.Subtract(kv.MultiplyByNum(comb[k-1]))
				}
			}
			if (eqValue.GreaterThanNum(max) || eqValue.LessThanNum(min)) || (natural && !eqValue.IsNatural()) {
				fitsLimit = false
			} else {
				tmpRes = append(tmpRes, eqValue)
			}
		}

		if fitsLimit {
			for j := len(comb) - 1; j >= 0; j-- {
				tmpRes = append(tmpRes, rational.New(comb[j], 1))
			}
			res2 = append(res2, tmpRes)
		}

		for k := unknowns - 1; k >= 0; k-- {
			if comb[k] == max {
				comb[k] = min
			} else {
				comb[k]++
				break
			}
		}
	}

	return
}
