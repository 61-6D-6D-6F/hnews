package stories

import (
	"fmt"
	"os"
)

func scanInput(currentPage *int) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error scanning input")
		return
	}

	switch input {
	case "x":
		os.Exit(0)
	case "n":
		if 500/9 < *currentPage+1 {
			*currentPage = 500 / 9
		} else {
			*currentPage += 1
		}
	case "p":
		if *currentPage-1 < 1 {
			*currentPage = 1
		} else {
			*currentPage -= 1
		}
	default:
		fmt.Println("Error: input not supported")
	}
}
