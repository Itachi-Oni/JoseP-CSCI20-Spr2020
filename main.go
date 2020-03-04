// Programmer name: Jose Padilla
// Date completed:  3/3/2020
// Description: You Code It 2.5.1

package main

import "fmt"

func main() {
    //Create a string array of at least 5 values.  It should hold 5 products to buy
    var products [5]string
    products[0] = "Soda"
    products[1] = "Chips"
    products[2] = "Candy"
    products[3] = "Ice Cream"
    products[4] = "Cake"

    fmt.Println(products)

    //Create a float array of at least 5 values.  It should hold a price for each of the products
var prices [5]float64
    prices[0] = 1.50
    prices[1] = 1.00
    prices[2] = 0.99
    prices[3] = 4.99
    prices[4] = 9.99

    fmt.Println(prices)

    //Iterate through the array and output the product and it's price (similar to a menu)
    fmt.Println(products[0],prices[0],products[1],prices[1],products[2],prices[2],products[3],prices[3],products[4],prices[4],)
}