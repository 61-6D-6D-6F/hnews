package internal

import "sort"

func sortStoriesList(stories []Story) []Story {
	sort.Slice(stories, func(left, right int) bool {
		return stories[left].Rank < stories[right].Rank
	})

	return stories
}
