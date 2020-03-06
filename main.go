// Programmer name: Jose Padilla
// Date completed:  3/3/2020
// Description: You Code It 2.5.1

package main

import (
  "fmt" 
"math/rand" 
)

func main() {
  
 var array [100]int
 
 
 for i:= 0; i < 100 ; i++{
   array[i] = rand.Intn(1000)}
   fmt.Println(array)
}