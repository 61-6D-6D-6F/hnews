package internal

const NUM_PER_PAGE = 9
const MAX_STORIES = 500

type State struct {
	Mode            Mode
	StoryIds        []int
	PageNumber      int
	FetchedStories  []Story
	SelectedStory   Story
	FetchedComment  Comment
	CurrentSiblings []int
	CurrentPos      int
	HistorySiblings [][]int
	HistoryPos      []int
	UiInfo          string
}
