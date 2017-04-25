package main

import (
  "fmt"
  "math"
)

func main() {

  currentCount := 1
  currentPrime := 2

  for i := currentPrime; currentCount <= 10001; i++  {

    if isPrime(i) {
      currentCount++
      currentPrime = i
    }

  }

  fmt.Printf("The 10001st prime is: %v\n", currentPrime)
  return

}

// speed could be improved by using Miller-Rabin.
func isPrime(toTest int) bool {

  sqrt := math.Sqrt(float64(toTest))

  for i := 2; i <= int(sqrt); i++ {
    if toTest % i == 0 {
      return false
    }
  }

  return true

}
