package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"sort"
	"sync"
)

func fetchStoryIds() []int {
	var storyIds []int

	url := ("https://hacker-news.firebaseio.com/v0/topstories.json")
	res, err := http.Get(url)
	if err != nil {
		fmt.Println()
		fmt.Println("Error: fetching story ids")
		return storyIds
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&storyIds); err != nil {
		fmt.Println()
		fmt.Println("Error: decoding story ids")
		return storyIds
	}

	return storyIds
}

func fetchStory(storyId int, rank int, ch chan<- Story, wg *sync.WaitGroup) {
	var story Story

	story.Rank = rank

	defer wg.Done()

	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", storyId)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println()
		fmt.Printf("Error: fetching story id: %d", storyId)
		return
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&story); err != nil {
		fmt.Println()
		fmt.Printf("Error: decoding story id: %d", storyId)
		return
	}

	ch <- story

	return
}

func sortStoriesList(stories []Story) []Story {
	sort.Slice(stories, func(left, right int) bool {
		return stories[left].Rank < stories[right].Rank
	})

	return stories
}

func scan() string {
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println()
		fmt.Println("Error: scanning input")
	}

	return input
}

func equalState(s1 State, s2 State) bool {
	if reflect.TypeOf(State{}).NumField() != 10 ||
		s1.Mode != s2.Mode ||
		!slices.Equal(s1.StoryIds, s2.StoryIds) ||
		s1.PageNumber != s2.PageNumber ||
		len(s1.FetchedStories) != len(s2.FetchedStories) ||
		!equalStory(s1.SelectedStory, s2.SelectedStory) ||
		!equalComment(s1.FetchedComment, s2.FetchedComment) ||
		!slices.Equal(s1.CurrentSiblings, s2.CurrentSiblings) ||
		s1.CurrentPos != s2.CurrentPos ||
		len(s1.HistorySiblings) != len(s2.HistorySiblings) ||
		!slices.Equal(s1.HistoryPos, s2.HistoryPos) {
		return false
	}
	for i, story1 := range s1.FetchedStories {
		if !equalStory(story1, s2.FetchedStories[i]) {
			return false
		}
	}
	for i, historySibs1 := range s1.HistorySiblings {
		if !slices.Equal(historySibs1, s2.HistorySiblings[i]) {
			return false
		}
	}

	return true
}

func equalStory(s1 Story, s2 Story) bool {
	if reflect.TypeOf(Story{}).NumField() != 8 ||
		s1.Rank != s2.Rank ||
		s1.By != s2.By ||
		s1.Title != s2.Title ||
		s1.Url != s2.Url ||
		s1.Text != s2.Text ||
		!slices.Equal(s1.Kids, s2.Kids) ||
		s1.Descendants != s2.Descendants ||
		s1.Deleted != s2.Deleted {
		return false
	}

	return true
}

func equalComment(c1 Comment, c2 Comment) bool {
	if reflect.TypeOf(Comment{}).NumField() != 4 ||
		c1.By != c2.By ||
		c1.Text != c2.Text ||
		!slices.Equal(c1.Kids, c2.Kids) ||
		c1.Deleted != c2.Deleted {
		return false
	}

	return true
}
