package internal

import (
	"fmt"
	"os"
	"regexp"
)

func scanStoriesListInput(currentPage *int) string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error scanning input")
		return input
	}

	num, _ := regexp.Compile("[1-9]")

	switch {
	case input == "x":
		os.Exit(0)
	case input == "n":
		if 500/9 < *currentPage+1 {
			*currentPage = 500 / 9
		} else {
			*currentPage += 1
		}
	case input == "p":
		if *currentPage-1 < 1 {
			*currentPage = 1
		} else {
			*currentPage -= 1
		}
	case num.MatchString(input):
		state = Detail
	default:
		fmt.Println("Error: input not supported")
	}

	return input
}

func scanStoryDetailsInput() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error scanning input")
	}

	switch input {
	case "x":
		os.Exit(0)
	case "b":
		state = List
	default:
		fmt.Println("Error: input not supported")
	}
}
