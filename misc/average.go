package misc

import "github.com/alex-ant/gomath/rational"

// Average stores dynamic average data.
type Average struct {
	num  int
	avg  float64
	avgR rational.Rational
}

// Feed recalculates the average basing on the new value.
func (a *Average) Feed(value interface{}) {
	floatVal, fSuccess := value.(float64)
	rationalVal, rSuccess := value.(rational.Rational)

	intVal, intSuccess := value.(int)
	if intSuccess {
		floatVal = float64(intVal)
		fSuccess = true
	}

	int64Val, int64Success := value.(int64)
	if int64Success {
		floatVal = float64(int64Val)
		fSuccess = true
	}

	var err error

	a.num++

	switch {
	case fSuccess:
		if a.num == 1 {
			a.avgR, err = rational.NewFromFloat(floatVal)
			if err != nil {
				a.num--
				return
			}
			a.avg = floatVal
		} else {
			var rVal rational.Rational
			rVal, err = rational.NewFromFloat(floatVal)
			if err != nil {
				a.num--
				return
			}
			a.avg = (a.avg*float64(a.num-1) + floatVal) / float64(a.num)
			a.avgR = a.avgR.MultiplyByNum(int64(a.num - 1)).Add(rVal).DivideByNum(int64(a.num))
		}
	case rSuccess:
		if a.num == 1 {
			a.avg = rationalVal.Float64()
			a.avgR = rationalVal
		} else {
			a.avg = (a.avg*float64(a.num-1) + rationalVal.Float64()) / float64(a.num)
			a.avgR = a.avgR.MultiplyByNum(int64(a.num - 1)).Add(rationalVal).DivideByNum(int64(a.num))
		}
	default:
		a.num--
		return
	}
}

// Get returns an average in the float64 format.
func (a *Average) Get() float64 {
	return a.avg
}

// GetR returns an average in the Rational format.
func (a *Average) GetR() rational.Rational {
	return a.avgR
}
