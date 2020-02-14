//Created by: Jose Padilla
//Created on: 2/13/2020
//Description: This program accepts user input and converts it into Kelvin and Celsius from Farenheit.
  
package main
import "fmt"
func main() {
  //Sets user input variable
  var far float64
  fmt.Println("This program is used to convert Farenheit into Celsius and Kelvin.")
  fmt.Print("Please enter a temperature in Farenheit.")
  //Scans for user input to assign to variable
  fmt.Scan( &far)
  //Converts Farenheit to Kelvin
  kel := (far-32) * 5/9 + 273.15
  //Converts Farenheit to Celsius
  cel := (far - 32) * 5/9
  //Prints the results
  fmt.Print("Farenheit: ", far," Kelvin: ", kel, " Celsius: ", cel)
}