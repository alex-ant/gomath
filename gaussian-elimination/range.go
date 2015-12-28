package gaussian

import (
	"encoding/json"
	"sync"

	"github.com/alex-ant/gomath/rational"
)

func sliceToString(data []int64) string {
	b, _ := json.Marshal(data)
	return string(b)
}

func stringToSlice(data string) (sl []int64) {
	json.Unmarshal([]byte(data), &sl)
	return
}

// GetAllOptionsInRange retuns a list of all the possible variables' values within
// the passed range.
func GetAllOptionsInRange(res [][]rational.Rational, min, max int64, natural bool) (res2 [][]rational.Rational) {
	unknowns := len(res[0]) - 1

	comb := make([]int64, unknowns)

	combFinal := func() (final bool) {
		final = true
		for _, v := range comb {
			if v != max {
				final = false
				break
			}
		}
		return
	}

	concCh := make(chan bool, 20)

	var resMutex sync.Mutex

	started := true

	for !combFinal() {
		for k := unknowns - 1; k >= 0; k-- {
			if comb[k] == max || started {
				comb[k] = min
			} else {
				comb[k]++
				break
			}
		}
		started = false

		concCh <- true
		go func(combStr string) {
			combSl := stringToSlice(combStr)
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
						eqValue = eqValue.Subtract(kv.MultiplyByNum(combSl[k-1]))
					}
				}
				if (eqValue.GreaterThanNum(max) || eqValue.LessThanNum(min)) || (natural && !eqValue.IsNatural()) {
					fitsLimit = false
				} else {
					tmpRes = append(tmpRes, eqValue)
				}
			}

			if fitsLimit {
				for j := len(combSl) - 1; j >= 0; j-- {
					tmpRes = append(tmpRes, rational.New(combSl[j], 1))
				}
				resMutex.Lock()
				res2 = append(res2, tmpRes)
				resMutex.Unlock()
			}
			<-concCh
		}(sliceToString(comb))
	}

	for i := 0; i < cap(concCh); i++ {
		concCh <- true
	}

	return
}
