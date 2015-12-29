# gomath
Extra mathematical algorithms and data types for Golang.

## Packages
The repository provides the following packages:
* [rational](https://github.com/alex-ant/gomath/tree/master/rational)
* [gaussian](https://github.com/alex-ant/gomath/tree/master/gaussian-elimination)
* [misc](https://github.com/alex-ant/gomath/tree/master/misc)

Follow the links above to find more about each package.

## Quick overview
Even though each package contains its own README, let's consider a few examples for each one here.

#### rational
The package provides the Rational data type and a kit of corresponding methods to work with rational numbers.

Imagine you have the following expression to solve:

![example 1](https://raw.githubusercontent.com/alex-ant/gomath/master/example1.gif)

Here's how to solve it using the rational package:
```go
package main

import (
	"fmt"
	"log"

	"github.com/alex-ant/gomath/rational"
)

func main() {
	// Create a new rational number defining it's numerator and denominator.
	r1 := rational.New(1, 2)
	// This also can be done using an existing float number.
	r1, err := rational.NewFromFloat(0.5)
	if err != nil {
		log.Fatal(err)
	}

	r2 := rational.New(3, 7)
	r3 := rational.New(4, 5)

  // Add 3/7 to 1/2, multiply by 2 and divide by 4/5.
	result := r1.Add(r2).MultiplyByNum(2).Divide(r3)
	fmt.Println(result)           // {65 28}
	fmt.Println(result.Float64()) // 2.3214285714285716
}
```
