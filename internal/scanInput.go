package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func scanStoryListInput(state State) State {
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println()
		fmt.Println("Error: scanning input")
		return state
	}

	numbers, _ := regexp.Compile("^[1-9]$")

	switch {
	case input == "x":
		os.Exit(0)
	case input == "n":
		if MAX_STORIES/NUM_PER_PAGE < state.PageNumber+1 {
			state.PageNumber = MAX_STORIES / NUM_PER_PAGE
			fmt.Println()
			fmt.Println("This is the last story")
		} else {
			state.PageNumber += 1
		}
	case input == "p":
		if state.PageNumber-1 < 1 {
			state.PageNumber = 1
			fmt.Println()
			fmt.Println("This is the first story")
		} else {
			state.PageNumber -= 1
		}
	case numbers.MatchString(input):
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println()
			fmt.Println("Error: parsing input number")
		} else {
			state.SelectedStory = state.FetchedStories[num-1]
			state.Mode = Details
		}
	default:
		fmt.Println()
		fmt.Println("Error: input not supported")
	}

	return state
}

func scanStoryDetailsInput(state State) State {
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println()
		fmt.Println("Error: scanning input")
		return state
	}

	switch input {
	case "x":
		os.Exit(0)
	case "b":
		state.Mode = List
	case "c":
		if len(state.SelectedStory.Kids) == 0 {
			fmt.Println()
			fmt.Println("Story has no comment yet")
		} else {
			state.CurrentSiblings = state.SelectedStory.Kids
			state.HistorySiblings = [][]int{}
			state.HistoryPos = []int{}
			state.CurrentPos = 0
			state.Mode = Comments
		}
	default:
		fmt.Println()
		fmt.Println("Error: input not supported")
	}

	return state
}

func scanCommentInput(state State) State {
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println()
		fmt.Println("Error: scanning input")
		return state
	}

	switch input {
	case "x":
		os.Exit(0)
	case "b":
		if len(state.HistorySiblings) == 0 {
			state.Mode = Details
		} else {
			state.CurrentSiblings = state.HistorySiblings[len(state.HistorySiblings)-1]
			state.CurrentPos = state.HistoryPos[len(state.HistoryPos)-1]

			state.HistorySiblings = state.HistorySiblings[:len(state.HistorySiblings)-1]
			state.HistoryPos = state.HistoryPos[:len(state.HistoryPos)-1]
		}
	case "r":
		if len(state.FetchedComment.Kids) == 0 {
			fmt.Println()
			fmt.Println("Comment has no reply yet")
		} else {
			state.HistorySiblings = append(state.HistorySiblings, state.CurrentSiblings)
			state.HistoryPos = append(state.HistoryPos, state.CurrentPos)
			state.CurrentSiblings = state.FetchedComment.Kids
			state.CurrentPos = 0
		}
	case "n":
		if len(state.CurrentSiblings)-1 < state.CurrentPos+1 {
			state.CurrentPos = len(state.CurrentSiblings) - 1
			fmt.Println()
			fmt.Println("This is the last comment of reply chain")
		} else {
			state.CurrentPos += 1
		}
	case "p":
		if state.CurrentPos-1 < 0 {
			state.CurrentPos = 0
			fmt.Println()
			fmt.Println("This is the first comment of reply chain")
		} else {
			state.CurrentPos -= 1
		}
	default:
		fmt.Println()
		fmt.Println("Error: input not supported")
	}

	return state
}
