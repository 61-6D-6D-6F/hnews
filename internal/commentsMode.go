package internal

import (
	"fmt"
	"os"
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
	fmt.Println(BANNER_COMMENT)
	fmt.Println()

	if c.state.FetchedComment.Deleted == true {
		fmt.Println()
		fmt.Println("Deleted comment")
		return
	}

	fmt.Printf("Comment tree :     ")
	for i, pos := range c.state.HistoryPos {
		fmt.Printf("[ %d / %d ]     ", pos+1, len(c.state.HistorySiblings[i]))
	}
	fmt.Printf("[ %d / %d ]\n", c.state.CurrentPos+1, len(c.state.CurrentSiblings))
	fmt.Println()

	fmt.Printf("By:         %s\n", c.state.FetchedComment.By)
	fmt.Printf("Replies:    %d\n", len(c.state.FetchedComment.Kids))
	fmt.Println()
	fmt.Printf("%s\n", c.state.FetchedComment.Text)
}

func (c *CommentsMode) ChangeState(input string) State {
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
			fmt.Println()
			fmt.Println("No reply yet")

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
			fmt.Println()
			fmt.Println("Last comment of the reply chain")

			return c.state
		}
		c.state.CurrentPos += 1
	case "p":
		if c.state.CurrentPos-1 < 0 {
			c.state.CurrentPos = 0
			fmt.Println()
			fmt.Println("First comment of the reply chain")

			return c.state
		}
		c.state.CurrentPos -= 1
	default:
		fmt.Println()
		fmt.Println("Error: input not supported")
	}

	return c.state
}
