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
			screen = NewListMenu(state)
		case Details:
			screen = NewDetailsMenu(state)
		case Comments:
			screen = NewCommentMenu(state)
		}

		display = NewDisplay(screen)

		display.Fetch()
		display.Render()
		state = display.Scan()
	}
}
