package main

import 
"fmt"

  
func main() {
  fmt.Println("This game is called the 99 trick.")

  fmt.Print("Enter a number between 10 and 49: ")
  var answer float64
  fmt.Scanf("%f", &answer)

  factor := 99 - answer

  fmt.Println("Yor factor is",factor)

  fmt.Println("Now enter a number between 50 and 99")
  var friend float64
  fmt.Scanf("%f", &friend)

  factor2 := factor + friend

  fmt.Println("Now take the hundreds from",factor2, "and add it to the resulting number.")
  var result float64
  fmt.Scanf("%f", &result)
  end := friend - result
  fmt.Println("Was this your number",end)















}

