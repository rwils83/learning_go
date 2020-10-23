package main

import "fmt"


func main() {
  defer fmt.Println("This is  the first line of main, but will run at the end")

  fmt.Println("Main running")
  fmt.Println("Main ended")
}
