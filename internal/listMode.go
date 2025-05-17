package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
		go fetchStory(storyId, rank, ch, &wg)
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
	fmt.Println(BANNER_LIST)
	fmt.Println()

	for i, story := range l.state.FetchedStories {
		fmt.Printf("%d | %d. %s\n", i+1, story.Rank, story.Title)
	}
}

func (l *ListMode) ChangeState(input string) State {
	numbers, _ := regexp.Compile("^[1-9]$")

	switch {
	case input == "x":
		os.Exit(0)
	case input == "n":
		if MAX_STORIES/NUM_PER_PAGE < l.state.PageNumber+1 {
			l.state.PageNumber = MAX_STORIES / NUM_PER_PAGE
			fmt.Println()
			fmt.Println("Last story on Hacker News")
		} else {
			l.state.PageNumber += 1
		}
	case input == "p":
		if l.state.PageNumber-1 < 1 {
			l.state.PageNumber = 1
			fmt.Println()
			fmt.Println("First story on Hacker News")
		} else {
			l.state.PageNumber -= 1
		}
	case numbers.MatchString(input):
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println()
			fmt.Println("Error: parsing input number")
		} else {
			l.state.SelectedStory = l.state.FetchedStories[num-1]
			l.state.Mode = Details
		}
	default:
		fmt.Println()
		fmt.Println("Error: input not supported")
	}

	return l.state
}
