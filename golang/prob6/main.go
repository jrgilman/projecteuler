package main

import (
  "fmt"
  "math"
)

func main() {

  sumOfSquares := float64(0)
  sum := float64(0)

  for i := float64(1); i <= 100; i++ {

    sum += i
    sumOfSquares += math.Pow(i, 2)

  }

  squareOfSum := math.Pow(sum, 2)

  fmt.Printf("The difference between the square of sums and the sum of squares is: %v\n", uint64(squareOfSum - sumOfSquares))
  return

}
