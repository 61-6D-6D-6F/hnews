package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"sync"
)

func main() {
	topIds := fetchTopIds()

	currentPage := 1

	runMainLoop(topIds, &currentPage)
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

func runMainLoop(topIds []int, currentPage *int) {
	stories := fetchStories(topIds, *currentPage)

	storiesSorted := sortStories(stories)

	renderStories(storiesSorted)

	scanInput(*&currentPage)

	runMainLoop(topIds, currentPage)
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

func sortStories(stories []Story) []Story {
	sort.Slice(stories, func(left, right int) bool {
		return stories[left].Rank < stories[right].Rank
	})

	return stories
}

const BANNER = `   __ ___  __             
  / // / |/ /__ _    _____      x - exit
 / _  /    / -_) |/|/ (_-<      n - next
/_//_/_/|_/\__/|__,__/___/      p - prev`

func renderStories(stories []Story) {
	fmt.Println(BANNER)
	fmt.Println()

	for i, story := range stories {
		fmt.Printf("%d | %d. %s\n", i+1, story.Rank, story.Title)
	}
}
func scanInput(currentPage *int) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error scanning choice")
		return
	}

	switch input {
	case "x":
		os.Exit(0)
	case "n":
		if 500/9 < *currentPage+1 {
			*currentPage = 500 / 9
		} else {
			*currentPage += 1
		}
	case "p":
		if *currentPage-1 < 1 {
			*currentPage = 1
		} else {
			*currentPage -= 1
		}
	default:
		fmt.Println("Error input not supported")
	}
}
