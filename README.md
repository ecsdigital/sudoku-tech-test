# Sudoku Tech Test (golang software developer)

Please create a clone of this repo on your local machine to faciliate testing. The files in the main folder relate to the code to auto-check the candidate's solution.

# The Question

Write a JSON API in Golang to Solve Sudoku

Hint: There is a recursive algorithm that solves sudoko puzzles

The API should expect a POST request to the following endpoint: http://localhost:8080/sudoku

The body of the POST request will be of the form

{"Grid":[[0,2,0,0,0,0,0,0,0],[0,0,0,6,0,0,0,0,3],[0,7,4,0,8,0,0,0,0],[0,0,0,0,0,3,0,0,2],[0,8,0,0,4,0,0,1,0],[6,0,0,5,0,0,0,0,0],[0,0,0,0,1,0,7,8,0],[5,0,0,0,0,9,0,0,0],[0,0,0,0,0,0,0,4,0]]

The API responses should be of the following forms

{"Result":"Success,"Grid":[1,2,6,4,3,7,9,5,8],[8,9,5,6,2,1,4,7,3],[3,7,4,9,8,5,1,2,6],[4,5,7,1,9,3,8,6,2],[9,8,3,2,4,6,5,1,7],[6,1,2,5,7,8,3,9,4],[2,6,9,3,1,4,7,8,5],[5,4,8,7,6,9,2,3,1],[7,3,1,8,5,2,6,4,9]]}

{"Result":"Failure,"Grid":[]}

Use the HTTP status code 400 for handling incorrect requests

# Auto-checking the Answer

## Set up the API Server

The code in this repo is for testing the answer to the above question.

Once the code has been sent by the candidate, please briefly review the code. For example, confirm that it is written in the golang software language and contains the file main.go

Follow the candidates instructions to run the code, e.g. from inside the cmd directory

$go run main.go

This should spin up the API server.

Please confirm the server is running on localhost:8080 by:
1. Reading the candidates instructions
2. Visit http://localhost:8080/sudoku in a browser window and confirm there is a server response

## Running the automated tests against the server

Once the Sudoku API Server is up and running, the automated tests can be run.

Note, the test are designed to run either fast or slow.

Initially they will run fast, however if they cause the API to crash they will automatically re-run slowly. In the slow case the API needs to be re-started (if crashed) before each test.

The location of the API for the tests is found in data/data.go here:

const Url string = "http://localhost:8080/sudoku"

Execute all of the test with the following from the cmd directory

$go run main.go

This will run all tests and give a report of the test outcomes.

Please note the tests are in three groups

1. OkRequests (Solvable requests)
2. UnsolvableRequests (Some combinations of allowable initial numbers in an unsolved Sudoku Grid result in no possible solution)
3. BadRequests (These includes a combination of the following deliberate errors - please see internal/test.go for further details):
    * "Misspelling of grid in json body"
    * "Wrong grid size"
	* "Invalid grid, duplicate 5 in centre 3x3 box"
	* "Negative number in grid"
	* "Double digit number in grid"
	* "Decimal in grid"
	* "Alphabet character in grid"


