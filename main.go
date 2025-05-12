package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
)

func main() {
	topIds := fetchTopIds()
	// fmt.Println(topIds)

	ch := make(chan Story)
	var wg sync.WaitGroup

	for i, id := range topIds {
		rank := i + 1
		wg.Add(1)
		go fetchStory(id, rank, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	storiesSorted := sortStories(ch)

	for _, story := range storiesSorted {
		fmt.Printf("%d. %s\n", story.Rank, story.Title)
	}
}

func fetchTopIds() []int {
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

type Story struct {
	Rank        int
	By          string `json:"by"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Text        string `json:"text"`
	Kids        []int  `json:"kids"`
	Descendants int    `json:"descendants"`
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

func sortStories(ch chan Story) []Story {
	var stories []Story

	for story := range ch {
		stories = append(stories, story)
	}

	sort.Slice(stories, func(left, right int) bool {
		return stories[left].Rank < stories[right].Rank
	})

	return stories
}
