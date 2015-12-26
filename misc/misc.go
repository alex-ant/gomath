package misc

import "math/big"

// MultiGCD returns the Greatest Common Delimiter for all the numbers in the slice.
func MultiGCD(line []int64) (gcd int64) {
	var uniq []int64
	for _, v := range line {
		var exists bool
		for _, d := range uniq {
			if v == d {
				exists = true
				break
			}
		}
		if !exists {
			uniq = append(uniq, v)
		}
	}

	gcd = uniq[0]
	for i := 1; i < len(uniq); i++ {
		n := big.NewInt(gcd)
		d := big.NewInt(uniq[i])
		gcd = new(big.Int).GCD(nil, nil, n, d).Int64()
	}

	return
}
