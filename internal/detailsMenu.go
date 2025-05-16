package internal

import (
	"fmt"
	"os"
)

const BANNER_DETAIL = `   __ ___  __             
  / // / |/ /__ _    _____      x - exit
 / _  /    / -_) |/|/ (_-<      b - back
/_//_/_/|_/\__/|__,__/___/      c - comments`

type DetailsMenu struct {
	state State
}

func NewDetailsMenu(s State) *DetailsMenu {
	return &DetailsMenu{
		state: s,
	}
}

func (d *DetailsMenu) Fetch() {
}

func (d *DetailsMenu) Render() {

	fmt.Println(BANNER_DETAIL)
	fmt.Println()

	if d.state.SelectedStory.Deleted == true {
		fmt.Println()
		fmt.Println("Deleted story")
		return
	}

	fmt.Printf("Title:      %s\n", d.state.SelectedStory.Title)
	fmt.Printf("By:         %s\n", d.state.SelectedStory.By)
	fmt.Printf("Url:        %s\n", d.state.SelectedStory.Url)
	fmt.Printf("Comments:   %d\n", d.state.SelectedStory.Descendants)
	if d.state.SelectedStory.Text != "" {
		fmt.Println()
		fmt.Printf("%s\n", d.state.SelectedStory.Text)
	}
}

func (d *DetailsMenu) Scan() State {
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println()
		fmt.Println("Error: scanning input")
		return d.state
	}

	switch input {
	case "x":
		os.Exit(0)
	case "b":
		d.state.Mode = List
	case "c":
		if len(d.state.SelectedStory.Kids) == 0 {
			fmt.Println()
			fmt.Println("No comment yet")
		} else {
			d.state.CurrentSiblings = d.state.SelectedStory.Kids
			d.state.HistorySiblings = [][]int{}
			d.state.HistoryPos = []int{}
			d.state.CurrentPos = 0
			d.state.Mode = Comments
		}
	default:
		fmt.Println()
		fmt.Println("Error: input not supported")
	}
	return d.state
}
