package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const BANNER_LIST = `
   __ ___  __                   x - exit
  / // / |/ /__ _    _____      n - next
 / _  /    / -_) |/|/ (_-<      p - prev
/_//_/_/|_/\__/|__,__/___/      1-9 - details`

type ListMode struct {
	state State
}

func NewListMode(s State) *ListMode {
	return &ListMode{
		state: s,
	}
}

func (l *ListMode) Fetch() {
	ch := make(chan Story)
	var wg sync.WaitGroup

	start := l.state.PageNumber*NUM_PER_PAGE - NUM_PER_PAGE
	end := l.state.PageNumber * NUM_PER_PAGE

	for i, storyId := range l.state.StoryIds[start:end] {
		rank := l.state.PageNumber*NUM_PER_PAGE - 8 + i
		wg.Add(1)
		go func() {
			var story Story

			story.Rank = rank

			defer wg.Done()

			url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", storyId)

			fetchUrl(url, &story)

			ch <- story
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var stories []Story

	for story := range ch {
		stories = append(stories, story)
	}

	l.state.FetchedStories = sortStoriesList(stories)
}

func (l *ListMode) Render() {
	var sB strings.Builder

	if l.state.UiInfo != "" {
		sB.WriteString("\n")
		sB.WriteString(l.state.UiInfo)
		sB.WriteString("\n")
	}

	sB.WriteString(BANNER_LIST)
	sB.WriteString("\n\n")

	for i, story := range l.state.FetchedStories {
		fmt.Fprintf(&sB, "%d | %d. %s\n", i+1, story.Rank, story.Title)
	}

	fmt.Println(sB.String())
}

func (l *ListMode) ChangeState(input string) State {
	l.state.UiInfo = ""

	numbers, _ := regexp.Compile("^[1-9]$")

	switch {
	case input == "x":
		os.Exit(0)
	case input == "n":
		if MAX_STORIES/NUM_PER_PAGE < l.state.PageNumber+1 {
			l.state.PageNumber = MAX_STORIES / NUM_PER_PAGE
			l.state.UiInfo = "Last story on Hacker News"

			return l.state
		}
		l.state.PageNumber += 1
	case input == "p":
		if l.state.PageNumber-1 < 1 {
			l.state.PageNumber = 1
			l.state.UiInfo = "First story on Hacker News"

			return l.state
		}
		l.state.PageNumber -= 1
	case numbers.MatchString(input):
		num, err := strconv.Atoi(input)
		if err != nil {
			l.state.UiInfo = "Error: parsing input number"

			return l.state
		}
		l.state.SelectedStory = l.state.FetchedStories[num-1]
		l.state.Mode = Details
	default:
		l.state.UiInfo = "Error: input not supported"
	}

	return l.state
}
