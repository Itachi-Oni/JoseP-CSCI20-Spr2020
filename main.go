// Programmer name: Jose Padilla
// Date completed:  2/27/2020
// Description: Rock Paper Scissors game-p1 vs cpu

package main

import (
    "fmt"
    "math/rand"
    "time"
) //adding the ability to do random numbers

func main() {

  var user int 
  var cpu int
    //create two variables - one for the computer and one for the user

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  cpu = r1.Intn(3)

    //use a random integer value representing the computer's choice in a game of Rock, Scissors, Paper. 0=rock, 1=scissors, 2=paper

  fmt.Println("Please enter 0 for rock, 1 for scissors or 2 for paper")
  fmt.Scanln(&user)
    //prompt the user for an integer value representing the player's choice
  if cpu == 0 {
    fmt.Println("Computer chose rock")
  } else if cpu == 1 {
    fmt.Println("Computer chose scissors")
  }else if cpu == 2 {
    fmt.Println("Computer chose papers")
  }
  if user == 0 {
    fmt.Println("you chose rock")
  } else if user == 1 {
    fmt.Println("you chose scissors")
  }else if user == 2 {
    fmt.Println("you chose papers")
  }
  
    //Print out the values using the words rock, scissors, paper.  ie. "Computer chose rock and player chose paper"
    //You will need to use decisions for this
}