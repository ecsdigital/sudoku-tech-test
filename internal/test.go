package internal

import (
	"strings"
	"fmt"
	"io"
)

type Response struct {
	Result	string
	Grid	[][]int
}

type Expected struct {
	TestName		string
	httpStatusCode	int
	Response		string
}

func OkRequests(apiRestart int) {
	expectedHttpStatusCode := 200
	const numberOfTests int = 7
	var payload [numberOfTests]io.Reader
	var Expected [numberOfTests]Expected
	payload[0] = strings.NewReader(`{"grid":[[0,0,0,2,6,0,7,0,1],[6,8,0,0,7,0,0,9,0],[1,9,0,0,0,4,5,0,0],[8,2,0,1,0,0,0,4,0],[0,0,4,6,0,2,9,0,0],[0,5,0,0,0,3,0,2,8],[0,0,9,3,0,0,0,7,4],[0,4,0,0,5,0,0,3,6],[7,0,3,0,1,8,0,0,0]]}`)
	payload[1] = strings.NewReader(`{"grid":[[0,2,0,6,0,8,0,0,0],[5,8,0,0,0,9,7,0,0],[0,0,0,0,4,0,0,0,0],[3,7,0,0,0,0,5,0,0],[6,0,0,0,0,0,0,0,4],[0,0,8,0,0,0,0,1,3],[0,0,0,0,2,0,0,0,0],[0,0,9,8,0,0,0,3,6],[0,0,0,3,0,6,0,9,0]]}`)
	payload[2] = strings.NewReader(`{"grid":[[2,0,0,3,0,0,0,0,0],[8,0,4,0,6,2,0,0,3],[0,1,3,8,0,0,2,0,0],[0,0,0,0,2,0,3,9,0],[5,0,7,0,0,0,6,2,1],[0,3,2,0,0,6,0,0,0],[0,2,0,0,0,9,1,4,0],[6,0,1,2,5,0,8,0,9],[0,0,0,0,0,1,0,0,2]]}`)
	payload[3] = strings.NewReader(`{"grid":[[0,2,0,0,0,0,0,0,0],[0,0,0,6,0,0,0,0,3],[0,7,4,0,8,0,0,0,0],[0,0,0,0,0,3,0,0,2],[0,8,0,0,4,0,0,1,0],[6,0,0,5,0,0,0,0,0],[0,0,0,0,1,0,7,8,0],[5,0,0,0,0,9,0,0,0],[0,0,0,0,0,0,0,4,0]]}`)
	payload[4] = strings.NewReader(`{"Grid":[[4,3,5,2,6,9,7,8,1],[6,8,2,5,7,1,4,9,3],[1,9,7,8,3,4,5,6,2],[8,2,6,1,9,5,3,4,7],[3,7,4,6,8,2,9,1,5],[9,5,1,7,4,3,6,2,8],[5,1,9,3,2,6,8,7,4],[2,4,8,9,5,7,1,3,6],[7,6,3,4,1,8,2,5,9]]}`)
	payload[5] = strings.NewReader(`{"grid":[[0,0,0,2,6,0,7,0,1],
	[6,8,0,0,7,0,0,9,0],
	[1,9,0,0,0,4,5,0,0],
	[8,2,0,1,0,0,0,4,0],
	[0,0,4,6,0,2,9,0,0],
	[0,5,0,0,0,3,0,2,8],
	[0,0,9,3,0,0,0,7,4],
	[0,4,0,0,5,0,0,3,6],
	[7,0,3,0,1,8,0,0,0]]
}`)
	payload[6] = strings.NewReader(`{"grid":[[0, 0 , 0, 2, 6, 0, 7, 0, 1], [6, 8, 0, 0, 7, 0, 0, 9, 0], [1, 9, 0, 0, 0, 4, 5, 0, 0], [8, 2, 0, 1, 0, 0, 0, 4, 0], [0, 0, 4, 6, 0, 2, 9, 0, 0], [0, 5, 0, 0, 0, 3, 0, 2, 8], [0, 0, 9, 3, 0, 0, 0, 7, 4], [0, 4, 0, 0, 5, 0, 0, 3, 6], [7, 0, 3, 0, 1, 8, 0, 0, 0]]}`)

	Expected[0].Response = `{"Result":"Success","Grid":[[4,3,5,2,6,9,7,8,1],[6,8,2,5,7,1,4,9,3],[1,9,7,8,3,4,5,6,2],[8,2,6,1,9,5,3,4,7],[3,7,4,6,8,2,9,1,5],[9,5,1,7,4,3,6,2,8],[5,1,9,3,2,6,8,7,4],[2,4,8,9,5,7,1,3,6],[7,6,3,4,1,8,2,5,9]]}`
	Expected[1].Response = `{"Result":"Success","Grid":[[1,2,3,6,7,8,9,4,5],[5,8,4,2,3,9,7,6,1],[9,6,7,1,4,5,3,2,8],[3,7,2,4,6,1,5,8,9],[6,9,1,5,8,3,2,7,4],[4,5,8,7,9,2,6,1,3],[8,3,6,9,2,4,1,5,7],[2,1,9,8,5,7,4,3,6],[7,4,5,3,1,6,8,9,2]]}`
	Expected[2].Response = `{"Result":"Success","Grid":[[2,7,6,3,1,4,9,5,8],[8,5,4,9,6,2,7,1,3],[9,1,3,8,7,5,2,6,4],[4,6,8,1,2,7,3,9,5],[5,9,7,4,3,8,6,2,1],[1,3,2,5,9,6,4,8,7],[3,2,5,7,8,9,1,4,6],[6,4,1,2,5,3,8,7,9],[7,8,9,6,4,1,5,3,2]]}`
	Expected[3].Response = `{"Result":"Success","Grid":[[1,2,6,4,3,7,9,5,8],[8,9,5,6,2,1,4,7,3],[3,7,4,9,8,5,1,2,6],[4,5,7,1,9,3,8,6,2],[9,8,3,2,4,6,5,1,7],[6,1,2,5,7,8,3,9,4],[2,6,9,3,1,4,7,8,5],[5,4,8,7,6,9,2,3,1],[7,3,1,8,5,2,6,4,9]]}`
	Expected[4].Response = `{"Result":"Success","Grid":[[4,3,5,2,6,9,7,8,1],[6,8,2,5,7,1,4,9,3],[1,9,7,8,3,4,5,6,2],[8,2,6,1,9,5,3,4,7],[3,7,4,6,8,2,9,1,5],[9,5,1,7,4,3,6,2,8],[5,1,9,3,2,6,8,7,4],[2,4,8,9,5,7,1,3,6],[7,6,3,4,1,8,2,5,9]]}`
	Expected[5].Response = `{"Result":"Success","Grid":[[4,3,5,2,6,9,7,8,1],[6,8,2,5,7,1,4,9,3],[1,9,7,8,3,4,5,6,2],[8,2,6,1,9,5,3,4,7],[3,7,4,6,8,2,9,1,5],[9,5,1,7,4,3,6,2,8],[5,1,9,3,2,6,8,7,4],[2,4,8,9,5,7,1,3,6],[7,6,3,4,1,8,2,5,9]]}`
	Expected[6].Response = `{"Result":"Success","Grid":[[4,3,5,2,6,9,7,8,1],[6,8,2,5,7,1,4,9,3],[1,9,7,8,3,4,5,6,2],[8,2,6,1,9,5,3,4,7],[3,7,4,6,8,2,9,1,5],[9,5,1,7,4,3,6,2,8],[5,1,9,3,2,6,8,7,4],[2,4,8,9,5,7,1,3,6],[7,6,3,4,1,8,2,5,9]]}`

	Expected[0].TestName = "Solvable grid"
	Expected[1].TestName = "Solvable grid"
	Expected[2].TestName = "Solvable grid"
	Expected[3].TestName = "Solvable grid"
	Expected[4].TestName = "Completed grid"
	Expected[5].TestName = "New Line grid"
	Expected[6].TestName = "Additional spaces grid"

	errorCount := 0
	var errorIndices[numberOfTests] int
	for i:=0; i<numberOfTests; i++ {
		if apiRestart == 1 {
			var input string = "n"
			for input != "y" {
				fmt.Println("TestRunner with API restart, please enter y when the candidate's program is ready to check")
				fmt.Scanln(&input)
			}
		}
		httpStatusCode, response, err := PostSudoku(payload[i])
		if err !=nil {
			fmt.Println("Unexpected error")
			errorIndices[i] = 1
			errorCount++
			continue
		}
		if httpStatusCode != expectedHttpStatusCode {
			errorIndices[i] = 1
			errorCount++
			continue
		}
		response = strings.Replace(response, "\n", "", -1)
		response = strings.Replace(response, " ", "", -1)
		if response != Expected[i].Response {
			fmt.Printf("Error, expected %v \n response %v", Expected[i].Response, response)
			errorIndices[i] = 1
			errorCount++
		}
	}

	fmt.Printf("Ok Requests tests %v, failures %v\n", numberOfTests, errorCount)
	for j:=0; j<numberOfTests; j++ {
		if errorIndices[j] == 1 {
			fmt.Printf("     %v request test failure\n", Expected[j].TestName)
		}
	}
}


