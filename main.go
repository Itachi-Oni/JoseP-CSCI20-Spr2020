// Programmer name: Jose Padilla
// Date completed:  2/20/2020
// Description: Project 1: A temperature converter.

package main
import "fmt"
func main() {
 var x [4]int
 var user0 int
 var user1 int
 var user2 int
 var user3 int

  fmt.Println("Please enter a single digit number. 1/4 ")
  fmt.Scanln(&user0)
  fmt.Println(user0)
  x[0] = user0

  fmt.Println("Please enter a single digit number. 2/4 ")
  fmt.Scanln(&user1)
  fmt.Println(user1)
  x[1] = user1

    fmt.Println("Please enter a single digit number. 3/4 ")
  fmt.Scanln(&user2)
  fmt.Println(user2)
  x[2] = user2

    fmt.Println("Please enter a single digit number. 4/4 ")
  fmt.Scanln(&user3)
  fmt.Println(user3)
  x[3] = user3

  fmt.Println(x)
  added := user0 + user1 + user2 + user3
  fmt.Println(added)
for added <= 9 && added >= 1{
  added = user0 + user1 + user2 + user3
  fmt.Println(added)
  }
} 