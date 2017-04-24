package main

import "fmt"

func main(){

  fibOne := 1
  fibTwo := 2

  sum := 2

  for fibTwo <= 4000000 {
    fibTwo = fibTwo + fibOne
    fibOne = fibTwo - fibOne

    if fibTwo % 2 == 0 {
      sum += fibTwo
    }
  }

  fmt.Printf("The answer is: %v\n", sum)
  return

}
