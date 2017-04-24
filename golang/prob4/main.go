package main

import (
  "fmt"
  "strings"
  "strconv"
)

func main() {

  largest := 0

  for i := 100; i < 1000; i++ {
    for j := 100; j < 1000; j++ {

      if i*j > largest && isPalindrome(i*j) {
        largest = i*j
      }

    }
  }

  fmt.Printf("The largest palindrome created by multiplying two 3-digit numbers is: %v\n", largest)
  return

}

func isPalindrome(toCheck int) bool {

  normal := strconv.Itoa(toCheck)
  reverse := ""

  for i := len(normal) - 1; i >= 0; i-- {
    reverse += string(normal[i])
  }

  if strings.Compare(normal, reverse) == 0 {
    return true
  }

  return false

}
