package main

import "fmt"

func main(){
  fmt.Println("This will use the logic while i < 20, print i")
  i := 0
  for i < 20{
    fmt.Println("i = ", i)
    i += 1
  }
    fmt.Println("Finished")
  }
