// Programmer name: Jose Padilla
// Date completed:  2/27/2020
// Description: Paystub generator and updated w/ win conditionsRock Paper Scissors game

package main

import (
    "fmt"
   "math/rand"
   "time"
)


func main() {
  //variable for wage
  fmt.Println("Please enter your hourly wage.")
  var wage float64
  fmt.Scan(&wage)
//variable for overtime wage
  fmt.Print("Please enter your overtime wage.")
  var overtime float64
  fmt.Scan(&overtime)
//variable for hours accrued
  fmt.Print("Please enter the amount of hours you have worked.")
  var hours float64
  fmt.Scan(&hours)
//subtracts 40 from hours if its greater than 40 to get overtime hours
  var hoursOt float64
    if hours > 40{
    hoursOt = hours - 40  
  }   else {
    hoursOt = 0
  }
//payStub
payStub := (((hours - hoursOt) * wage) + (hoursOt * overtime)) * 0.98
fmt.Print("Your total pay after a 12% tax rate will be ",payStub)



//begin updated rock paper scissors
  var user int 
  var cpu int
  //Seeding the random
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  cpu = r1.Intn(3)

  fmt.Println("Please enter 8 for rock, 13 for scissors or 18 for paper")
  fmt.Scanln(&user)
 //cpu choice
    switch cpu {
    case 0: fmt.Println("Computer chose rock")
    case 1: fmt.Println("Computer chose scissors")
    case 2: fmt.Println("Computer chose papers")
  }//user choice
    switch user {
    case 8: fmt.Println("You chose rock")
    case 13: fmt.Println("You chose scissors")
    case 18: fmt.Println("You chose papers")
  }
 //win conditions
  win := cpu + user
    switch win {
      case 8: fmt.Println("You tie")//rock,rock
      case 9: fmt.Println("You lose")//scissors,rock
      case 10: fmt.Println("You lose")//paper,rock

      case 13: fmt.Println("You lose")//rock,scissors
      case 14: fmt.Println("You tie")//scissors,scissors
      case 15: fmt.Println("You win")//paper,scissors

      case 18: fmt.Println("You win")//rock,paper
      case 19: fmt.Println("You lost")//scissors,paper
      case 20: fmt.Println("You tie")//paper,paper

    }
    fmt.Println(win)
}