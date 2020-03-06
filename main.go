// Programmer name: Jose Padilla
// Date completed:  3/5/2020
// Description: Attempt 1 for Tic Tac Toe and for the value counter

package main
import ("fmt" 
"math/rand" 
"time"
) 
func main() {

//--------------------------------------
//set variables and the o's to fill them  
var line1 [3]string
var line2 [3]string
var line3 [3]string
line1[0] = "O"
line1[1] = "O"
line1[2] = "O"
line2[0] = "O"
line2[1] = "O"
line2[2] = "O"
line3[0] = "O"
line3[1] = "O"
line3[2] = "O"

fmt.Println(line1)
fmt.Println(line2)
fmt.Println(line3)

var change string
//give the user 8 triesw to fill up the entire board
  for i := 1; i <= 9; i++{
    fmt.Println("To change a space on the board please enter a, b, or c for the row followed by the number 1, 2, or 3 for the column.")
    fmt.Scan(&change)
    //changes based on user input
    switch change {
      case "a1" : line1[0] = "x"
      case "a2" : line1[1] = "x"
      case "a3" : line1[2] = "x"
      case "b1" : line2[0] = "x"
      case "b2" : line2[1] = "x"
      case "b3" : line2[2] = "x"
      case "c1" : line3[0] = "x"
      case "c2" : line3[1] = "x"
      case "c3" : line3[2] = "x"
      default : fmt.Println("Please enter a valid combination.")
      }
    fmt.Println(line1)
    fmt.Println(line2)
    fmt.Println(line3)  

//--------------------------------------

s1 := rand.NewSource(time.Now().UnixNano())
r1 := rand.New(s1)

//varcount := make(map[int]int)
//varcount[1] = 0
//fmt.Println(varcount[1])

varcount := make([]int, 50)
    


for i := 1; i <= 50 ; i++{
   varcount[i] = r1.Intn(20)}
   fmt.Println(varcount)



}
