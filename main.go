package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(fetchTopIds())
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

	if res.StatusCode != 200 {
		fmt.Println("Hacker News not available")
		return topIds
	}

	if err := json.NewDecoder(res.Body).Decode(&topIds); err != nil {
		fmt.Println("Error decoding ids of top stories response")
		return topIds
	}

	return topIds
}
