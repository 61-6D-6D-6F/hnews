package internal

func runStoryList(state State) {
	state.FetchedStories = fetchStories(state.StoryIds, state.PageNumber)

	renderStoryList(state.FetchedStories)

	state = scanStoryListInput(state)

	switch state.Mode {
	case List:
		runStoryList(state)
	case Details:
		runStoryDetails(state)
	}
}

func runStoryDetails(state State) {
	renderStoryDetails(state.SelectedStory)

	state = scanStoryDetailsInput(state)

	switch state.Mode {
	case List:
		runStoryList(state)
	case Details:
		runStoryDetails(state)
	case Comments:
		runComment(state)
	}
}

func runComment(state State) {
	state.FetchedComment = fetchComment(state.CurrentSiblings[state.CurrentPos])

	renderComment(state.FetchedComment)

	state = scanCommentInput(state)

	switch state.Mode {
	case Details:
		runStoryDetails(state)
	case Comments:
		runComment(state)
	}
}
