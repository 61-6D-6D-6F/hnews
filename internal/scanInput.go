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
			fmt.Println("This is the last story")
		} else {
			*currentPage += 1
		}
	case input == "p":
		if *currentPage-1 < 1 {
			*currentPage = 1
			fmt.Println("This is the first story")
		} else {
			*currentPage -= 1
		}
	case num.MatchString(input):
		mode = Details
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
		mode = List
	case "c":
		mode = Comments
	default:
		fmt.Println("Error: input not supported")
	}
}

func scanCommentInput(comment Comment, storyKids []int, id int) (string, int) {
	var input string

	newId := id

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error scanning input")
		return input, newId
	}

	switch input {
	case "x":
		os.Exit(0)
	case "b":
	case "r":
		if len(comment.Kids) != 0 {
			newId = 0
			input = "replies"
		} else {
			fmt.Println("Comment has no reply yet")
		}
	case "n":
		if len(storyKids)-1 < id+1 {
			newId = len(storyKids) - 1
			fmt.Println("This is the last comment of reply chain")
		} else {
			newId += 1
		}
	case "p":
		if newId-1 < 0 {
			newId = 0
			fmt.Println("This is the first comment of reply chain")
		} else {
			newId -= 1
		}
	default:
		fmt.Println("Error: input not supported")
	}

	return input, newId
}
