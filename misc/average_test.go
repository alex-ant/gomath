package misc

import (
	"testing"

	"github.com/alex-ant/gomath/rational"
)

func TestFeedAverage(t *testing.T) {
	var avg Average

	checkFloatAndRational := func(numerator int) {
		if avg.Get() != float64(numerator) {
			t.Error("failed to get float64 average")
		}

		if avg.GetR() != rational.New(int64(numerator), 1) {
			t.Error("failed to get Rarional average")
		}
	}

	avg.Feed(float64(10))
	checkFloatAndRational(10)

	avg.Feed(float64(8))
	checkFloatAndRational(9)

	avg.Feed(float64(6))
	checkFloatAndRational(8)

	avg.Feed(float64(4))
	checkFloatAndRational(7)
}
