package stories

import (
	"fmt"
	"strconv"
)

func ShowList(topIds []int, currentPage *int) {
	fetchedStories := fetchStories(topIds, *currentPage)

	sortedStories := sortList(fetchedStories)

	renderList(sortedStories)

	input := scanInput(*&currentPage)

	switch Mode {
	case List:
		ShowList(topIds, currentPage)
	case Detail:
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error parsing input number")
			Mode = List
			ShowList(topIds, currentPage)
		} else {
			showDetail(topIds, sortedStories[num-1], currentPage)
		}
	}
}

func showDetail(topIds []int, story Story, currentPage *int) {
	renderDetail(story)

	scanDetailInput()

	switch Mode {
	case List:
		ShowList(topIds, currentPage)
	default:
		showDetail(topIds, story, currentPage)
	}
}
