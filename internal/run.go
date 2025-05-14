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

	switch mode {
	case List:
		runStoriesList(topIds, currentPage)
	case Details:
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error parsing input number")
			mode = List
			runStoriesList(topIds, currentPage)
		} else {
			runStoryDetails(topIds, sortedStories[num-1], currentPage)
		}
	}
}

func runStoryDetails(topIds []int, story Story, currentPage *int) {
	renderStoryDetails(story)

	scanStoryDetailsInput()

	switch mode {
	case List:
		runStoriesList(topIds, currentPage)
	case Details:
		runStoryDetails(topIds, story, currentPage)
	case Comments:
		parentList := []string{"story"}
		siblingList := [][]int{}
		siblingPosition := []int{}
		runComment(topIds, story, *&currentPage, story.Kids, 0, parentList, siblingList, siblingPosition)
	}
}

func runComment(topIds []int, story Story, currentPage *int, kids []int,
	currentComment int, parentList []string, siblingList [][]int, siblingPosition []int) {

	fetchedComment := fetchComment(kids[currentComment])

	renderComment(fetchedComment)

	input, newId := scanCommentInput(fetchedComment, kids, currentComment)

	switch {
	case input == "b":
		runParent(topIds, story, currentPage, kids, newId, parentList, siblingList, siblingPosition)
	case input == "r":
		runComment(topIds, story, currentPage, kids, newId, parentList, siblingList, siblingPosition)
	case input == "replies":
		parentList = append(parentList, "comment")
		siblingList = append(siblingList, kids)
		siblingPosition = append(siblingPosition, currentComment)
		runComment(topIds, story, currentPage, fetchedComment.Kids, newId, parentList, siblingList, siblingPosition)
	case input == "n" || input == "p":
		runComment(topIds, story, currentPage, kids, newId, parentList, siblingList, siblingPosition)
	default:
		runComment(topIds, story, currentPage, kids, newId, parentList, siblingList, siblingPosition)
	}
}

func runParent(topIds []int, story Story, currentPage *int, kids []int,
	newId int, parentList []string, siblingList [][]int, siblingPosition []int) {

	if len(parentList) == 1 {
		mode = Details
		runStoryDetails(topIds, story, currentPage)
	} else {
		kids = siblingList[len(parentList)-2]
		newId = siblingPosition[len(parentList)-2]

		parentList = parentList[:len(parentList)-1]
		siblingList = siblingList[:len(parentList)-1]
		siblingPosition = siblingPosition[:len(parentList)-1]

		runComment(topIds, story, currentPage, kids, newId, parentList, siblingList, siblingPosition)
	}
}
