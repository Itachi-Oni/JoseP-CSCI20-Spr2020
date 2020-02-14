//Created by: Jose Padilla
//Created on: 2/13/2020
//Description: This program accepts user input and converts it into Kelvin and Celsius from Farenheit.
  
package main
import "fmt"
func main() {
  var far float64
  fmt.Println("This program is used to convert Farenheit into Celsius and Kelvin.")
  fmt.Print("Please enter a temperature in Farenheit.")
  fmt.Scan( &far)
  kel := (far-32) * 5/9 + 273.15
  cel := (far - 32) * 5/9
  fmt.Print("Farenheit: ", far," Kelvin: ", kel, " Celsius: ", cel)
}