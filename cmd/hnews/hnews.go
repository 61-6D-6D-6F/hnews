package main

import (
	"github.com/61-6D-6D-6F/hnews/internal/stories"
)

func main() {
	topIds := stories.FetchIds()

	stories.ShowList(topIds, &stories.CurrentPage)
}
