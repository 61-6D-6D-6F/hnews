package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
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
