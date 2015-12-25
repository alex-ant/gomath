package gaussian

import (
	"testing"

	"github.com/alex-ant/gomath/rational"
)

func TestSolveGaussian(t *testing.T) {
	m := make([][]int64, 3)
	m[0] = []int64{1, 2, 1, 20}
	m[1] = []int64{3, 1, 2, 29}
	m[2] = []int64{1, 1, 1, 15}

	m2 := make([][]rational.Rational, len(m))
	for i, iv := range m {
		mr := make([]rational.Rational, len(m[i]))
		for j, jv := range iv {
			mr[j] = rational.New(jv, 1)
		}
		m2[i] = mr
	}

	resTest := make([][]rational.Rational, 3)
	resTest[0] = append(resTest[0], rational.New(4, 1))
	resTest[1] = append(resTest[1], rational.New(5, 1))
	resTest[2] = append(resTest[2], rational.New(6, 1))

	res, gausErr := SolveGaussian(m2, false)
	if gausErr != nil {
		t.Error(gausErr)
	}

	success := true
	for i, v := range res {
		if v[0] != resTest[i][0] {
			success = false
		}
	}

	if !success {
		t.Error("failed to solve the system of linear equations")
	}
}