func UnsolvableRequests(apiRestart int) {
	const numberOfTests int = 1
	var payload [numberOfTests]io.Reader
	var Expected [numberOfTests]Expected
	payload[0] = strings.NewReader(`{"grid":[[0,0,9,0,2,8,7,0,0],[8,0,6,0,0,4,0,0,5],[0,0,3,0,0,0,0,0,4],[6,0,0,0,0,0,0,0,0],[0,2,0,7,1,3,4,5,0],[0,0,0,0,0,0,0,0,2],[3,0,0,0,0,0,5,0,0],[9,0,0,4,0,0,8,0,7],[0,0,1,2,5,0,3,0,0]]}`)

	Expected[0].Response = `{"Result":"Failure"`
	Expected[0].TestName = "Unolvable grid"

	errorCount := 0
	var errorIndices[numberOfTests] int
	for i:=0; i<numberOfTests; i++ {
		if apiRestart == 1 {
			var input string = "n"
			for input != "y" {
				fmt.Println("TestRunner with API restart, please enter y when the candidate's program is ready to check")
				fmt.Scanln(&input)
			}
		}
		_, response, err := PostSudoku(payload[i])
		if err !=nil {
			fmt.Println("Unexpected error")
			errorIndices[i] = 1
			errorCount++
			continue
		}

		response = strings.Replace(response, "\n", "", -1)
		response = strings.Replace(response, " ", "", -1)
		response = response[0:19]
		if response != Expected[i].Response {
			fmt.Printf("Error, expected beginning of response %v \n actual %v", Expected[i].Response, response)
			errorIndices[i] = 1
			errorCount++
		}
	}

	fmt.Printf("Unsolvable Requests tests %v, failures %v\n", numberOfTests, errorCount)
	for j:=0; j<numberOfTests; j++ {
		if errorIndices[j] == 1 {
			fmt.Printf("     %v request test failure\n", Expected[j].TestName)
		}
	}
}

