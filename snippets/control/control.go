package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument.")
		return
	}
	argument := os.Args[1]

	// With expression after switch
	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("Two, Three, or Four!")
		fallthrough
		default:
			fmt.Println("Value:", argument)
			_, err := strconv.Atoi(argument)
			if err != nil {
				fmt.Println("cannot convert to integer:", argument)
			}
			
			return 
	}
}
