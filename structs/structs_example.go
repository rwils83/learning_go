package main

import "fmt"

type Student struct {
  FirstName string
  LastName string
  GradeLevel int
}

func main() {
  //Make an instance
  studentOne := Student{"Michael", "Myers", 16}
  fmt.Println(studentOne.FirstName)
  fmt.Println(studentOne)
  //Initialize second student with only a first name
  studentTwo := Student{FirstName:"Jim"}
  fmt.Println(studentTwo)
  //pointer to studentOne
  myPointer := &studentOne
  fmt.Println(myPointer.FirstName)
  //and create a pointer randomly
  myPointer2 := &Student{"Freddie", "Kruger", 2}
  fmt.Println(myPointer2.LastName)
}
