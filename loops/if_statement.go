package main

import "fmt"

func main(){
  fmt.Println("This will take in two numbers from the user, assign the numbers to a(First number) or b(Second Number) and return if a or b is larger")
  fmt.Println("Enter your first number: ")
  var a int
  fmt.Scanln(&a)
  var b int
  fmt.Println("Enter your second number: ")
  fmt.Scanln(&b)
  if a > b {
    fmt.Println("a is > b")
  } else{
    fmt.Println("b is > a")
  }
}
