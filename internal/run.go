package internal

import (
	"fmt"
	"strconv"
)

func runStoriesList(topIds []int, currentPage *int) {
	fetchedStories := fetchStories(topIds, *currentPage)

	sortedStories := sortStoriesList(fetchedStories)

	renderStoriesList(sortedStories)

	input := scanStoriesListInput(*&currentPage)

	switch state {
	case List:
		runStoriesList(topIds, currentPage)
	case Detail:
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error parsing input number")
			state = List
			runStoriesList(topIds, currentPage)
		} else {
			runStoryDetails(topIds, sortedStories[num-1], currentPage)
		}
	}
}

func runStoryDetails(topIds []int, story Story, currentPage *int) {
	renderStoryDetails(story)

	scanStoryDetailsInput()

	switch state {
	case List:
		runStoriesList(topIds, currentPage)
	case Detail:
		runStoryDetails(topIds, story, currentPage)
	}
}
