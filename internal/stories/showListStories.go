package stories

func ShowList(topIds []int, currentPage *int) {
	fetchedStories := fetchStories(topIds, *currentPage)

	sortedStories := sortList(fetchedStories)

	renderList(sortedStories)

	scanInput(*&currentPage)

	ShowList(topIds, currentPage)
}
