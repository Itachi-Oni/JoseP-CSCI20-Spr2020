// Programmer name: Jose Padilla
// Date completed:  2/20/2020
// Description: Project 1: A temperature converter.

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