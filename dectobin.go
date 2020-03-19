//A program that converts decimals to binary

package main

import "fmt"

func main() {
	var input int
	fmt.Println("Input decimal")
	fmt.Scan(&input)                        //Prompts user for input
	fmt.Printf("The binary is %b\n", input) //returns decimal in binary form
}
