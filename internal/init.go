package internal

func Init() {
	var state State

	state.Mode = List
	state.StoryIds = fetchStoryIds()
	state.PageNumber = 1

	var screen Screen
	var display *Display

	for {
		switch state.Mode {
		case List:
			screen = NewListMode(state)
		case Details:
			screen = NewDetailsMode(state)
		case Comments:
			screen = NewCommentsMode(state)
		}

		display = NewDisplay(screen)

		display.Fetch()
		display.Render()

		input := scan()

		state = display.ChangeState(input)
	}
}
