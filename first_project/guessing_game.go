package main

import (
  "fmt"
  "math/rand"
  "time"
  "bufio"
  "os"

)

func main() {
  fmt.Println("Welcome to the guessing game. This game will generate a random number. You can provide your guess, it will tell you if the number is higher or lower than your guess")
  rand.Seed(time.Now().UnixNano())
  computer_number := rand.Intn(100)
  number_of_guesses := 0
  
  fmt.Println("Enter your name to begin: ")
  input := bufio.NewScanner(os.Stdin)
  input.Scan()
  user_name := input.Text()
  
  if user_name == "test" {
    fmt.Println(computer_number)
  }
  
  var guess int
  for guess != computer_number {
    if number_of_guesses == 0 {
      fmt.Println("All right,", user_name, ". I have chosen a number between 1 and 100. What is your first guess?")
      number_of_guesses ++
    } else {
      fmt.Println("What is your next guess?")
      number_of_guesses ++
    }
    fmt.Scanln(&guess)
    if guess < computer_number{
      fmt.Println("Guess Higher")
    }
    if guess > computer_number {
      fmt.Println("Guess Lower")
    }
  }
  
  if number_of_guesses == 1 {
    fmt.Println("You win! It took you 1 guess")
  } else {
    fmt.Println("You win! It took you", number_of_guesses, "guesses")
  }
  //fmt.Println(computer_number)
  //fmt.Println(user_name)
}
