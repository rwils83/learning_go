package main

import (
  "fmt"
  "math/rand" // This is not cryptographically secure!
  "time" 
)

func main() {
  // Seeding rand
  fmt.Println("This will use switch/case to print whether a random number is larger or smaller than 50")
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Choosing a random number:")
  num := rand.Intn(100)
  fmt.Println("The random number is: ", num)
  switch {
  case num < 50:
    fmt.Println("The random number is less than 50")
  case num == 50:
    fmt.Println("The random number equals 50")
  default:
    fmt.Println("More than 50") 
}
}

