package internal

func Init() {
	var state State

	state.Mode = List
	state.StoryIds = fetchStoryIds()
	state.PageNumber = 1

	runStoryList(state)
}
