//Created by: Jose Padilla
//Created on: Some day
//Description: This program runs the code for a lab game called the 99 trick.
  
package main

import "fmt"

func main() {

  var noun0 string
  fmt.Print("Enter a noun: ")
  fmt.Scan(&noun0)
  fmt.Println(noun0)
  
   var adjective0 string
  fmt.Print("Enter an adejective: ")
  fmt.Scan(&adjective0)
  fmt.Println(adjective0)

   var noun1 string
  fmt.Print("Enter a noun: ")
  fmt.Scan(&noun1)
  fmt.Println(noun1)

   var location0 string
  fmt.Print("Enter a location: ")
  fmt.Scan(&location0)
  fmt.Println(location0)

   var adjective1 string
  fmt.Print("Enter an adjective: ")
  fmt.Scan(&adjective1)
  fmt.Println(adjective1)

  var noun2 string
  fmt.Print("Enter a noun: ")
  fmt.Scan(&noun2)
  fmt.Println(noun2)

  fmt.Print("There once was a ", noun0, " who had a ", adjective0," pet ", noun1,". ", "They went on an adventure to ", location0, ", to recover the ", adjective1," ", noun2,"." )

  var num int
  fmt.Print("Enter a number: ")
  fmt.Scan(&num)
  fmt.Println(num)

  var ber int
  fmt.Print("Enter another number: ")
  fmt.Scan(&ber)
  fmt.Println(ber)
  fmt.Print(num, "+", ber, "=",num + ber, num, "-", ber, "=",num - ber, num, "*", ber, "=", num * ber, num, "/", ber, "=", num / ber)
}