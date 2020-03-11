// Programmer name: Jose Padilla
// Date completed:  3/10/2020
// Description: The game of pig

package main
import ("fmt" 
"math/rand" 
"time"
) 
func main() {

s1 := rand.NewSource(time.Now().UnixNano())
r1 := rand.New(s1)

var playertotal int
var playerscore int
var choice string
fmt.Println("Would you like to roll or stop")
fmt.Scan(&choice)
switch choice {
case roll:
case Roll:
case stop:
case Stop:
}




for playertotal := 0; playertotal <= 100; playertotal++ {
  playerscore = r1.Intn(6)}
  switch playerscore{
  case 0: playerscore = playerscore + r1.Intn(6)
  case 1: playerscore = 0
  case 2: playerscore = playerscore + 2
  case 3: playerscore = playerscore + 3
  case 4: playerscore = playerscore + 4
  case 5: playerscore = playerscore + 5
  case 6: playerscore = playerscore + 6
  }
  playertotal = playertotal + playerscore
fmt.Print(playertotal)

}
