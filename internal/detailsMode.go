package internal

import (
	"fmt"
	"os"
	"strings"
)

const BANNER_DETAIL = `   __ ___  __             
  / // / |/ /__ _    _____      x - exit
 / _  /    / -_) |/|/ (_-<      b - back
/_//_/_/|_/\__/|__,__/___/      c - comments`

type DetailsMode struct {
	state State
}

func NewDetailsMode(s State) *DetailsMode {
	return &DetailsMode{
		state: s,
	}
}

func (d *DetailsMode) Fetch() {
}

func (d *DetailsMode) Render() {
	var sB strings.Builder

	if d.state.UiInfo != "" {
		sB.WriteString("\n")
		sB.WriteString(d.state.UiInfo)
		sB.WriteString("\n")
	}

	sB.WriteString(BANNER_DETAIL)
	sB.WriteString("\n\n")

	if d.state.SelectedStory.Deleted == true {
		sB.WriteString("\n")
		sB.WriteString("Deleted story")

		fmt.Println(sB.String())
		return
	}

	fmt.Fprintf(&sB, "Title:      %s\n", d.state.SelectedStory.Title)
	fmt.Fprintf(&sB, "By:         %s\n", d.state.SelectedStory.By)
	fmt.Fprintf(&sB, "Url:        %s\n", d.state.SelectedStory.Url)
	fmt.Fprintf(&sB, "Comments:   %d\n", d.state.SelectedStory.Descendants)
	if d.state.SelectedStory.Text != "" {
		sB.WriteString("\n")
		fmt.Fprintf(&sB, "%s\n", d.state.SelectedStory.Text)
	}

	fmt.Println(sB.String())
}

func (d *DetailsMode) ChangeState(input string) State {
	d.state.UiInfo = ""

	switch input {
	case "x":
		os.Exit(0)
	case "b":
		d.state.Mode = List
	case "c":
		if len(d.state.SelectedStory.Kids) == 0 {
			d.state.UiInfo = "No comment yet"

			return d.state
		}
		d.state.CurrentSiblings = d.state.SelectedStory.Kids
		d.state.HistorySiblings = [][]int{}
		d.state.HistoryPos = []int{}
		d.state.CurrentPos = 0
		d.state.Mode = Comments
	default:
		d.state.UiInfo = "Error: input not supported"
	}

	return d.state
}
