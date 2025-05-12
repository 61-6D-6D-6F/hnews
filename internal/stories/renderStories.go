package stories

import "fmt"

const BANNER = `   __ ___  __             
  / // / |/ /__ _    _____      x - exit
 / _  /    / -_) |/|/ (_-<      n - next
/_//_/_/|_/\__/|__,__/___/      p - prev`

func renderList(stories []Story) {
	fmt.Println(BANNER)
	fmt.Println()

	for i, story := range stories {
		fmt.Printf("%d | %d. %s\n", i+1, story.Rank, story.Title)
	}
}
