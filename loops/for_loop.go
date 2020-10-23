package main

import "fmt"

func main() {
  sum := 0
  fmt.Println("The sum at the beginning is: ",sum)
  for i :=0;i < 20; i++ {
    sum += i
    previous_sum := sum - i
    fmt.Println(previous_sum, " + ", i, " = ",sum)
  }
  fmt.Println(sum)
}

