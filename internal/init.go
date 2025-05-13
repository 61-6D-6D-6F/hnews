package internal

func Init() {
	runStoriesList(fetchTopStoriesIds(), &currentPage)
}
