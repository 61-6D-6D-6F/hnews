package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func fetchTopStoriesIds() []int {
	var topIds []int

	url := ("https://hacker-news.firebaseio.com/v0/topstories.json")
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching ids of top stories")
		return topIds
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&topIds); err != nil {
		fmt.Println("Error decoding ids of top stories response")
		return topIds
	}

	return topIds
}

func fetchStories(ids []int, currentPage int) []Story {
	ch := make(chan Story)
	var wg sync.WaitGroup

	for i, id := range ids[currentPage*9-9 : currentPage*9] {
		rank := currentPage*9 - 8 + i
		wg.Add(1)
		go fetchStory(id, rank, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var stories []Story

	for story := range ch {
		stories = append(stories, story)
	}

	return stories
}

func fetchStory(id int, rank int, ch chan<- Story, wg *sync.WaitGroup) {
	var story Story

	story.Rank = rank

	defer wg.Done()

	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching story id: %d", id)
		return
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&story); err != nil {
		fmt.Printf("Error decoding story id: %d", id)
		return
	}

	ch <- story

	return
}