func BadRequests(apiRestart int) { 
	expectedHttpStatusCode := 400
	const numberOfTests int = 9
	var payload [numberOfTests]io.Reader
	var Expected [numberOfTests]Expected
	payload[0] = strings.NewReader(`{"grd":[[0,0,9,0,7,0,0,0,5],[0,0,2,1,0,0,9,0,0],[1,0,0,0,2,8,0,0,0],[0,7,0,0,0,5,0,0,1],[0,0,8,5,1,0,0,0,0],[0,5,0,0,0,0,3,0,0],[0,0,0,0,0,3,0,0,6],[8,0,0,0,0,0,0,0,0],[2,1,0,0,0,0,0,8,7]]}`)
	payload[1] = strings.NewReader(`{"grid":[[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0]]}`)
	payload[2] = strings.NewReader(`{"grid":[[0,0,0,2,6,0,7,0],[6,8,0,0,7,0,0,9,0],[1,9,0,0,0,4,5,0,0],[8,2,0,1,0,0,0,4],[0,0,4,6,0,2,9,0,0],[0,5,0,0,0,3,0,2,8],[0,0,9,3,0,0,0,7,4],[0,4,0,0,5,0,0,3,6],[7,0,3,0,1,8,0,0,0]]}`)
	payload[3] = strings.NewReader(`{"grid":[[0,2,0,6,0,8,0,0,0],[5,8,0,0,0,9,7,0,0],[0,0,0,0,4,0,0,0,0],[3,7,0,0,0,0,5,0,0],[6,0,0,0,0,0,0,0,4],[0,0,8,0,0,0,0,1,3],[0,0,0,0,2,0,0,0,0],[0,0,9,8,0,0,0,3,6],[0,0,0,3,0,6,0,9,0]]}`)
	payload[4] = strings.NewReader(`{"grid":[[0,0,9,0,7,0,0,0,5],[0,0,2,1,0,0,9,0,0],[1,0,0,0,2,8,0,0,0],[0,7,0,0,0,5,0,0,1],[0,0,8,5,1,0,0,0,0],[0,5,0,0,0,0,3,0,0],[0,0,0,0,0,3,0,0,6],[8,0,0,0,0,0,0,0,0],[2,1,0,0,0,0,0,8,7]]}`)
	payload[5] = strings.NewReader(`{"grid":[[2,0,0,-3,0,0,0,0,0],[8,0,4,0,6,2,0,0,3],[0,1,3,8,0,0,2,0,0],[0,0,0,0,2,0,3,9,0],[5,0,7,0,0,0,6,2,1],[0,3,2,0,0,6,0,0,0],[0,2,0,0,0,9,1,4,0],[6,0,1,2,5,0,8,0,9],[0,0,0,0,0,1,0,0,2]]}`)
	payload[6] = strings.NewReader(`{"grid":[[2,0,0,37,0,0,0,0,0],[8,0,4,0,6,2,0,0,3],[0,1,3,8,0,0,2,0,0],[0,0,0,0,2,0,3,9,0],[5,0,7,0,0,0,6,2,1],[0,3,2,0,0,6,0,0,0],[0,2,0,0,0,9,1,4,0],[6,0,1,2,5,0,8,0,9],[0,0,0,0,0,1,0,0,2]]}`)
	payload[7] = strings.NewReader(`{"grid":[[2,0,0,3.27,0,0,0,0,0],[8,0,4,0,6,2,0,0,3],[0,1,3,8,0,0,2,0,0],[0,0,0,0,2,0,3,9,0],[5,0,7,0,0,0,6,2,1],[0,3,2,0,0,6,0,0,0],[0,2,0,0,0,9,1,4,0],[6,0,1,2,5,0,8,0,9],[0,0,0,0,0,1,0,0,2]]}`)
	payload[8] = strings.NewReader(`{"grid":[[A,0,0,3,0,0,0,0,0],[8,0,4,0,6,2,0,0,3],[0,1,3,8,0,0,2,0,0],[0,0,0,0,2,0,3,9,0],[5,0,7,0,0,0,6,2,1],[0,3,2,0,0,6,0,0,0],[0,2,0,0,0,9,1,4,0],[6,0,1,2,5,0,8,0,9],[0,0,0,0,0,1,0,0,2]]}`)

	Expected[0].TestName = "Misspelling of grid in json body"
	Expected[1].TestName = "Wrong grid size"
	Expected[2].TestName = "Wrong grid size"
	Expected[3].TestName = "Wrong grid size"
	Expected[4].TestName = "Invalid grid, duplicate 5 in centre 3x3 box"
	Expected[5].TestName = "Negative number in grid"
	Expected[6].TestName = "Double digit number in grid"
	Expected[7].TestName = "Decimal in grid"
	Expected[8].TestName = "Alphabet character in grid"

	errorCount := 0
	var errorIndices[numberOfTests] int
	for i:=0; i<numberOfTests; i++ {
		if apiRestart == 1 {
			var input string = "n"
			for input != "y" {
				fmt.Println("TestRunner with API restart, please enter y when the candidate's program is ready to check")
				fmt.Scanln(&input)
			}
		}
		httpStatusCode, _, err := PostSudoku(payload[i])
		if err !=nil {
			fmt.Println("Unexpected error")
			errorIndices[i] = 1
			errorCount++
			continue
		}
		if httpStatusCode != expectedHttpStatusCode {
			errorIndices[i] = 1
			errorCount++
			continue
		}
	}

	fmt.Printf("Bad Requests tests %v, failures %v\n", numberOfTests, errorCount)
	for j:=0; j<numberOfTests; j++ {
		if errorIndices[j] == 1 {
			fmt.Printf("     %v request test failure\n", Expected[j].TestName)
		}
	}
}

