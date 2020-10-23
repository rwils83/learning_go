package main

import "fmt"

func main(){
  fmt.Println("This is the opposite of the incremementing, using -- to decrement")
  i := 20
  for i > 0 {
    fmt.Println(i)
    i--
  }
  fmt.Println("Finished")
}
