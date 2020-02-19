// Programmer name: Jose Padilla
// Date completed:  2/18/2020
// Description: 2.1.1 Lab Part 1

package main

import (
    "fmt"
    "math/rand"
) //adding the ability to do random numbers

func main() {
    var count int

    //ask the user to enter a max range for the guessing game and store that value in variable max.
    fmt.Println("Enter a max range")
    var max int
    fmt.Scanln(&max)
    //this next line creates a random number from 1 to that guess for the computer to know.  You can test this by printing out the variable computerGuess
    var computerGuess = rand.Intn(max)
   
    //ask the user to enter a guess for the computer number
  fmt.Println("Please guess the computers number")
  
  var userGuess int
  fmt.Scanln(&userGuess)
  for userGuess != computerGuess {
    count++
    
  
    fmt.Println("Wrong")
    fmt.Println("Please try again")
    fmt.Scanln(&userGuess)
  }
  fmt.Println("You got the answer in ", count, " try(ies).")
 
}