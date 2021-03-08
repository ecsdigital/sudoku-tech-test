package main

import (
	"../internal"
	"../data"
	"fmt"
)

func main() {
	Ready()
	apiRestart := 0
	defer func() {
		if r:= recover(); r !=nil {
			fmt.Println("the candidates API crashed from invalid input, rerunning all tests with API restarts")
			apiRestart = 1
			TestRunner(1)
		}
	}()

	TestRunner(0)
}

func Ready() {
	fmt.Println("This program performs automated checks against an interview candidate's golang JSON API Sudoku Solver")
	fmt.Println("1. Please run the candidate's program with the command $go run main.go in the relevant directory")
	fmt.Printf("2. Please visit %v in a browser to confirm the server is running", data.Url)
	fmt.Println("")
	var input string = "n"
	for input != "y" {
		fmt.Println("Please enter y when the candidate's program is ready to check")
		fmt.Scanln(&input)
	}
	fmt.Printf("\n\n")
}

func TestRunner(apiRestart int) {
	internal.OkRequests(apiRestart)
	internal.UnsolvableRequests(apiRestart)
	internal.BadRequests(apiRestart)	
}
