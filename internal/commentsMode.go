package internal

import (
	"fmt"
	"os"
	"strings"
)

const BANNER_COMMENT = `   __ ___  __             
  / // / |/ /__ _    _____      x - exit
 / _  /    / -_) |/|/ (_-<      b - back    r - replies
/_//_/_/|_/\__/|__,__/___/      n - next    p - prev`

type CommentsMode struct {
	state State
}

func NewCommentsMode(s State) *CommentsMode {
	return &CommentsMode{
		state: s,
	}
}

func (c *CommentsMode) Fetch() {
	commentId := c.state.CurrentSiblings[c.state.CurrentPos]

	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", commentId)

	fetchUrl(url, &c.state.FetchedComment)
}

func (c *CommentsMode) Render() {
	var sB strings.Builder

	if c.state.UiInfo != "" {
		sB.WriteString("\n")
		sB.WriteString(c.state.UiInfo)
		sB.WriteString("\n")
	}

	sB.WriteString(BANNER_COMMENT)
	sB.WriteString("\n\n")

	if c.state.FetchedComment.Deleted == true {
		sB.WriteString("\n")
		sB.WriteString("Deleted comment")

		fmt.Println(sB.String())
		return
	}

	fmt.Fprintf(&sB, "Comment tree :     ")
	for i, pos := range c.state.HistoryPos {
		fmt.Fprintf(&sB, "[ %d / %d ]     ", pos+1, len(c.state.HistorySiblings[i]))
	}
	fmt.Fprintf(&sB, "[ %d / %d ]\n", c.state.CurrentPos+1, len(c.state.CurrentSiblings))
	sB.WriteString("\n")

	fmt.Fprintf(&sB, "By:         %s\n", c.state.FetchedComment.By)
	fmt.Fprintf(&sB, "Replies:    %d\n", len(c.state.FetchedComment.Kids))
	sB.WriteString("\n")
	fmt.Fprintf(&sB, "%s\n", c.state.FetchedComment.Text)

	fmt.Println(sB.String())
}

func (c *CommentsMode) ChangeState(input string) State {
	c.state.UiInfo = ""

	switch input {
	case "x":
		os.Exit(0)
	case "b":
		if len(c.state.HistorySiblings) == 0 {
			c.state.Mode = Details

			return c.state
		}
		c.state.CurrentSiblings = c.state.HistorySiblings[len(c.state.HistorySiblings)-1]
		c.state.CurrentPos = c.state.HistoryPos[len(c.state.HistoryPos)-1]

		c.state.HistorySiblings = c.state.HistorySiblings[:len(c.state.HistorySiblings)-1]
		c.state.HistoryPos = c.state.HistoryPos[:len(c.state.HistoryPos)-1]
	case "r":
		if len(c.state.FetchedComment.Kids) == 0 {
			c.state.UiInfo = "No reply yet"

			return c.state
		}
		c.state.HistorySiblings = append(c.state.HistorySiblings, c.state.CurrentSiblings)
		c.state.HistoryPos = append(c.state.HistoryPos, c.state.CurrentPos)
		c.state.CurrentSiblings = c.state.FetchedComment.Kids
		c.state.CurrentPos = 0
		c.state.FetchedComment.Kids = []int{}
	case "n":
		if len(c.state.CurrentSiblings)-1 < c.state.CurrentPos+1 {
			c.state.CurrentPos = len(c.state.CurrentSiblings) - 1
			c.state.UiInfo = "Last comment of the comment chain"

			return c.state
		}
		c.state.CurrentPos += 1
	case "p":
		if c.state.CurrentPos-1 < 0 {
			c.state.CurrentPos = 0
			c.state.UiInfo = "First comment of the comment chain"

			return c.state
		}
		c.state.CurrentPos -= 1
	default:
		c.state.UiInfo = "Error: input not supported"
	}

	return c.state
}
