package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func fetchStories(storyIds []int, pageNumber int) []Story {
	ch := make(chan Story)
	var wg sync.WaitGroup

	start := pageNumber*NUM_PER_PAGE - NUM_PER_PAGE
	end := pageNumber * NUM_PER_PAGE

	for i, storyId := range storyIds[start:end] {
		rank := pageNumber*NUM_PER_PAGE - 8 + i
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

	return sortStoriesList(stories)
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

func fetchComment(commentId int) Comment {
	var comment Comment

	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", commentId)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println()
		fmt.Printf("Error: fetching comment id: %d", commentId)
		return comment
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&comment); err != nil {
		fmt.Println()
		fmt.Printf("Error: decoding comment id: %d", commentId)
		return comment
	}

	return comment
}
