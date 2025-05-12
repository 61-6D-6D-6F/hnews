package stories

import "fmt"

const BANNER_LIST = `   __ ___  __             
  / // / |/ /__ _    _____      x - exit
 / _  /    / -_) |/|/ (_-<      n - next
/_//_/_/|_/\__/|__,__/___/      p - prev`

const BANNER_DETAIL = `   __ ___  __             
  / // / |/ /__ _    _____      x - exit
 / _  /    / -_) |/|/ (_-<      b - back
/_//_/_/|_/\__/|__,__/___/              `

func renderList(stories []Story) {
	fmt.Println(BANNER_LIST)
	fmt.Println()

	for i, story := range stories {
		fmt.Printf("%d | %d. %s\n", i+1, story.Rank, story.Title)
	}
}

func renderDetail(story Story) {
	fmt.Println(BANNER_DETAIL)
	fmt.Println()

	fmt.Printf("Title:  %s\n", story.Title)
	fmt.Printf("By:     %s\n", story.By)
	fmt.Printf("Url:    %s\n", story.Url)
	fmt.Printf("Comments: %d\n", story.Descendants)
	if story.Text != "" {
		fmt.Println()
		fmt.Printf("%s\n", story.Text)
	}
}
