package main

import (
  "fmt"
  "math"
)

func main(){

  number := 600851475143
  largestPrimeFactor := 0

  // slow. O(n^2), second for loop in isPrime function
  for i := 2; i < int(math.Sqrt(float64(number))) ; i++ {
    if isPrime(i) {

      if number % i == 0 && i > largestPrimeFactor {
        largestPrimeFactor = i
        fmt.Println(largestPrimeFactor)
      }

    }
  }

  fmt.Printf("Largest prime is: %v\n", largestPrimeFactor)
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
